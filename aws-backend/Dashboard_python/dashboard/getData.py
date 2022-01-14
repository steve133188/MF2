import boto3
import time
import math
import pandas as pd

from boto3.dynamodb.conditions import Key


class GetData:

    # Dynamodb
    def __init__(self):
        dynamodb = boto3.resource('dynamodb', region_name='ap-southeast-1')
        self.user_table = dynamodb.Table('MF2_TCO_USER')
        self.role_table = dynamodb.Table('MF2_TCO_ROLE')
        self.customer_table = dynamodb.Table('MF2_TCO_CUSTOMER')
        self.tag_table = dynamodb.Table('MF2_TCO_TAG')
        self.msg_table = dynamodb.Table('MessageTable')
        self.log_table = dynamodb.Table('Activity')

        now = round(time.time())
        self.end = str(now)
        self.start = str(now - 3600 * 24 * 1)
        print(self.start, self.end)

    #################################################################################
    def get_message(self):
        msg_filter = {
            'FilterExpression': '#ts between :s and :e',
            'ExpressionAttributeValues': {
                ':s': self.start,
                ':e': self.end
            },
            'ExpressionAttributeNames': {
                '#ts': 'timestamp'
            }
        }

        msgs = self.msg_table.scan(**msg_filter)
        msgs_data = msgs['Items']

        while msgs.get('LastEvaluatedKey'):
            msgs = self.msg_table.scan(ExclusiveStartKey=msgs['LastEvaluatedKey'])
            msgs_data.extend(msgs['Items'])

        all_msgs_counts = msgs['Count']
        print('msg_count ', all_msgs_counts)

        return pd.DataFrame(msgs_data)

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
    def get_customer(self):
        customers = self.customer_table.scan()
        customers_data = customers['Items']

        while customers.get('LastEvaluatedKey'):
            customers = self.customer_table.scan(ExclusiveStartKey=customers['LastEvaluatedKey'])
            customers_data.extend(customers['Items'])

        all_customers_counts = customers['Count']
        print('customers count ', all_customers_counts)

        return pd.DataFrame(customers_data)

    #################################################################################
    def get_log(self):
        log_filter = {
            'FilterExpression': '#ts between :s and :e AND #ac = :cc',
            'ExpressionAttributeValues': {
                ':s': self.start,
                ':e': self.end,
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
            logs = self.log_table.scab(ExclusiveStartKey=logs['LastEvaluatedKey'])
            logs_data.extend(logs['Items'])

        newly_added_customers = logs['Count']
        print('logs_count ', newly_added_customers)

        return pd.DataFrame(logs_data)

    #################################################################################
    def get_communication(self):
        waba_communication_list = []
        wts_communication_list = []
        global z
        for x in range(24):
            start_time = str(self.now - (24 - x) * 3600)
            end_time = str(self.now - (24 - x - 1) * 3600)
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

            comm_msgs = self.msg_table.scan(**msg_filter)
            comm_msgs_data = comm_msgs['Items']

            while comm_msgs.get('LastEvaluatedKey'):
                comm_msgs = self.msg_table.scan(ExclusiveStartKey=comm_msgs['LastEvaluatedKey'])
                comm_msgs_data.extend(comm_msgs['Items'])

            if comm_msgs['Count'] == 0:
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
        for tag in tags:

            customer_filter = {
                'FilterExpression': 'contains(tags_id, :tagid)',
                'ExpressionAttributeValues': {

                    ':tagid': tag['tag_id']
                }
            }

            customers = self.customer_table.scan(**customer_filter)
            customer_data = custs['Items']

            while custs.get('LastEvaluatedKey'):
                custs = customer_data.scan(ExclusiveStartKey=customers['LastEvaluatedKey'])
                customer_data.extend(custs['Items'])

            tag_result_data[tag['tag_name']] = custs['Count']
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
