#!/bin/sh

## TODO: Use all the modules in the workspaces
echo "Linking linter to workdirectory"
ln -s $(go env GOPATH)/bin/golangci-lint .
./golangci-lint --version

echo "Running the linters in the golang project"

go build -o ./find_workspace ./pipelines/find_workspace/main.go
./find_workspace | awk '{print $1"/..."}' | xargs -l ./golangci-lint run
