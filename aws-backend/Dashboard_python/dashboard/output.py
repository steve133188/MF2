from dashboard.getData import GetData
from dashboard.dataHandler import DataHandler
import time


class Output:
    def __init__(self, index, start, end):
        self.get_from_logic = DataHandler(0, start, end)
        self.get_from_db = GetData(0, start, end)
        self.end = 0
        if index == 1:
            self.get_from_logic = DataHandler(1, start, end)
            self.get_from_db = GetData(1, start, end)
            self.end = end
            print('Output ', start, end)

    def construct_data(self):

        agent, wts_data = self.get_from_logic.get_wts_agent_dashboard()
        waba_data = self.get_from_logic.get_waba_contact()
        if len(wts_data) == 0 & len(waba_data) == 0:
            return {'PK': 'PK',
                    'timestamp': int(round(time.time())) if self.end == 0 else int(self.end)}
        waba_com, wts_com = self.get_from_logic.get_communication_hour()
        waba_new_contact, wts_new_contact = self.get_from_logic.get_new_contact()
        data_dash = {'PK': 'PK',
                     'timestamp': int(round(time.time())) if self.end != 0 else int(self.end),

                     'communication_hours': {'WABA': waba_com,
                                             'Whatsapp': wts_com},

                     'new_added_contacts': {'WABA': int(waba_new_contact),
                                            'Whatsapp': int(wts_new_contact)},
                     'all_contacts': {'WABA': int(self.get_from_logic.get_all_contact()),
                                      'Whatsapp': int(wts_data['assigned_contacts'] + wts_new_contact)},
                     'active_contacts': {'WABA': int(waba_data['active_contacts']),
                                         'Whatsapp': int(wts_data['active_contacts'])},
                     'delivered_contacts': {'Whatsapp': int(wts_data['delivered_contacts'])},
                     'unhandled_contacts': {'Whatsapp': int(wts_data['unhandled_contacts'])},

                     'total_msg_sent': {'WABA': int(waba_data['msg_sent']),
                                        'Whatsapp': int(wts_data['msg_sent'])},
                     'total_msg_recv': {'WABA': int(waba_data['msg_recv']),
                                        'Whatsapp': int(wts_data['msg_recv'])},

                     'avg_resp_time': {'WABA': int(waba_data['resp_time']),
                                       'Whatsapp': int(wts_data['avg_response_time'])},
                     'avg_first_time': {'WABA': int(waba_data['first_time']),
                                        'Whatsapp': int(wts_data['avg_first_response_time'])},

                     'tags': self.get_from_db.get_all_tags(),

                     'agents': agent
                     }
        return data_dash

    def insert_data(self):
        # dynamodb = boto3.resource('dynamodb', endpoint_url="http://localhost:8000")

        table = self.get_from_db.dynamodb.Table('Mf2_TCO_DASHBOARD')
        items = self.construct_data()
        response = table.put_item(
            Item=items
        )
        print(response)
        return response
