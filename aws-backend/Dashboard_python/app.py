from functools import lru_cache

import flask
import pandas as pd
from flask import Flask, request
from dashboard import output, getData
from apscheduler.schedulers.background import BackgroundScheduler

import time

app = Flask(__name__)
obj = output.Output(0, 0, 0)
default_data = dict()


def get_data(start, end, default_index):
    table = getData.GetData(0, 0, 0).dynamodb.Table('Mf2_TCO_DASHBOARD')
    dashboard_filter = {
        'FilterExpression': '#ts between :s and :e',
        'ExpressionAttributeValues': {
            ':s': int(start),
            ':e': int(end)
        },
        'ExpressionAttributeNames': {
            '#ts': 'timestamp'
        }
    }
    dashboard = table.scan(**dashboard_filter)
    dashboard_data = dashboard['Items']

    while dashboard.get('LastEvaluatedKey'):
        dashboard = table.scan(ExclusiveStartKey=dashboard['LastEvaluatedKey'])
        dashboard_data.extend(dashboard['Items'])

    df_dashboard = pd.DataFrame(dashboard_data)
    output_data = {}
    for (name, data) in df_dashboard.iteritems():
        if name == 'agents':
            agent_temp = []
            for i in range(len(data.values)):
                temp = data.values[i]
                for j in range(len(temp)):
                    agent_temp.append(temp[j])

            agent_df = pd.DataFrame(agent_temp)
            agent_df = agent_df.groupby(['Name', 'Role', 'Status', 'Team']).agg('sum').reset_index()

            output_data[name] = agent_df.T.to_dict(orient='dict')

        elif name == 'communication_hours':
            temp = pd.DataFrame(data.values.tolist())
            waba_temp = pd.DataFrame(temp['WABA'].to_list()).sum()
            wts_temp = pd.DataFrame(temp['Whatsapp'].to_list()).sum()
            output_data[name] = {'WABA': list(map(int, waba_temp.to_list())),
                                 'Whatsapp': list(map(int, wts_temp.to_list()))}

        elif name == 'PK' or name == 'timestamp' or name == 'tags':
            print(name)

        elif name == 'unhandled_contacts' or name == 'delivered_contacts':
            temp = pd.DataFrame(data.values.tolist())
            output_data[name] = {'Whatsapp': list(map(int, temp['Whatsapp'].to_list()))}

        else:
            temp = pd.DataFrame(data.values.tolist())
            output_data[name] = {'WABA': list(map(int, temp['WABA'].to_list())),
                                 'Whatsapp': list(map(int, temp['Whatsapp'].to_list()))}

    output_data['tags'] = df_dashboard['tags'][len(df_dashboard) - 1]

    if default_index == 1:
        default_data[0] = output_data

    return output_data


scheduler = BackgroundScheduler()
scheduler.start()
scheduler.add_job(
    obj.insert_data,
    trigger='cron',
    hour=16,
)
scheduler.add_job(
    get_data,
    trigger='cron',
    hour=16,
    args=(round(time.time()) - 3600 * 24 * 7 - 3600, round(time.time()), 1)
)


@app.route('/test')
def test():
    now = round(time.time())
    end = str(now)
    start = str(now - 3600 * 24)
    print('===================================================================')
    test_obj = output.Output(0, start, end)
    print('|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||')
    print(test_obj.get_from_logic.get_all_contact())
    return 'testing'

@app.route('/')
def default():  # put application's code here
    # if len(default_data) == 0:
    #     get_data(round(time.time()) - 3600 * 24 * 7 - 3600, round(time.time()), 1)
    #
    # return default_data[0]
    now = round(time.time())
    print(now)
    end = str(now)
    start = str(now - 3600 * 24 * 1)
    return output.Output(0, start, end)
    # return 'dashboard api is running'


# @app.route('/dashboard')
# def dashboard():  # put application's code here
#     start = request.args.get('start')
#     end = request.args.get('end')
#
#     return get_data(start, end, 0)


@app.route('/migration')
def migration():  # put application's code here
    start = request.args.get('start')
    end = request.args.get('end')

    i = int(end)
    while i >= int(start):
        migrate = output.Output(1, i - 24 * 7 * 3600, i)
        err = migrate.insert_data()

        i = i - 24 * 3600

    return flask.Response(status=200)


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)
