#!/bin/bash
#
# Build the cli, gui is currently not supported. 
# Run in the build folder
# Writes to the bin folder

cd ..

# clear bin directory
rm -rf bin

# CLI BUILD
cli="vsps_cli"
# Linux
GOOS=linux GOARCH=amd64 go build -o "${cli}_linux_amd64"
GOOS=linux GOARCH=arm64 go build -o "${cli}_linux_aarch64"
# MacOS
GOOS=darwin GOARCH=amd64 go build -o "${cli}_darwin_amd64"
GOOS=darwin GOARCH=arm64 go build -o "${cli}_darwin_arm64"
# Windows
GOOS=windows GOARCH=amd64 go build -o "${cli}_windows_amd64"

mkdir -p bin
mv "${cli}_"* "bin/"

# Create the completion scripts for differnt shells
cd build
bash gen_completion.sh
