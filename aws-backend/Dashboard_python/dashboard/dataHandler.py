import pandas as pd
from getData import GetData


# from dynamodb_json import json_util as json

class DataHandler:

    def __init__(self):
        data = GetData()
        self.messages = pd.DataFrame(data.get_message())
        self.users = pd.DataFrame(data.get_user())
        self.customers = pd.DataFrame(data.get_customer())
        self.tags = pd.DataFrame(data.get_tag())
        self.logs = pd.DataFrame(data.get_log())
        self.roles = pd.DataFrame(data.get_role())

    def get_msg_number(self):
        all_assigned_contacts = 0
        for customer in self.customers:
            if len(customer['agents_id']) != 0:
                all_assigned_contacts += 1

        # for comparing
        waba_msg = self.messages.loc[self.messages['channel'] == 'WABA']
        waba_total_msg_sent = len(waba_msg.loc[waba_msg['from_me']])
        waba_total_msg_rec = len(waba_msg) - waba_total_msg_sent

        wts_msg = self.messages.loc[self.messages['channel'] == 'Whatsapp']
        whatsapp_total_msg_sent = len(wts_msg.loc[wts_msg['from_me']])
        whatsapp_total_msg_rec = len(wts_msg) - whatsapp_total_msg_sent

        print('waba_total_msg_sent ', waba_total_msg_sent)
        print('waba_total_msg_rec ', waba_total_msg_rec)
        print('whatsapp_total_msg_sent ', whatsapp_total_msg_sent)
        print('whatsapp_total_msg_rec ', whatsapp_total_msg_rec)

    def get_waba_contact(self):

        waba_newly_customers = len(self.logs.loc[self.customer['payload'] == 'WABA'])
        print('WABA newly added customer ', waba_newly_customers)

        # avg response time and avg 1st response time
        waba_total_resp_time = 0
        waba_total_resp_time_count = 0

        waba_total_active_contacts_count = 0

        waba_active_list = []
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

    def get_agent_dashboard(self):

        wts_newly_customers = len(self.logs.loc[self.customer['payload'] == 'Whatsapp'])
        print('Whatsapp newly added customer ', wts_newly_customers)

        assigned_list = self.customers.loc[(len(self.customers['agents_id']) != 0)]
        wts_assigned_contacts = len(assigned_list)
        print('Whatsapp assigned customer ', wts_assigned_contacts)

        agent_dashboard = []
        for user in self.users:

            user_dash = {'Name': user['username'],
                         'Role': self.roles.loc[self.roles['role_id'] == user['role_id']]['role_name'],
                         'Status': "",
                         'assigned_contact':
                             len(assigned_list.loc[assigned_list['agents_id'].isin(user['user_id'])])
                         }

            user_msg = self.messages.loc[
                self.messages['sender'] == user['user_id'] &
                self.messages['recipient'] == user['user_id'] &
                self.messages['channel'] == 'WABA'
                ]

            if len(user_msg) == 0:
                break

            # user contact
            for msg in user_msg:
                contact_list = []
                if (contact_list == 0) or (msg['room_id'] in contact_list):
                    room_msg = user_msg.loc[user_msg['room_id'] == msg['room_id']]
                    contact_list.append(msg['room_id'])
                    if room_msg[0]['from_me']:
                        for j in range(1, room_msg):
                            if not room_msg[j]['from_me']:
                                user_dash['active_contact'] += 1
                                break

                            if j >= len(room_msg):
                                user_dash['delivered_contact'] += 1

                    if not room_msg[0]['from_me']:
                        for j in range(1, room_msg):
                            if room_msg[j]['from_me']:
                                user_dash['active_contact'] += 1
                                break

                            if j >= len(room_msg):
                                user_dash['unhandled_contact'] += 1

            user_dash['message_sent'] = len(user_msg.loc[user_msg['from_me']])
            user_dash['message_recv'] = len(user_msg) - user_dash['message_sent']

            resp_time = []
            for i in range(user_msg):
                if not user_msg[i]['from_me']:
                    for j in range(i, user_msg):
                        if (user_msg[j]['from_me']) & (user_msg[j]['room_id'] == user_msg[i]['room_id']):
                            resp_time.append(int(user[j]['timestamp']) - int(user[i]['timestamp']))
                            break

            if len(resp_time) == 0:
                user_dash['avg_response_time'] = 0
                user_dash['first_response_time'] = 0
            else:
                user_dash['avg_response_time'] = sum(resp_time) / len(resp_time)
                user_dash['first_response_time'] = resp_time[0]

            agent_dashboard.append(user_dash)

        print(agent_dashboard)

