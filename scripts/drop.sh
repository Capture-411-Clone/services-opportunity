#!/bin/bash

# Get current directory
CURR_DIR=$(dirname "$(realpath "${0}")")

# Load environment variables
. "$CURR_DIR"/base/init.sh

# Execute docker compose down
docker compose -f \
  "$PROJECT_ROOT_DIR"/docker-compose.yaml \
  --env-file "$PROJECT_ROOT_DIR"/.env.production \
  down --remove-orphans