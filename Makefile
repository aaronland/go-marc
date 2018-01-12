CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src/github.com/thisisaaronland/go-marc; then rm -rf src/github.com/thisisaaronland/go-marc; fi
	mkdir -p src/github.com/thisisaaronland/go-marc/fields
	cp -r assets src/github.com/thisisaaronland/go-marc/
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
	@GOPATH=$(GOPATH) go get -u "github.com/jteeuwen/go-bindata/"
	@GOPATH=$(GOPATH) go get -u "github.com/elazarl/go-bindata-assetfs/"
	rm -rf src/github.com/jteeuwen/go-bindata/testdata

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	cp -r src/* vendor/
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

assets: self
	@GOPATH=$(GOPATH) go build -o bin/go-bindata ./vendor/github.com/jteeuwen/go-bindata/go-bindata/
	rm -rf templates/*/*~
	rm -rf assets
	mkdir -p assets/html
	@GOPATH=$(GOPATH) bin/go-bindata -pkg html -o assets/html/html.go templates/html

static: self
	@GOPATH=$(GOPATH) go build -o bin/go-bindata ./vendor/github.com/jteeuwen/go-bindata/go-bindata/
	@GOPATH=$(GOPATH) go build -o bin/go-bindata-assetfs vendor/github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs/main.go
	rm -f static/css/*~ static/javascript/*~ static/tangram/*~ static/fonts/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg http static/javascript static/css
	if test -f http/static_fs.go; then rm http/static_fs.go; fi
	mv bindata_assetfs.go http/static_fs.go

build:
	@make assets
	@make static
	@make bin

debug:
	@make build
	bin/marc-034d

fmt:
	go fmt cmd/*.go
	go fmt fields/*.go
	go fmt http/*.go
	go fmt *.go

bin: 	rmdeps self
	rm -rf bin/*
	@GOPATH=$(GOPATH) go build -o bin/marc-034 cmd/marc-034.go
	@GOPATH=$(GOPATH) go build -o bin/marc-034d cmd/marc-034d.go

docker-build:
	docker build -t marc-034d .

docker-debug: docker-build
	docker run -it -p 8080:8080 marc-034d
