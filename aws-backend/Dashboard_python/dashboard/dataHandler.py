from dashboard.getData import GetData
import pandas as pd


# from dynamodb_json import json_util as json

class DataHandler:

    def __init__(self, index, start, end):
        self.data = GetData(index, start, end)
        self.messages = self.data.get_message()
        self.users = self.data.get_user()
        self.customers = self.data.get_customer()
        self.tags = self.data.get_tag()
        self.logs = self.data.get_log()
        self.roles = self.data.get_role()
        self.teams = self.data.get_org()
        self.prev_dash = self.data.get_previous_dash()

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
        if len(self.logs) == 0:
            return 0, 0
        waba_newly_contacts = len(self.logs.loc[self.logs['payload'] == 'WABA'])
        print('WABA newly added customer ', waba_newly_contacts)

        wts_newly_contacts = len(self.logs.loc[self.logs['payload'] == 'Whatsapp'])
        print('Whatsapp newly added customer ', wts_newly_contacts)

        return waba_newly_contacts, wts_newly_contacts

    #################################################################################
    def get_all_contact(self):
        customer_list = self.customers['channels']
        wts_all_contact = 0
        waba_all_contact = 0
        for i in range(len(customer_list)):
            if 'WABA' in customer_list[i]:
                waba_all_contact += 1
            if 'Whatsapp' in customer_list[i]:
                wts_all_contact += 1

        # print('Whatsapp all contact ', wts_all_contact)

        return waba_all_contact, wts_all_contact

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

        if self.prev_dash != 0:
            waba_total_msg_sent = self.prev_dash[0]['total_msg_sent']['WABA']
            waba_total_msg_rec = self.prev_dash[0]['total_msg_recv']['WABA']
            waba_total_active_contacts_count = self.prev_dash[0]['active_contacts']['WABA']

        if len(self.messages) == 0:
            return {'active_contacts': waba_total_active_contacts_count,
                    'msg_sent': waba_total_msg_sent,
                    'msg_recv': waba_total_msg_rec,
                    'resp_time': waba_total_resp_time,
                    'first_time': waba_first_time,
                    'longest_time': waba_longest_time
                    }

        waba_msg = self.messages.loc[self.messages['channel'] == 'WABA']
        print('=' * 50)
        print(waba_msg)

        for i in waba_msg.index:
            if waba_msg['from_me'][i]:
                waba_total_msg_sent += 1
                # j = i + 1
                for j in range(i + 1, len(waba_msg)):
                    # find the next msg with same room id
                    if waba_msg['room_id'][j] == waba_msg['room_id'][i]:
                        if not waba_msg['from_me'][j]:
                            # bi-direction communication checking for active contacts
                            if not waba_msg['room_id'][j] in waba_active_list:
                                waba_total_active_contacts_count += 1
                                waba_active_list.append(waba_msg['room_id'][j])
            else:
                waba_total_msg_rec += 1
                # j = i + 1
                for j in range(i + 1, len(waba_msg)):
                    if waba_msg['room_id'][j] == waba_msg['room_id'][i]:
                        if waba_msg['from_me'][j]:
                            # bi-direction communication checking for active contacts
                            if not waba_msg['room_id'][j] in waba_active_list:
                                waba_total_active_contacts_count += 1
                                waba_active_list.append(waba_msg['room_id'][j])

                            time = int(waba_msg['timestamp'][j]) - int(waba_msg['timestamp'][i])
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
                     'resp_time': round(waba_total_resp_time / 60, 2),
                     'first_time': round(waba_first_time / 60, 2),
                     'longest_time': round(waba_longest_time / 60, 2)
                     }

        print('---------------------------------------------')
        print(waba_data)

        return waba_data

    #################################################################################
    def get_wts_agent_dashboard(self):
        # print(self.customers['agents_id'].map(len) > 0)
        assigned_list = self.customers[self.customers['agents_id'].map(len) > 0]
        wts_assigned_contacts = len(assigned_list)
        print('Whatsapp assigned customer ', wts_assigned_contacts)

        agent_dashboard = {'username': [],
                           'role': [],
                           'team': [],
                           'status': [],
                           'assigned_contact': [],
                           'active_contact': [],
                           'delivered_contact': [],
                           'unhandled_contact': [],
                           'message_sent': [],
                           'message_recv': [],
                           'avg_response_time': [],
                           'first_response_time': [],
                           'longest_response_time': []
                           }

        wts_dashboard = {'assigned_contacts': wts_assigned_contacts,
                         'active_contacts': 0,
                         'delivered_contacts': 0,
                         'unhandled_contacts': 0,
                         'msg_sent': 0,
                         'msg_recv': 0,
                         'avg_response_time': 0,
                         'avg_first_response_time': 0
                         }

        if self.prev_dash != 0:
            wts_dashboard['active_contacts'] = int(self.prev_dash[0]['active_contacts']['Whatsapp'])
            wts_dashboard['delivered_contacts'] = int(self.prev_dash[0]['delivered_contacts']['Whatsapp'])
            wts_dashboard['unhandled_contacts'] = int(self.prev_dash[0]['unhandled_contacts']['Whatsapp'])
            wts_dashboard['msg_sent'] = int(self.prev_dash[0]['total_msg_sent']['Whatsapp'])
            wts_dashboard['msg_recv'] = int(self.prev_dash[0]['total_msg_recv']['Whatsapp'])

        print('User length: ', len(self.users))

        for user in self.users.index:
            if self.users['user_id'][user] == 1:
                continue

            team = 'None'
            if self.users['team_id'][user] != 0:
                team = self.teams.loc[self.teams['org_id'] == self.users['team_id'][user]]['name'].to_string(
                    index=False)
            print('=' * 50)
            user_id = str(self.users['user_id'][user])
            agent_dashboard['username'].append({user_id: self.users['username'][user]})
            agent_dashboard['team'].append({user_id: team})
            agent_dashboard['role'].append({user_id: self.roles.loc[self.roles['role_id'] ==
                                                                    self.users['role_id'][user]]['role_name'].to_string(
                index=False)})

            agent_dashboard['status'].append({user_id: ''})
            assignee_count = 0
            for assignee in assigned_list['agents_id']:
                if self.users['user_id'][user] in assignee:
                    assignee_count += 1
            agent_dashboard['assigned_contact'].append({user_id: assignee_count})

            contact_list = []
            active_contact = 0
            delivered_contact = 0
            unhandled_contact = 0
            agent_msg_sent = 0
            agent_msg_recv = 0

            if self.prev_dash != 0:
                active_contact = int(next(item for item in self.prev_dash[0]['agents']['active_contact']
                                          if user_id in item.keys())[user_id])
                delivered_contact = int(next(item for item in self.prev_dash[0]['agents']['delivered_contact']
                                             if user_id in item.keys())[user_id])
                unhandled_contact = int(next(item for item in self.prev_dash[0]['agents']['unhandled_contact']
                                             if user_id in item.keys())[user_id])
                agent_msg_sent = int(next(item for item in self.prev_dash[0]['agents']['message_sent']
                                          if user_id in item.keys())[user_id])
                agent_msg_recv = int(next(item for item in self.prev_dash[0]['agents']['message_recv']
                                          if user_id in item.keys())[user_id])

            if len(self.messages) == 0 or len(self.messages.loc[self.messages['channel'] == 'Whatsapp']) == 0:
                agent_dashboard['active_contact'].append({user_id: active_contact})
                agent_dashboard['delivered_contact'].append({user_id: delivered_contact})
                agent_dashboard['unhandled_contact'].append({user_id: unhandled_contact})
                agent_dashboard['message_sent'].append({user_id: agent_msg_sent})
                agent_dashboard['message_recv'].append({user_id: agent_msg_recv})
                agent_dashboard['avg_response_time'].append({user_id: 0})
                agent_dashboard['first_response_time'].append({user_id: 0})
                agent_dashboard['longest_response_time'].append({user_id: 0})
                continue
            else:
                wts_msg = self.messages.loc[self.messages['channel'] == 'Whatsapp']
                print('Whatsapp Messages length: ', len(wts_msg))

            user_msg = wts_msg.loc[
                (wts_msg['sender'] == str(self.users['user_id'][user])) |
                (wts_msg['recipient'] == str(self.users['user_id'][user]))
                ].reset_index()
            print('User msg: ', len(user_msg))
            print('=' * 50)

            if len(user_msg) == 0:
                agent_dashboard['active_contact'].append({user_id: active_contact})
                agent_dashboard['delivered_contact'].append({user_id: delivered_contact})
                agent_dashboard['unhandled_contact'].append({user_id: unhandled_contact})
                agent_dashboard['message_sent'].append({user_id: agent_msg_sent})
                agent_dashboard['message_recv'].append({user_id: agent_msg_recv})
                agent_dashboard['avg_response_time'].append({user_id: 0})
                agent_dashboard['first_response_time'].append({user_id: 0})
                agent_dashboard['longest_response_time'].append({user_id: 0})
                continue

            # user contact
            for msg in user_msg.index:

                if (len(contact_list) == 0) or not (user_msg['room_id'][msg] in contact_list):
                    room_msg = user_msg.loc[user_msg['room_id'] == user_msg['room_id'][msg]].reset_index()
                    contact_list.append(user_msg['room_id'][msg])
                    if room_msg['from_me'][0]:
                        if len(room_msg) == 1:
                            delivered_contact += 1
                            wts_dashboard['delivered_contacts'] += 1
                        else:
                            for j in range(1, len(room_msg)):
                                if not room_msg['from_me'][j]:
                                    active_contact += 1
                                    wts_dashboard['active_contacts'] += 1
                                    break

                                elif j == len(room_msg) - 1:
                                    delivered_contact += 1
                                    wts_dashboard['delivered_contacts'] += 1

                    if not room_msg['from_me'][0]:
                        if len(room_msg) == 1:
                            unhandled_contact += 1
                            wts_dashboard['unhandled_contacts'] += 1
                        else:
                            for j in range(1, len(room_msg)):
                                if room_msg['from_me'][j]:
                                    active_contact += 1
                                    wts_dashboard['active_contacts'] += 1
                                    break

                                elif j == len(room_msg) - 1:
                                    unhandled_contact += 1
                                    wts_dashboard['unhandled_contacts'] += 1

            agent_dashboard['active_contact'].append({user_id: active_contact})
            agent_dashboard['delivered_contact'].append({user_id: delivered_contact})
            agent_dashboard['unhandled_contact'].append({user_id: unhandled_contact})

            msg_sent = len(user_msg.loc[user_msg['from_me']])
            agent_dashboard['message_sent'].append({user_id: msg_sent + agent_msg_sent})
            agent_dashboard['message_recv'].append({user_id: len(user_msg) - msg_sent + agent_msg_recv})
            wts_dashboard['msg_sent'] += msg_sent
            wts_dashboard['msg_recv'] += len(user_msg) - msg_sent

            print('=' * 50)
            print('active_contact', agent_dashboard['active_contact'],
                  '\n delivered_contact', agent_dashboard['delivered_contact'],
                  '\n unhandled_contact', agent_dashboard['unhandled_contact'],
                  '\n message_sent', agent_dashboard['message_sent'],
                  '\n message_recv', agent_dashboard['message_recv'])

            resp_time = []
            for i in user_msg.index:
                if not user_msg['from_me'][i]:
                    for j in range(i, len(user_msg)):
                        if (user_msg['from_me'][j]) and (user_msg['room_id'][j] == user_msg['room_id'][i]):
                            resp_time.append(int(user_msg['timestamp'][j]) - int(user_msg['timestamp'][i]))
                            break

            if len(resp_time) == 0:
                agent_dashboard['avg_response_time'].append({user_id: 0})
                agent_dashboard['first_response_time'].append({user_id: 0})
                agent_dashboard['longest_response_time'].append({user_id: 0})
            else:
                avg_time = round(sum(resp_time) / len(resp_time) / 60, 2)
                first_time = round(resp_time[0] / 60, 2)
                long_time = round(max(resp_time) / 60, 2)

                agent_dashboard['avg_response_time'].append({user_id: avg_time})
                agent_dashboard['first_response_time'].append({user_id: first_time})
                agent_dashboard['longest_response_time'].append({user_id: long_time})

                wts_dashboard['avg_response_time'] += avg_time
                wts_dashboard['avg_first_response_time'] += first_time

        wts_dashboard['avg_response_time'] = round(
            wts_dashboard['avg_response_time'] / len(agent_dashboard['username']), 2)
        wts_dashboard['avg_first_response_time'] = round(
            wts_dashboard['avg_first_response_time'] / len(agent_dashboard['username']), 2)

        print('Agent: \n', agent_dashboard)
        print('Whatsapp Dashboard: \n', wts_dashboard)
        return agent_dashboard, wts_dashboard
