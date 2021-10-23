cli:
	go build -mod vendor -o bin/marc-034 cmd/marc-034/main.go
	go build -mod vendor -o bin/marc-034d cmd/marc-034d/main.go

debug:
	go run -mod vendor cmd/marc-034d/main.go -nextzen-api-key $(APIKEY)

debug-tilepack:
	go run -mod vendor cmd/marc-034d/main.go -nextzen-tilepack-database $(TILEPACK)

docker:
	docker build -t marc-034d .

docker-debug:
	docker run -it -p 8080:8080 -e 'MARC_NEXTZEN_APIKEY=$(APIKEY)' marc-034d
