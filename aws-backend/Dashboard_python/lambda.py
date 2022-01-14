import json
import boto3
import time
import math

from boto3.dynamodb.conditions import Key
from dashboard import dataHandler


def lambda_handler(event, context):

    dynamodb = boto3.resource('dynamodb', region_name='ap-southeast-1')
    user_table = dynamodb.Table('MF2_TCO_USER')
    customer_table = dynamodb.Table('MF2_TCO_CUSTOMER')
    tag_table = dynamodb.Table('MF2_TCO_TAG')
    msg_table = dynamodb.Table('MessageTable')
    log_table = dynamodb.Table('Activity')
    # TODO implement
    # users = user_table.scan()
    # users_list = []
    # for user in users:
    #     users_list.append(user)
    global time
    print(time.time())
    now = round(time.time())
    end = str(now)
    start = str(now - 3600 * 24 * 1)
    print(start)
    print(end)

    # obtain message data
    msg_filter = {
        'FilterExpression': '#ts between :s and :e',
        'ExpressionAttributeValues': {
            ':s': start,
            ':e': end
        },
        'ExpressionAttributeNames': {
            '#ts': 'timestamp'
        }
    }

    msgs = msg_table.scan(**msg_filter)
    msgs_data = msgs['Items']

    while msgs.get('LastEvaluatedKey'):
        msgs = msg_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
        msgs_data.extend(msgs['Items'])

    all_msgs_counts = msgs['Count']

    def customSort(k):
        return k['timestamp']

    # msgs_data.sort(key=customSort, reverse=False)
    # print('msg_data sort test',msgs_data)
    print('msg_count ', all_msgs_counts)
    #################################################################################

    # obtain user data: user counts
    users = user_table.scan()
    users_data = users['Items']

    while users.get('LastEvaluatedKey'):
        users = user_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
        users_data.extend(users['Items'])

    all_users_counts = users['Count']
    print('user_count ', all_users_counts)
    #################################################################################

    # obtain customer data: custoemr count
    customers = customer_table.scan()
    customers_data = customers['Items']

    while customers.get('LastEvaluatedKey'):
        customers = customer_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
        customers_data.extend(customers['Items'])

    all_customers_counts = customers['Count']
    print('customers count ', all_customers_counts)
    #################################################################################

    # obtain tag data
    tag_filter = {
        'ProjectionExpression': "tag_id, tag_name"
    }
    tags = tag_table.scan(**tag_filter)
    tags_data = tags['Items']

    while tags.get('LastEvaluatedKey'):
        tags = tag_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
        tags_data.extend(tags['Items'])

    all_tags_counts = tags['Count']
    #################################################################################

    # obtain activity log of creating customer
    log_filter = {
        'FilterExpression': '#ts between :s and :e AND #ac = :cc',
        'ExpressionAttributeValues': {
            ':s': start,
            ':e': end,
            ':cc': 'CREATED_CUSTOMER'
        },
        'ExpressionAttributeNames': {
            '#ts': 'timestamp',
            '#ac': 'action'
        }
    }

    logs = log_table.scan(**log_filter)
    logs_data = logs['Items']

    while logs.get('LastEvaluatedKey'):
        logs = log_table.scab(ExclusiveStartKey=response['LastEvaluatedKey'])
        logs_data.extend(logs['Items'])

    newly_added_customers = logs['Count']
    print('logs_count ', newly_added_customers)

    #################################################################################
    # count assigned contacts
    all_assigned_contacts = 0
    for customer in customers_data:
        if len(customer['agents_id']) != 0:
            all_assigned_contacts += 1

    # for comparing
    whatsapp_total_msg_sent = 0
    whatsapp_total_msg_rec = 0
    waba_total_msg_sent = 0
    waba_total_msg_rec = 0
    for msg in msgs_data:
        if msg['channel'] == 'WABA':
            if msg['from_me']:
                waba_total_msg_sent += 1
            else:
                waba_total_msg_rec += 1
        elif msg['channel'] == 'Whatapp':
            if msg['from_me']:
                whatsapp_total_msg_sent += 1
            else:
                whatsapp_total_msg_rec += 1

    print('waba_total_msg_sent ', waba_total_msg_sent)
    print('waba_total_msg_rec ', waba_total_msg_rec)
    print('whatsapp_total_msg_sent ', whatsapp_total_msg_sent)
    print('whatsapp_total_msg_rec ', whatsapp_total_msg_rec)
    # for comparing
    #################################################################################

    # avg response time and avg 1st response time
    waba_total_resp_time = 0
    waba_total_resp_time_count = 0

    waba_total_active_contacts_count = 0

    # waba_total_first_resp_time = 0
    # waba_total_first_resp_time_count = 0

    waba_total_msg_sent = 0
    waba_total_msg_rec = 0

    # waba_most_communication_time = 0
    # waba_roomlist = []
    waba_active_list = []
    # waba_first_resp_list = []

    whatsapp_total_resp_time = 0
    whatsapp_total_resp_time_count = 0

    whatsapp_total_first_resp_time = 0
    whatsapp_total_first_resp_time_count = 0

    whatsapp_avg_resp_time = []
    whatsapp_avg_first_resp_time = []

    waba_temp_room_list = []
    whatsapp_temp_user_list = []

    # get most communication hours list
    waba_communication_list = []
    global z
    for x in range(24):
        start_time = str(now - (24 - x) * 3600)
        end_time = str(now - (24 - x - 1) * 3600)
        msg_filter = {
            'FilterExpression': '#ts between :s and :e',
            'ExpressionAttributeValues': {
                ':s': start_time,
                ':e': end_time
            },
            'ExpressionAttributeNames': {
                '#ts': 'timestamp'
            }
        }

        comm_msgs = msg_table.scan(**msg_filter)
        comm_msgs_data = comm_msgs['Items']

        while comm_msgs.get('LastEvaluatedKey'):
            comm_msgs = msg_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
            comm_msgs_data.extend(comm_msgs['Items'])

        if comm_msgs['Count'] == 0:
            waba_communication_list.append(0)
        else:
            comm_time = 0
            for y in range(comm_msgs['Count']):
                if comm_msgs_data[y]['channel'] == 'WABA':
                    if comm_msgs_data[y]['from_me']:
                        # z = y + 1
                        for z in range(y + 1, comm_msgs['Count']):
                            if (not comm_msgs_data[z]['from_me']) & (
                                    comm_msgs_data[z]['room_id'] == comm_msgs_data[y]['room_id']):
                                time = int(comm_msgs_data[z]['timestamp']) - int(comm_msgs_data[y]['timestamp'])
                                if time > comm_time:
                                    comm_time = time
                                    break
                    elif not comm_msgs_data[y]['from_me']:
                        # z = y + 1
                        for z in range(y + 1, comm_msgs['Count']):
                            if (comm_msgs_data[z]['from_me']) & (
                                    comm_msgs_data[z]['room_id'] == comm_msgs_data[y]['room_id']):
                                time = int(comm_msgs_data[z]['timestamp']) - int(comm_msgs_data[y]['timestamp'])
                                if time > comm_time:
                                    comm_time = time
                                    break
            waba_communication_list.append(comm_time)

    global j
    for i in range(all_msgs_counts):
        if msgs_data[i]['channel'] == 'WABA':
            # if not msgs_data[i]['room_id'] in waba_roomlist:
            #   waba_roomlist.append(msgs_data[i]['room_id'])

            if msgs_data[i]['from_me']:
                waba_total_msg_sent += 1
                # j = i + 1
                for j in range(i + 1, all_msgs_counts):
                    # find the next msg with same room id
                    if msgs_data[j]['room_id'] == msgs_data[i]['room_id']:
                        if not msgs_data[j]['from_me']:
                            # bi-direction communication checking for active contacts
                            if not msgs_data[j]['room_id'] in waba_active_list:
                                waba_total_active_contacts_count += 1
                                waba_active_list.append(msgs_data[j]['room_id'])

                            # time = int(msgs_data[j]['timestamp']) - int(msgs_data[i]['timestamp'])

                            # if time > waba_most_communication_time:
                            #     waba_most_communication_time = time

            else:
                waba_total_msg_rec += 1
                # j = i + 1
                for j in range(i + 1, all_msgs_counts):
                    if msgs_data[j]['room_id'] == msgs_data[i]['room_id']:
                        if msgs_data[j]['from_me']:
                            # bi-direction communication checking for active contacts
                            if not msgs_data[j]['room_id'] in waba_active_list:
                                waba_total_active_contacts_count += 1
                                waba_active_list.append(msgs_data[j]['room_id'])

                            # first resp checking
                            # if not msgs_data[j]['room_id'] in waba_first_resp_list:
                            #     time = int(msgs_data[j]['timestamp']) - int(msgs_data[i]['timestamp'])

                            time = int(msgs_data[j]['timestamp']) - int(msgs_data[i]['timestamp'])
                            waba_total_resp_time += time
                            waba_total_resp_time_count += 1

    # obtain tag data
    tag_result_data = {}
    for i in range(all_tags_counts):
        tag_id = tags_data[i]['tag_id']
        tag_name = tags_data[i]['tag_name']

        cust_filter = {
            'FilterExpression': 'contains(tags_id, :tagid)',
            'ExpressionAttributeValues': {

                ':tagid': tag_id
            }
        }

        custs = customer_table.scan(**cust_filter)
        custs_data = custs['Items']

        while custs.get('LastEvaluatedKey'):
            custs = custs_data.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
            custs_data.extend(custs['Items'])

        tag_result_data[tag_name] = custs['Count']
        # print(tag_result_data[tag_name])

    tag_json_data = json.dumps(tag_result_data)

    # total number of customer with channel waba
    cust_filter = {
        'FilterExpression': 'contains(#ch, :chName)',
        'ExpressionAttributeValues': {
            ':chName': 'WABA'
        },
        'ExpressionAttributeNames': {
            '#ch': 'channels'
        }
    }
    custs = customer_table.scan(**cust_filter)
    custs_data = custs['Items']

    while custs.get('LastEvaluatedKey'):
        custs = customer_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
        custs_data.extend(custs['Items'])

    all_custs_counts_channel = custs['Count']

    print('---------------------------------------------')
    print('channel customer number ', all_custs_counts_channel)
    print('all contacts ', all_customers_counts)
    print('active contacts ', waba_total_active_contacts_count)
    print('total message sent ', waba_total_msg_sent)
    print('total message receive ', waba_total_msg_rec)
    print('waba_total_resp_time', waba_total_resp_time)
    print('waba_total_resp_time_count ', waba_total_resp_time_count)
    print('average response time ', round(waba_total_resp_time / waba_total_resp_time_count) / 60)  # mins
    print('most communication hours ', waba_communication_list)  # second
    print('newly added customers ', newly_added_customers)
    print('all tags ', all_tags_counts)
    print('tag result ', tag_json_data)

    agent = dataHandler.DataHandler()
    agent.get_msg_number()
    agent.get_agent_dashboard()
