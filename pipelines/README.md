# Linter Strategy

TODO:

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