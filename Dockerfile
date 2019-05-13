FROM instrumentisto/dep:0.5-alpine

WORKDIR /go/src/enix.io/dothill-api-go

RUN apk add --update curl

COPY *.go ./

COPY test ./test

RUN dep init

WORKDIR /go/src/enix.io/dothill-api-go/test

ENTRYPOINT [ "/bin/sh" ]

CMD [ "-c", "sleep 1 && go test" ]
