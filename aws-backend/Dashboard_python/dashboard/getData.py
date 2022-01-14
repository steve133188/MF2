import boto3
import time
import json
import math

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
            msgs = self.msg_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
            msgs_data.extend(msgs['Items'])

        all_msgs_counts = msgs['Count']
        print('msg_count ', all_msgs_counts)

        return msgs_data

    #################################################################################
    def get_user(self):
        users = self.user_table.scan()
        users_data = users['Items']

        while users.get('LastEvaluatedKey'):
            users = self.user_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
            users_data.extend(users['Items'])

        all_users_counts = users['Count']
        print('user_count ', all_users_counts)

        return users_data

    #################################################################################
    def get_role(self):
        roles = self.role_table.scan()
        roles_data = roles['Items']

        while roles.get('LastEvaluatedKey'):
            roles = self.role_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
            roles_data.extend(roles['Items'])

        all_roles_counts = roles['Count']
        print('role_count ', all_roles_counts)

        return roles_data

    #################################################################################
    def get_customer(self):
        customers = self.customer_table.scan()
        customers_data = customers['Items']

        while customers.get('LastEvaluatedKey'):
            customers = self.customer_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
            customers_data.extend(customers['Items'])

        all_customers_counts = customers['Count']
        print('customers count ', all_customers_counts)

        return customers_data

    def get_tag(self):
        tags = self.tag_table.scan()
        tags_data = tags['Items']

        while tags.get('LastEvaluatedKey'):
            tags = self.tag_table.scan(ExclusiveStartKey=response['LastEvaluatedKey'])
            tags_data.extend(tags['Items'])

        all_tags_counts = tags['Count']
        print('tags count ', all_tags_counts)

        return tags_data

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
            logs = self.log_table.scab(ExclusiveStartKey=response['LastEvaluatedKey'])
            logs_data.extend(logs['Items'])

        newly_added_customers = logs['Count']
        print('logs_count ', newly_added_customers)

        return logs_data
