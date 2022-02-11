from functools import lru_cache

import flask
import pandas as pd
from flask import Flask, request
from flask_cors import CORS
from dashboard import output, getData
from apscheduler.schedulers.background import BackgroundScheduler

import time

app = Flask(__name__)
default_data = dict()
CORS(app)


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
        dashboard = table.scan(ExclusiveStartKey=dashboard['LastEvaluatedKey'], **dashboard_filter)
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
                # if agent_column == 'role' or agent_column == 'team' or agent_column == 'username' or agent_column == 'status':
                if agent_column == 'status':
                    continue

                if agent_column == 'role' or agent_column == 'team' or agent_column == 'username':
                    agent_output[agent_column] = agent_temp[agent_column][len(agent_temp) - 1]
                    continue

                if agent_column == 'avg_response_time' or agent_column == 'longest_response_time' or \
                        agent_column == 'first_response_time':
                    sum_df = pd.DataFrame(agent_temp[agent_column].to_list()).mean().astype(int)
                    agent_output[agent_column] = sum_df.to_dict()
                    continue
                print(agent_column)
                sum_df = pd.DataFrame(agent_temp[agent_column].to_list()).mean().astype(int)
                agent_output[agent_column] = sum_df.to_dict()

            output_data[name] = agent_output

            # Team info
            temp = pd.DataFrame(agent_output)
            teams = temp.groupby(['team']).sum()
            teams['avg_response_time'] = round(teams['avg_response_time'] / len(teams), 1)
            teams['first_response_time'] = round(teams['first_response_time'] / len(teams), 1)
            teams['longest_response_time'] = round(teams['longest_response_time'] / len(teams), 1)
            output_data['teams'] = teams.to_dict()

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
    output.Output(0, 0, 0).insert_data,
    trigger='cron',
    hour=16,
)
scheduler.add_job(
    get_data,
    trigger='cron',
    hour=16,
    args=(round(time.time()) - 3600 * 24 * 7 - 3600, round(time.time()), 1)
)


@app.route('/')
def start():
    return 'MF2 dashboard server is running'


@app.route('/test')
def test():
    # obj.insert_data()
    return 'testing'


@app.route('/insert')
def insert():
    test_obj = output.Output(0, 0, 0)
    print('|' * 50)
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


@app.route('/dashboard/livechat')
def livechat():  # put application's code here
    start = request.args.get('start')
    end = request.args.get('end')

    original_data = get_data(start, end, 0)
    selected_data = {'active_contacts': original_data['active_contacts'],
                     'all_contacts': original_data['all_contacts'],
                     'total_msg_sent': original_data['total_msg_sent'],
                     'total_msg_recv': original_data['total_msg_recv'],
                     'new_added_contacts': original_data['new_added_contacts'],
                     'avg_resp_time': original_data['avg_resp_time'],
                     'communication_hours': original_data['communication_hours'],
                     'tags': original_data['tags']
                     }

    return selected_data


@app.route('/dashboard/agent')
def agent():  # put application's code here
    start = request.args.get('start')
    end = request.args.get('end')

    return get_data(start, end, 0)


@app.route('/dashboard/team')
def team():  # put application's code here
    start = request.args.get('start')
    end = request.args.get('end')

    original_data = get_data(start, end, 0)
    temp = pd.DataFrame(original_data['agents'])
    teams = temp.groupby(['team']).sum()
    teams['avg_response_time'] = round(teams['avg_response_time'] / len(teams), 1)
    teams['first_response_time'] = round(teams['first_response_time'] / len(teams), 1)
    teams['longest_response_time'] = round(teams['longest_response_time'] / len(teams), 1)

    return teams.to_dict()


@app.route('/migration')
def migration():  # put application's code here
    start = request.args.get('start')
    end = request.args.get('end')

    i = int(start)
    while i <= int(end):
        print("#" * 50)
        print(i)
        print("#" * 50)
        migrate = output.Output(1, i, i + 24 * 3600)
        err = migrate.insert_data()
        i = i + 24 * 3600
        time.sleep(5)

    return flask.Response(status=200)


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)
