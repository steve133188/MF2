#!/bin/bash

chmod 777 start.sh

rm main
rm dashboard.zip
rm dashboardreq.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main . 
zip dashboardreq.zip main

echo finished