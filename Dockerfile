FROM golang:1.17-alpine as builder

RUN apk update && apk upgrade &&  apk add libc-dev gcc

ADD . /go-marc

RUN cd /go-marc && go build -mod vendor -o bin/marc-034d cmd/marc-034d/main.go

FROM alpine

COPY --from=builder /go-marc/bin/marc-034d /usr/local/bin/marc-034d

RUN apk update && apk upgrade \
    && apk add ca-certificates

RUN mkdir -p /usr/local/data/tiles

COPY tiles/*.db /usr/local/data/tiles/

# ENTRYPOINT ["/usr/local/bin/marc-034d"]