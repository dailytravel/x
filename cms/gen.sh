#!/bin/bash

clear

go mod tidy
go run github.com/99designs/gqlgen generate
echo "Generated GraphQL files"