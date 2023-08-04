#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Please enter service names as arguments"
  exit 1
fi

for s in "$@"; do
  if [[ "$s" != "" ]]; then
    cd "$s"
    ./gen.sh
    cd ../
    echo "Gen for $s is done!"
  else
    echo "Service name cannot be empty. Skipping..."
  fi
done
