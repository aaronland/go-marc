CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-bbox/parser
	cp bbox.go src/github.com/whosonfirst/go-whosonfirst-bbox/
	cp parser/*.go src/github.com/whosonfirst/go-whosonfirst-bbox/parser
	cp -r vendor/src/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/facebookgo/grace/gracehttp"
	@GOPATH=$(GOPATH) go get -u "github.com/thisisaaronland/go-marc"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt bbox.go
	go fmt cmd/*.go

bin: 	rmdeps self
	@GOPATH=$(GOPATH) go build -o bin/parse-bbox cmd/parse-bbox.go
	@GOPATH=$(GOPATH) go build -o bin/parse-bbox-server cmd/parse-bbox-server.go
