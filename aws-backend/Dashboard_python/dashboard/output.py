from getData import GetData
from dataHandler import DataHandler


class Output:

    def __init__(self):
        self.get_from_logic = DataHandler()
        self.get_from_db = GetData()

    def construct_data(self):
        agent, wts_data = self.get_from_logic.get_wts_agent_dashboard()
        waba_data = self.get_from_logic.get_waba_contact()
        waba_com, wts_com = self.get_from_logic.get_communication_hour()
        waba_new_contact, wts_new_contact = self.get_from_logic.get_new_contact()
        data_dash = {'communication_hours': {'WABA': waba_com,
                                             'Whatsapp': wts_com},

                     'new_added_contacts': {'WABA': waba_new_contact,
                                            'Whatsapp': wts_new_contact},
                     'all_contacts': {'WABA': self.get_from_logic.get_all_contact(),
                                      'Whatsapp': wts_data['assigned_contacts'] + wts_new_contact},
                     'active_contacts': {'WABA': waba_data['active_contacts'],
                                         'Whatsapp': wts_data['active_contacts']},
                     'delivered_contacts': {'Whatsapp': wts_data['delivered_contacts']},
                     'unhandled_contacts': {'Whatsapp': wts_data['unhandled_contacts']},

                     'total_msg_sent': {'WABA': waba_data['message_sent'],
                                        'Whatsapp': wts_data['msg_sent']},
                     'total_msg_recv': {'WABA': waba_data['message_sent'],
                                        'Whatsapp': wts_data['msg_sent']},

                     'avg_resp_time': {'WABA': waba_data['resp_time'],
                                       'Whatsapp': wts_data['avg_response_time']},
                     'avg_first_time': {'WABA': waba_data['first_time'],
                                        'Whatsapp': wts_data['avg_first_response_time']},
                     
                     'tags': self.get_from_db.get_all_tags(),

                     'agents': agent
                     }
        return data_dash
