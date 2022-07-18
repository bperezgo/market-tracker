#!/bin/sh
# Remember change the persmissions of this file

echo "Setting up the environment"
chmod +x ./pipelines/**/*.sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2