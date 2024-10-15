#!/bin/bash
# if no $1 make it empty string
if [ -z "$1" ]; then
  set -- ""
fi

# Execute docker compose logs
docker compose -f docker-compose.local.yaml logs -f "$1"