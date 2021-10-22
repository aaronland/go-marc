FROM golang:1.17-alpine as builder

RUN apk update && apk upgrade

ADD . /go-marc

RUN cd /go-marc && go build -mod vendor -o bin/marc-034d cmd/marc-034d/main.go

FROM alpine

COPY --from=builder /go-marc/bin/marc-034d /usr/local/bin/marc-034d

RUN apk update && apk upgrade \
    && apk add ca-certificates

ENTRYPOINT ["/usr/local/bin/marc-034d"]