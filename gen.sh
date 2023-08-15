#!/bin/bash

echo "Generating..."
go mod tidy
go run github.com/99designs/gqlgen generate
echo "Done!"