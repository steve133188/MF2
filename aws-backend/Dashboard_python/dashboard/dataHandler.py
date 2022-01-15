from dashboard.getData import GetData
import pandas as pd


# from dynamodb_json import json_util as json

class DataHandler:

    def __init__(self):
        self.data = GetData()
        self.messages = self.data.get_message()
        self.users = self.data.get_user()
        self.customers = self.data.get_customer()
        self.tags = self.data.get_tag()
        self.logs = self.data.get_log()
        self.roles = self.data.get_role()

    #################################################################################
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

        return waba_total_msg_sent, waba_total_msg_rec, whatsapp_total_msg_sent, waba_total_msg_rec

    #################################################################################
    def get_communication_hour(self):
        waba_com, wts_com = self.data.get_communication()
        print('WABA communication hour ', waba_com)
        print('Whatsapp communication hour ', wts_com)

        return waba_com, wts_com

    #################################################################################
    def get_new_contact(self):
        waba_newly_contacts = len(self.logs.loc[self.customer['payload'] == 'WABA'])
        print('WABA newly added customer ', waba_newly_contacts)

        wts_newly_contacts = len(self.logs.loc[self.customer['payload'] == 'Whatsapp'])
        print('Whatsapp newly added customer ', wts_newly_contacts)

        return waba_newly_contacts, wts_newly_contacts

    #################################################################################
    def get_all_contact(self):
        waba_all_contact = len(self.customers.loc[self.customers['channel'].isin('WABA')])
        print('WABA all contact ', waba_all_contact)

        return waba_all_contact

    #################################################################################
    def get_waba_contact(self):

        # avg response time and avg 1st response time
        waba_total_msg_sent = 0
        waba_total_msg_rec = 0

        waba_total_resp_time = 0
        waba_total_resp_time_count = 0
        waba_first_time = 0
        waba_longest_time = 0

        waba_total_active_contacts_count = 0
        waba_active_list = []

        waba_msg = self.messages.loc[self.messages['channel'] == 'WABA']
        for i in range(waba_msg):
            if waba_msg[i]['from_me']:
                waba_total_msg_sent += 1
                # j = i + 1
                for j in range(i + 1, waba_msg):
                    # find the next msg with same room id
                    if waba_msg[j]['room_id'] == waba_msg[i]['room_id']:
                        if not waba_msg[j]['from_me']:
                            # bi-direction communication checking for active contacts
                            if not waba_msg[j]['room_id'] in waba_active_list:
                                waba_total_active_contacts_count += 1
                                waba_active_list.append(waba_msg[j]['room_id'])
            else:
                waba_total_msg_rec += 1
                # j = i + 1
                for j in range(i + 1, waba_msg):
                    if waba_msg[j]['room_id'] == waba_msg[i]['room_id']:
                        if waba_msg[j]['from_me']:
                            # bi-direction communication checking for active contacts
                            if not waba_msg[j]['room_id'] in waba_active_list:
                                waba_total_active_contacts_count += 1
                                waba_active_list.append(waba_msg[j]['room_id'])

                            time = int(waba_msg[j]['timestamp']) - int(waba_msg[i]['timestamp'])
                            waba_total_resp_time += time
                            waba_total_resp_time_count += 1

                            if waba_first_time == 0:
                                waba_first_time = time
                            if time > waba_longest_time:
                                waba_longest_time = time

        if waba_total_resp_time_count != 0:
            waba_total_resp_time = waba_total_resp_time / waba_total_resp_time_count

        waba_data = {'active_contacts': waba_total_active_contacts_count,
                     'msg_sent': waba_total_msg_sent,
                     'msg_recv': waba_total_msg_rec,
                     'resp_time': waba_total_resp_time,
                     'first_time': waba_first_time,
                     'longest_time': waba_longest_time
                     }

        print('---------------------------------------------')
        print('active contacts ', waba_total_active_contacts_count)
        print('total message sent ', waba_total_msg_sent)
        print('total message receive ', waba_total_msg_rec)
        print('average response time ', waba_total_resp_time)  # mins

        return waba_data

    #################################################################################
    def get_wts_agent_dashboard(self):

        assigned_list = self.customers.loc[(self.customers['agents_id'].str.len() != 0)]
        print(assigned_list['agents_id'])
        wts_assigned_contacts = len(assigned_list)
        print('Whatsapp assigned customer ', wts_assigned_contacts)

        agent_dashboard = []
        print([self.users])
        for user in self.users.index:

            print(self.users['username'][user])
            print(self.users['role_id'][user])
            print(self.roles.loc[self.roles['role_id'] == self.users['role_id'][user]])
            print('===================================================================')
            user_dash = {'Name': self.users['username'][user],
                         'Role': self.roles.loc[self.roles['role_id'] == self.users['role_id'][user]]['role_name'],
                         'Status': "",
                         'assigned_contact':
                             len(assigned_list.loc[assigned_list['agents_id'].isin([self.users['user_id'][user]])])
                         }
            print(user_dash)

            user_msg = self.messages.loc[
                self.messages['sender'] == user['user_id'] &
                self.messages['recipient'] == user['user_id'] &
                self.messages['channel'] == 'Whatsapp'
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
                user_dash['longest_response_time'] = 0
            else:
                user_dash['avg_response_time'] = sum(resp_time) / len(resp_time)
                user_dash['first_response_time'] = resp_time[0]
                user_dash['longest_response_time'] = max(resp_time)

            agent_dashboard.append(user_dash)

        agent = pd.DataFrame(agent_dashboard)
        wts_dashboard = {'assigned_contacts': wts_assigned_contacts,
                         'active_contacts': agent['active_contact'].sum(),
                         'delivered_contacts': agent['delivered_contact'].sum(),
                         'unhandled_contacts': agent['unhandled_contact'].sum(),
                         'message_sent': agent['message_sent'].sum(),
                         'message_recv': agent['message_recv'].sum(),
                         'avg_response_time': agent['avg_response_time'].mean(),
                         'avg_first_response_time': agent['first_response_time'].mean()
                         }

        print(agent_dashboard, wts_dashboard)
