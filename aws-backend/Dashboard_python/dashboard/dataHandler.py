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

        waba_msg = self.messages.loc[self.messages['channel'] == 'WABA']

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

