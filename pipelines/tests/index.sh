#!/bin/sh

go build -o ./find_workspace ./pipelines/find_workspace/main.go
if [ -f temp_test_results ]
then
    rm temp_test_results
    touch temp_test_results
else
    touch temp_test_results
fi

./find_workspace | awk '{print $1"/... -covermode=count -coverprofile "$1"/coverage $(go list ./... | grep -v /vendor/)"}' | xargs -l go test -v | grep -i '^FAIL$' >> temp_test_results
./find_workspace | awk '{print " -func="$1"/coverage"}' | xargs -l  go tool cover

if [ -s temp_test_results ]
then
    echo "Some tests FAILED"
    exit 1
fi
