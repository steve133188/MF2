from flask import Flask
# from output import Output

app = Flask(__name__)


@app.route('/')
def hello_world():  # put application's code here
    # data = Output.construct_data()
    # print(data)
    return 'Hello World!'


if __name__ == '__main__':
    app.run(debug=True)