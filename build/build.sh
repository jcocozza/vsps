#!/bin/bash
#
# Build the cli, gui is currently not supported.
# Run in the build folder
# Writes to the bin folder

VERSION="$1"
if [[ -z "${VERSION}" ]]; then
    echo "version required"
    exit 1
fi

CHECKVERSION=$(cat ../cmd/root.go | grep "const version" | sed 's/.*= "\(.*\)"/\1/')

if [[ ${VERSION} != ${CHECKVERSION} ]]; then
    echo "version mismatch."
    echo "cmd/root.go has version: ${CHECKVERSION}"
    echo "you passed: ${VERSION}"
    exit 1
fi

echo "creating builds for version: ${VERSION}"

cd ..

# CLI BUILD
cli="vsps_cli"
# Linux
GOOS=linux GOARCH=amd64 go build -o "${cli}_linux_amd64_${VERSION}"
GOOS=linux GOARCH=arm64 go build -o "${cli}_linux_aarch64_${VERSION}"
# MacOS
GOOS=darwin GOARCH=amd64 go build -o "${cli}_darwin_amd64_${VERSION}"
GOOS=darwin GOARCH=arm64 go build -o "${cli}_darwin_arm64_${VERSION}"
# Windows
GOOS=windows GOARCH=amd64 go build -o "${cli}_windows_amd64_${VERSION}"

BINDIR="bin/${VERSION}"
# Check if the build files exist
 if ls ${cli}_* 1> /dev/null 2>&1; then
    BINDIR="bin/${VERSION}"
    mkdir -p "${BINDIR}"
    mv ${cli}_* "${BINDIR}"
    echo "Builds moved to ${BINDIR}"
else
    echo "Build files not found. Something went wrong during the build process."
    exit 1
fi

# Create the completion scripts for different shells
cd build
bash gen_completion.sh "${VERSION}"
