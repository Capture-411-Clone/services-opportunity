#!/bin/bash

if [ ! -x "$(command -v go)" ]; then
  echo "Go is not installed!"
else
  go mod download &&
    go mod vendor
fi
