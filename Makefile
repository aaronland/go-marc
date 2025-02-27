GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/marc-034 cmd/marc-034/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/marc-034d cmd/marc-034d/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/marc-034-convert cmd/marc-034-convert/main.go

debug:
	go run -mod $(GOMOD) -ldflags="$(LDFLAGS)" cmd/marc-034d/main.go

docker:
	docker build -t marc-034d .

docker-debug:
	docker run -it -p 8080:8080 -e 'MARC_NEXTZEN_APIKEY=$(APIKEY)' marc-034d
