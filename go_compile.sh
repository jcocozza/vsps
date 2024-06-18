#!/bin/bash
#
# Simple script to build compile golang applications

app_name="$1"

# Linux
GOOS=linux GOARCH=amd64 go build -o "$1_Linux_amd64"
GOOS=linux GOARCH=arm64 go build -o "$1_Linux_aarch64"
# MacOS
GOOS=darwin GOARCH=amd64 go build -o "$1_Darwin_amd64"
GOOS=darwin GOARCH=arm64 go build -o "$1_Darwin_arm64"

# Windows
GOOS=windows GOARCH=amd64 go build -o "$1_Windows_amd64"

mkdir -p bin
mv "${app_name}_"* "bin/"

