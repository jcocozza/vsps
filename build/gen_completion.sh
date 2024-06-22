#!/bin/bash
#
# Create shell completions for different shells
cd ..

mkdir -p bin/completions

for shell in bash zsh fish powershell; do
    go run . completion "${shell}" > "bin/completions/vsps.${shell}"
done

