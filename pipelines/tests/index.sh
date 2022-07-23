#!/bin/sh

go build -o ./find_workspace ./pipelines/find_workspace/main.go
./find_workspace | awk '{print $1"/... -covermode=count -coverprofile "$1"/coverage $(go list ./... | grep -v /vendor/)"}' | xargs -l go test
./find_workspace | awk '{print " -func="$1"/coverage"}' | xargs -l  go tool cover
# ./find_workspace | awk '{print $1"/coverage"}' | xargs -l rm