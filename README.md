# dothill-api-go

[![Build status](https://gitlab.com/enix.io/dothill-api-go/badges/master/pipeline.svg)](https://gitlab.com/enix.io/dothill-api-go/-/pipelines)
[![Go Report Card](https://goreportcard.com/badge/github.com/enix/dothill-api-go)](https://goreportcard.com/report/github.com/enix/dothill-api-go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/enix/dothill-api-go)](https://pkg.go.dev/github.com/enix/dothill-api-go)

A Go implementation of the [Dothill API](https://www.seagate.com/files/dothill-content/support/documentation/AssuredSAN_4004_Series_CLI_Reference_Guide_GL105.pdf).

## Run tests

In order to run tests, you will need to install node.js and npm to run the mock server. When it's done, go to the `mock` directory, install dependencies and start the mock server.

```sh
cd ./mock
npm install
npm run start
```

You're now ready to go, just run `go test` to run the tests suite.

You can also skip previous steps and just run tests with docker-compose:

```sh
docker-compose up --build --abort-on-container-exit --exit-code-from tests
```
