#!/bin/bash

chmod 777 start.sh

rm main
rm user_without_auth.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main . 
zip user_without_auth.zip main

echo finished