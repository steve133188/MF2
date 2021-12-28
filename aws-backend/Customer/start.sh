#!/bin/bash

chmod 777 start.sh

rm main
rm customer.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main . 
zip customer.zip main

echo finished