|❗ Warning |
|:----------|
|As of 2023/01, this project is no longer maintained by Enix and Seagate as take over the project. We archived the repository and we invite you to join this project [Seagate/seagate-exos-x-csi](https://github.com/Seagate/seagate-exos-x-csi).|

# dothill-api-go

[![Build status](https://gitlab.com/enix.io/dothill-api-go/badges/master/pipeline.svg)](https://gitlab.com/enix.io/dothill-api-go/-/pipelines)
[![Go Report Card](https://goreportcard.com/badge/github.com/enix/dothill-api-go/v2)](https://goreportcard.com/report/github.com/enix/dothill-api-go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/enix/dothill-api-go/v2)](https://pkg.go.dev/github.com/enix/dothill-api-go/v2)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)

A Go implementation of the [Dothill API](https://www.seagate.com/files/dothill-content/support/documentation/AssuredSAN_4004_Series_CLI_Reference_Guide_GL105.pdf).

## Run tests using our mock server

In order to run tests, you will need to start a mock server by running `go run cmd/mock/mock.go`. It will expose the mocked api on `localhost:8080`.

You're now ready to go, just run `go test -v` to run the tests suite.

You can also run tests with docker-compose, using this command, which is the one used in the CI:

```sh
docker-compose up --build --abort-on-container-exit --exit-code-from tests
```

Both the mocked api and the tests use as username and passwords variables from the environment, or from the `.env` file if missing. Since the `.env` file is pre-filled, you should not have to add any environment variable in order to make the tests work.

## Test using a live system

This option runs the golang test cases against a live storage system. Two steps are required:

- Update .env with the correct system IP address and credentials
- Run `go test -v`

Another option is to define environment variables, which take precedence over .env values

```bash
export TEST_STORAGEIP=http://<ipaddress>
export TEST_USERNAME=<username>
export TEST_PASSWORD=<password>
go test -v
unset TEST_STORAGEIP TEST_PASSWORD TEST_USERNAME
```
