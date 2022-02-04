import boto3
import time
# import math
import pandas as pd
from boto3.dynamodb.conditions import Key, Attr

# from boto3.dynamodb.conditions import Key


class GetData:

    # Dynamodb
    def __init__(self, index, start, end):
        key_id1 = 'AKIATRVR34'
        key_id2 = 'WXY767NBFP'
        access_key1 = 'E/X8SfmdBx0SNRO4q4W'
        access_key2 = 'fTLb0CrNrd+2UL5fO/z1r'

        self.dynamodb = boto3.resource('dynamodb', region_name='ap-southeast-1', aws_access_key_id=key_id1 + key_id2,
                                       aws_secret_access_key=access_key1 + access_key2)
        self.user_table = self.dynamodb.Table('MF2_TCO_USER')
        self.role_table = self.dynamodb.Table('MF2_TCO_ROLE')
        self.customer_table = self.dynamodb.Table('MF2_TCO_CUSTOMER')
        self.org_table = self.dynamodb.Table('MF2_TCO_ORG')
        self.tag_table = self.dynamodb.Table('MF2_TCO_TAG')
        self.msg_table = self.dynamodb.Table('MessageTable')
        self.log_table = self.dynamodb.Table('Activity')
        self.dash_table = self.dynamodb.Table('Mf2_TCO_DASHBOARD')

        if index == 1:
            self.start = start
            self.end = end
        else:
            self.now = round(time.time())
            self.end = self.now
            self.start = self.now - 3600 * 24 * 1 - 5

        print('Get Data ', self.start, self.end)

    #################################################################################
    def get_message(self):
        start_time = str(self.start)
        end_time = str(self.end)
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

        msgs = self.msg_table.scan(**msg_filter)
        msgs_data = msgs['Items']

        while msgs.get('LastEvaluatedKey'):
            msgs = self.msg_table.scan(ExclusiveStartKey=msgs['LastEvaluatedKey'], **msg_filter)
            msgs_data.extend(msgs['Items'])

        all_msgs_counts = len(msgs_data)
        print('msg_count ', all_msgs_counts)

        return pd.DataFrame(msgs_data)

    #################################################################################
    def get_previous_dash(self):
        start_t = int(self.start)
        end_t = int(self.end)
        dash_filter = {
            'FilterExpression': '#ts between :s and :e',
            'ExpressionAttributeValues': {
                ':s': start_t,
                ':e': end_t
            },
            'ExpressionAttributeNames': {
                '#ts': 'timestamp'
            },
            'Limit': 1
        }

        dash = self.dash_table.scan(**dash_filter)
        dash_data = dash['Items']

        while dash.get('LastEvaluatedKey'):
            dash = self.dash_table.scan(ExclusiveStartKey=dash['LastEvaluatedKey'], **dash_filter)
            dash_data.extend(dash['Items'])

        if len(dash_data) == 0:
            return 0

        print('Previous Dash: \n', len(dash_data))
        return dash_data

    #################################################################################
    def get_user(self):
        users = self.user_table.scan()
        users_data = users['Items']

        while users.get('LastEvaluatedKey'):
            users = self.user_table.scan(ExclusiveStartKey=users['LastEvaluatedKey'])
            users_data.extend(users['Items'])

        all_users_counts = users['Count']
        print('user_count ', all_users_counts)

        return pd.DataFrame(users_data)

    #################################################################################
    def get_role(self):
        roles = self.role_table.scan()
        roles_data = roles['Items']

        while roles.get('LastEvaluatedKey'):
            roles = self.role_table.scan(ExclusiveStartKey=roles['LastEvaluatedKey'])
            roles_data.extend(roles['Items'])

        all_roles_counts = roles['Count']
        print('role_count ', all_roles_counts)

        return pd.DataFrame(roles_data)

    #################################################################################
    def get_org(self):
        orgs = self.org_table.scan()
        orgs_data = orgs['Items']

        while orgs.get('LastEvaluatedKey'):
            orgs = self.role_table.scan(ExclusiveStartKey=orgs['LastEvaluatedKey'])
            orgs_data.extend(orgs['Items'])

        all_roles_counts = orgs['Count']
        print('org_count ', all_roles_counts)

        return pd.DataFrame(orgs_data)

    #################################################################################
    def get_customer(self):
        customers = self.customer_table.scan()
        customers_data = customers['Items']

        while customers.get('LastEvaluatedKey'):
            customers = self.customer_table.scan(ExclusiveStartKey=customers['LastEvaluatedKey'])
            customers_data.extend(customers['Items'])

        all_customers_counts = int(customers['Count'])
        print('customers count ', all_customers_counts)

        return pd.DataFrame(customers_data)

    #################################################################################
    def get_log(self):
        start_time = str(self.start)
        end_time = str(self.end)
        log_filter = {
            'FilterExpression': '#ts between :s and :e AND #ac = :cc',
            'ExpressionAttributeValues': {
                ':s': start_time,
                ':e': end_time,
                ':cc': 'CREATED_CUSTOMER'
            },
            'ExpressionAttributeNames': {
                '#ts': 'timestamp',
                '#ac': 'action'
            }
        }

        logs = self.log_table.scan(**log_filter)
        logs_data = logs['Items']

        while logs.get('LastEvaluatedKey'):
            logs = self.log_table.scan(ExclusiveStartKey=logs['LastEvaluatedKey'], **log_filter)
            logs_data.extend(logs['Items'])

        newly_added_customers = logs['Count']
        print('logs_count ', newly_added_customers)

        return pd.DataFrame(logs_data)

    #################################################################################
    def get_communication(self):
        waba_communication_list = []
        wts_communication_list = []
        global z
        now_time = int(self.end)
        for x in range(24):
            start_time = str(now_time - (24 - x) * 3600)
            end_time = str(now_time - (24 - x - 1) * 3600)
            # print('start_time ', start_time)
            # print('end_time ', end_time)
            msg_filter = {
                'ExpressionAttributeValues': {
                    ':s': start_time,
                    ':e': end_time
                },
                'ExpressionAttributeNames': {
                    '#ts': 'timestamp'
                },
                'FilterExpression': '#ts between :s and :e'
            }

            comm_msgs = self.msg_table.scan(**msg_filter)
            comm_msgs_data = comm_msgs['Items']

            while 'LastEvaluatedKey' in comm_msgs:
                comm_msgs = self.msg_table.scan(ExclusiveStartKey=comm_msgs['LastEvaluatedKey'], **msg_filter)
                comm_msgs_data.extend(comm_msgs['Items'])

            count = len(comm_msgs_data)
            if count == 0:
                waba_communication_list.append(0)
                wts_communication_list.append(0)
            else:
                com_msg = pd.DataFrame(comm_msgs_data)
                waba_communication_list.append(
                    len(com_msg.loc[com_msg['channel'] == 'WABA'])
                )
                wts_communication_list.append(
                    len(com_msg.loc[com_msg['channel'] == 'Whatsapp'])
                )

        return waba_communication_list, wts_communication_list

    #################################################################################
    def get_all_tags(self):
        # obtain tag data
        tag_result_data = {}
        tags = self.get_tag()
        for tag in tags.index:

            customer_filter = {
                'FilterExpression': 'contains(tags_id, :tagid)',
                'ExpressionAttributeValues': {
                    ':tagid': tags['tag_id'][tag]
                }
            }

            customers = self.customer_table.scan(**customer_filter)
            customer_data = customers['Items']

            while customers.get('LastEvaluatedKey'):
                customers = customer_data.scan(ExclusiveStartKey=customers['LastEvaluatedKey'], **customer_filter)
                customer_data.extend(customers['Items'])

            tag_result_data[tags['tag_name'][tag]] = customers['Count']
            # print(tag_result_data[tag_name])

        print('tag result ', tag_result_data)

        return tag_result_data

    #################################################################################
    def get_tag(self):
        tags = self.tag_table.scan()
        tags_data = tags['Items']

        while tags.get('LastEvaluatedKey'):
            tags = self.tag_table.scan(ExclusiveStartKey=tags['LastEvaluatedKey'])
            tags_data.extend(tags['Items'])

        all_tags_counts = tags['Count']
        print('tags count ', all_tags_counts)

        return pd.DataFrame(tags_data)
