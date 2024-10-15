#!/bin/bash

# Set project root directory
ROOT=$(realpath "$(dirname "${0}")"/..)

# Set environment variables
export GITLAB_URI=gitag.ir
export PROJECT_ROOT_DIR=$ROOT
export SCRIPTS_DIR=$ROOT/scripts
export LOCAL_BIN_DIR=$ROOT/bin
export DOCKER_EXTERNAL_NETWORK=capture411Bridge

# Create bin directory if not exists
mkdir -p "$PROJECT_ROOT_DIR"/bin