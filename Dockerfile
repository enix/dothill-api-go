FROM instrumentisto/dep:0.5-alpine

WORKDIR /go/src/enix.io/dothill-api-go

COPY *.go ./

RUN dep init

RUN go test
