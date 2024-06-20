#!/bin/bash
#
# Build the cli, gui is currently not supported. 
# Run in the build folder
# Writes to the bin folder

cd ..

# CLI BUILD
cli="vsps_cli"
# Linux
GOOS=linux GOARCH=amd64 go build -o "${cli}_Linux_amd64"
GOOS=linux GOARCH=arm64 go build -o "${cli}_Linux_aarch64"
# MacOS
GOOS=darwin GOARCH=amd64 go build -o "${cli}_Darwin_amd64"
GOOS=darwin GOARCH=arm64 go build -o "${cli}_Darwin_arm64"
# Windows
GOOS=windows GOARCH=amd64 go build -o "${cli}_Windows_amd64"

mkdir -p bin
mv "${cli}_"* "bin/"
