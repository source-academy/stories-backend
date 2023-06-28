#!/bin/bash

set -euo pipefail

echo "Running pre-commit hook..."

if [[ ! -z `gofmt -l .` ]]; then
    echo "Go code is not formatted, commit blocked!"
    echo "The following files are not formatted:"
    echo
    gofmt -l . | sed 's/^/.\//'
    echo
    echo "Run 'make format' to format your files and try again."
    exit 1
fi
