#!/bin/bash

# Get current directory
CURR_DIR=$(dirname "$(realpath "${0}")")

# Load environment variables
. "$CURR_DIR"/base/init.sh

cp -rf "$PROJECT_ROOT_DIR"/scripts/hooks/pre-commit.sh "$PROJECT_ROOT_DIR"/.git/hooks/pre-commit

echo "git hooks prepared"