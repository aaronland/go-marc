CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/thisisaaronland/go-marc; then rm -rf src/github.com/thisisaaronland/go-marc; fi
	mkdir -p src/github.com/thisisaaronland/go-marc/fields
	cp -r fields src/github.com/thisisaaronland/go-marc/
	cp -r http src/github.com/thisisaaronland/go-marc/
	cp marc.go src/github.com/thisisaaronland/go-marc/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-sanitize"	
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-bbox"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	cp -r src/* vendor/
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt fields/*.go
	go fmt http/*.go
	go fmt *.go

bin: 	rmdeps self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/marc-034 cmd/marc-034.go
	@GOPATH=$(GOPATH) go build -o bin/marc-034d cmd/marc-034d.go
