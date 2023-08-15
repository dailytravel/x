#!/bin/bash

echo "Generating GraphQL files in hrm"
go mod tidy
go run github.com/99designs/gqlgen generate
echo "Done!"