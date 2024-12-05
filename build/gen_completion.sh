#!/bin/bash
#
# Create shell completions for different shells

VERSION="$1"

cd ..

echo "creating completions directory"
mkdir -p bin/${VERSION}/completions

for shell in bash zsh fish powershell; do
    echo "creating completion for: ${shell}"
    go run . completion "${shell}" > "bin/${VERSION}/completions/vsps_${VERSION}.${shell}"
done

