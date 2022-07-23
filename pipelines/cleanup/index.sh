#!/bin/sh

echo "Removing coverage files"
./find_workspace | awk '{print $1"/coverage"}' | xargs -l rm

echo "Removing golangci-lint temp file"
rm ./golangci-lint

echo "Removing find_workspace temp file"
rm ./find_workspace