#!/bin/sh

echo "Removing coverage files"
./find_workspace | awk '{print $1"/coverage"}' | xargs -l rm

if [ -f ./golangci-lint ]
then
    echo "Removing golangci-lint temp file"
    rm ./golangci-lint
fi

if [ -f ./find_workspace ]
then
    echo "Removing find_workspace temp file"
    rm ./find_workspace
fi

if [ -f temp_test_results ]
then
    rm temp_test_results
fi