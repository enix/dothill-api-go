FROM golang:1.12-alpine3.9

WORKDIR /go/src/enix.io/dothill-api-go

RUN apk add --update curl g++

COPY *.go ./

COPY test ./test

WORKDIR /go/src/enix.io/dothill-api-go/test

ENTRYPOINT [ "/bin/sh" ]

CMD [ "-c", "sleep 1 && go test" ]
