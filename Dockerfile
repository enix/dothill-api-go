FROM golang:1.12-alpine3.9

RUN apk add --update git curl g++

COPY . /app

WORKDIR /app

ENTRYPOINT [ "/bin/sh" ]

CMD [ "-c", "sleep 1 && go test" ]
