from flask import Flask
from dashboard import output

app = Flask(__name__)


@app.route('/')
def hello_world():  # put application's code here
    obj = output.Output()
    data = obj.construct_data()
    print(data)
    return 'Hello World!'


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=8080)