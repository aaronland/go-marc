CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/thisisaaronland/go-marc; then rm -rf src/github.com/thisisaaronland/go-marc; fi
	mkdir -p src/github.com/thisisaaronland/go-marc/fields
	cp fields/*.go src/github.com/thisisaaronland/go-marc/fields
	cp marc.go src/github.com/thisisaaronland/go-marc/
	# cp -r vendor/src/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

deps:
	echo "no deps"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor/src; then rm -rf vendor/src; fi
	cp -r src vendor/src
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt fields/*.go
	go fmt *.go

bin: 	rmdeps self
	@GOPATH=$(GOPATH) go build -o bin/034-to-bbox cmd/034-to-bbox.go
