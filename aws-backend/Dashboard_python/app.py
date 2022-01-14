from flask import Flask
from dashboard import output

import schedule

app = Flask(__name__)

schedule.every().day.at().do(output.Output().insert_data())


@app.route('/')
def hello_world():  # put application's code here
    return 'Hello World!'


if __name__ == '__main__':
    app.run()
