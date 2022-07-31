#!/bin/sh

# this file is only used to initialize the environment for local development

# Remember change the persmissions of this file
chmod +x ./pipelines/**/*.sh
pre-commit install
pre-commit install --hook-type post-commit
# TODO: Used to initialize the dev project, and work with the dev config in local
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.46.2
golangci-lint --version
