debug:
	go run -mod vendor cmd/marc-034d/main.go -nextzen-api-key $(APIKEY)

docker-build:
	@make build
	docker build -t marc-034d .

docker-debug: docker-build
	docker run -it -p 8080:8080 -e 'MAPZEN_APIKEY=$(MAPZEN_APIKEY)' marc-034d
