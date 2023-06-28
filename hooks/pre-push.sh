#!/bin/bash

set -euo pipefail

echo "Running pre-push hook..."

echo "Running tests..."
make test

echo "Running linters..."
make lint
