#!/bin/sh
# Copyright 2022 Daqnext Foundation Ltd.

rm -f -R ./build
mkdir build

#todo mac is for dev, not for pro
echo "Compiling MAC amd64 version"
GOOS=darwin GOARCH=amd64 go build -a -o "./build/daemon-darwin-amd64"

echo "Compiling Linux  amd64 version"
GOOS=linux GOARCH=amd64  go build -a -o "./build/daemon-linux-amd64"

echo "Compiling Linux ARM64 version"
GOOS=linux GOARCH=arm64   go build -a -o "./build/daemon-linux-arm64"

echo "Compiling Linux ARM32 version"
GOOS=linux GOARCH=arm   go build -a -o "./build/daemon-linux-arm32"