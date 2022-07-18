# Linter Strategy

We use [golangci-lint](https://github.com/golangci/golangci-lint) to verify if the code accomplish with the rules needed to pass

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
```

The basic strategy is run the linter inside each module specified in the go.work with the command

```bash
golangci-lint run <workspace>/...
```

# Testing Strategy

We use the next strategy based on [this article](https://blog.friendsofgo.tech/posts/tests-coverage-en-go/).
First, it is generated the file with the results in a coverage file, with the next command

```bash
go test -covermode=count -coverprofile coverage $(go list ./... | grep -v /vendor/ | tr '\n' ' ')
```

And with the next command is printed the results of the tests

```bash
go tool cover -func=coverage
```

TODO: The next action needed to do is delete the coverage file genereted

# Deploy strategy

TODO: Review this info https://docs.microsoft.com/en-us/azure/architecture/microservices/ci-cd-kubernetes