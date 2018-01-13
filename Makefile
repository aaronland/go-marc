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
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-http-mapzenjs"
	rm -rf src/github.com/jteeuwen/go-bindata/testdata

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	cp -r src/* vendor/
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

maps: tangram refill mapzenjs

tangram:
	if test ! -d static/tangram; then mkdir -p static/tangram; fi
	curl -s -o static/javascript/tangram.js https://mapzen.com/tangram/tangram.debug.js
	curl -s -o static/javascript/tangram.min.js https://mapzen.com/tangram/tangram.min.js

refill:
	if test ! -d static/tangram; then mkdir -p static/tangram; fi
	curl -s -o static/tangram/refill-style.zip https://mapzen.com/carto/refill-style/refill-style.zip

mapzenjs:
	if test ! -d static/javascript; then mkdir -p static/javascript; fi
	if test ! -d static/css; then mkdir -p static/css; fi
	curl -s -o static/css/mapzen.js.css https://mapzen.com/js/mapzen.css
	curl -s -o static/javascript/mapzen.js https://mapzen.com/js/mapzen.js
	curl -s -o static/javascript/mapzen.min.js https://mapzen.com/js/mapzen.min.js

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
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg http static/javascript static/css static/images
	if test -f http/static_fs.go; then rm http/static_fs.go; fi
	mv bindata_assetfs.go http/static_fs.go

leaflet:
	if test -d tmp; then rm -rf tmp; fi
	mkdir tmp
	curl -s -o tmp/leaflet.zip http://cdn.leafletjs.com/leaflet/v1.2.0/leaflet.zip
	unzip -d tmp tmp/leaflet.zip
	mv tmp/leaflet.js static/javascript/
	mv tmp/leaflet-src*.js static/javascript/
	mv tmp/leaflet*.css static/css/
	mv tmp/images/*.png static/images/
	rm -rf tmp

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
