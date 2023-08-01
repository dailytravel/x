#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Please enter service names as arguments"
  exit 1
fi

for s in "$@"; do
  if [[ "$s" != "" ]]; then
    mkdir -p "$s"
    cp ./tools.go "$s/tools.go"
    cp ./gen.sh "$s/gen.sh"
    cd "$s"
    go mod init github.com/dailytravel/x/"$s"
    go mod tidy
    go run github.com/99designs/gqlgen init
    go run github.com/99designs/gqlgen generate
    cd ../
    echo "Setup for $s is done!"
  else
    echo "Service name cannot be empty. Skipping..."
  fi
done
