#!/usr/bin/env bash

if [ $# -eq 0 ]; then
  echo "Please enter service names as arguments"
  exit 1
fi

for s in "$@"; do
  if [ -n "$s" ]; then
    protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/$s/$s.proto
  else
    echo "Service name cannot be empty. Skipping..."
  fi
done
