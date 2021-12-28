#!/bin/bash

chmod 777 start.sh

rm main
rm tag.zip

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main . 
zip tag.zip main

echo finished