#!/bin/bash

PROJECT_DIR=$(dirname "$0")
GENERATE_DIR="$PROJECT_DIR/cmd/gormgen"

cd "$GENERATE_DIR" || exit

echo "Start Generating"
go run .
