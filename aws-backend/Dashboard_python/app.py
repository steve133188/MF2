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
    print(df_dashboard)
    output_data = {}
    for (name, data) in df_dashboard.iteritems():
        if name == 'agents':
            agent_temp = pd.DataFrame(data.values.tolist())
            agent_output = {}
            print(agent_temp)
            for (agent_column, agent_data) in agent_temp.iteritems():
                if agent_column == 'role' or agent_column == 'status' or agent_column == 'team' or \
                        agent_column == 'username':
                    agent_output[agent_column] = agent_temp[agent_column][len(agent_temp) - 1]
                    continue

                column_temp = agent_temp[agent_column][0]
                for i in range(1, len(agent_temp[agent_column])):
                    column_temp += agent_temp[agent_column][i]

                sum_df = pd.DataFrame.from_dict(column_temp).mean().astype(int)
                temp_put = []
                for (column, column_data) in sum_df.iteritems():
                    temp_put.append({column: column_data})
                agent_output[agent_column] = temp_put

            output_data[name] = agent_output

        elif name == 'communication_hours':
            temp = pd.DataFrame(data.values.tolist())
            waba_temp = pd.DataFrame(temp['WABA'].to_list()).sum()
            wts_temp = pd.DataFrame(temp['Whatsapp'].to_list()).sum()
            output_data[name] = {'WABA': list(map(int, waba_temp.to_list())),
                                 'Whatsapp': list(map(int, wts_temp.to_list()))}

        elif name == 'PK' or name == 'timestamp' or name == 'tags':
            print(name)

        elif name == 'unhandled_contacts' or name == 'delivered_contacts' or name == 'assigned_contacts':
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
    start = str(1642200846)
    print('===================================================================')
    test_obj = output.Output(1, start, end)
    print('|||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||')
    return test_obj.insert_data()


@app.route('/default')
def default():  # put application's code here
    if len(default_data) == 0:
        get_data(round(time.time()) - 3600 * 24 * 7 - 3600, round(time.time()), 1)

    return default_data[0]


# now = round(time.time())
# print(now)
# end = str(now)
# start = str(now - 3600 * 24 * 1)
# return output.Output(0, start, end)
# return 'dashboard api is running'


@app.route('/dashboard')
def dashboard():  # put application's code here
    start = request.args.get('start')
    end = request.args.get('end')

    return get_data(start, end, 0)


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
