#!/bin/bash

chmod 777 start.sh

rm main
rm dashboardv2.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main . 
zip dashboardv2.zip main

echo finished