#!/bin/bash

# Check .gitconfig file exists or not and create it if not exists
if [ ! -f "$PROJECT_ROOT_DIR"/.gitconfig ]; then
  echo ".gitconfig file not found, creating..."

  if [[ "$GITLAB_CI_USERNAME" == "" ]]; then
    read -rp "Gitlab Username: " GITLAB_CI_USERNAME
  fi
  if [[ "$GITLAB_CI_PAT" == "" ]]; then
    read -rp "Personal Access Token(PAT): " GITLAB_CI_PAT
  fi

  printf "[url \"https://%s:%s@%s\"]\n\tinsteadOf = https://gitlab.hoitek.fi\n" "$GITLAB_CI_USERNAME" "$GITLAB_CI_PAT" "$GITLAB_URI" > "$PROJECT_ROOT_DIR"/.gitconfig
fi
