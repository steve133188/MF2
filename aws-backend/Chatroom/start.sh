#!/bin/bash

chmod 777 start.sh

rm main
rm chatroom.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main . 
zip chatroom.zip main

echo finished