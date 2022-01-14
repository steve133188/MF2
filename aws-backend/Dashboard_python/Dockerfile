FROM python:3.7.2-stretch

WORKDIR /app

ADD . /app

RUN pip3 install -r requirements.txt

EXPOSE 8080

CMD python app.py