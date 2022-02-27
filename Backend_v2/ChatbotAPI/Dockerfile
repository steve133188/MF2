FROM golang:1.16-alpine

ENV MEMORYDBURL "clustercfg.mf2-redis.2j4s5t.memorydb.ap-east-1.amazonaws.com:6379"

ENV CHATBOTDBURL "0.0.0.0:6379"

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /channel-router

EXPOSE 8080

CMD [ "/channel-router" ]