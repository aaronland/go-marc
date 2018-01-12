# https://blog.docker.com/2016/09/docker-golang/
# https://blog.golang.org/docker

# docker build -t 034d .
# docker run -it -p 8080:8080 034d

FROM golang:alpine AS build-env

RUN apk add --update alpine-sdk

ADD . /go-marc

RUN cd /go-marc; make bin

FROM alpine

COPY --from=build-env /go-marc/bin/marc-034d /marc-034d

EXPOSE 8080

ENTRYPOINT /marc-034d

