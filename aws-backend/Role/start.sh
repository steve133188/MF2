#!/bin/bash

chmod 777 start.sh

rm main
rm role.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main . 
zip role.zip main

echo finished