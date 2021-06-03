# dothill-api-go

[![Build status](https://gitlab.com/enix.io/dothill-api-go/badges/master/pipeline.svg)](https://gitlab.com/enix.io/dothill-api-go/-/pipelines)
[![Go Report Card](https://goreportcard.com/badge/github.com/enix/dothill-api-go)](https://goreportcard.com/report/github.com/enix/dothill-api-go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/enix/dothill-api-go)](https://pkg.go.dev/github.com/enix/dothill-api-go)

A Go implementation of the [Dothill API](https://www.seagate.com/files/dothill-content/support/documentation/AssuredSAN_4004_Series_CLI_Reference_Guide_GL105.pdf).

## Run tests

In order to run tests, you will need to start a mock server by running `go run cmd/mock/mock.go`. It will expose the mocked api on `localhost:8080`.

You're now ready to go, just run `go test` to run the tests suite.

You can also run tests with docker-compose:

```sh
docker-compose up --build --abort-on-container-exit --exit-code-from tests
```

## Test Using A Live System

This option runs the Go language test cases against a live storage system. Two steps are required:
- Update .env with the correct system IP Address and credentials
- Run `go test -v`

Another option is to define environment variables, which take precedence over .env values
- export TEST_STORAGEIP=http://<ipaddress>
- export TEST_USERNAME=<username>
- export TEST_PASSWORD=<password>
- Run `go test -v`
- unset TEST_STORAGEIP TEST_PASSWORD TEST_USERNAME
