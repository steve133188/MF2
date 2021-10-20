# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ADD . /Users/ben/go/app

WORKDIR /Users/ben/go/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-stella-api

EXPOSE 8001

CMD [ "/docker-stella-api" ]