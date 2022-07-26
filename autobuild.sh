#!/bin/sh
# Copyright 2022 Daqnext Foundation Ltd.

rm -f -R ./build
mkdir build

#todo mac is for dev, not for pro
echo "Compiling MAC 64bit version"
GOOS=darwin GOARCH=amd64 go build -a -o "./build/service-darwin-amd64"

echo "Compiling Linux 64bit version"
GOOS=linux GOARCH=amd64  go build -a -o "./build/service-linux-amd64"

echo "Compiling ARM64 version"
GOOS=linux GOARCH=arm64   go build -a -o "./build/service-linux-arm64"

echo "Compiling ARM version"
GOOS=linux GOARCH=arm   go build -a -o "./build/service-linux-arm"