GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

MAP_PROVIDER=protomaps
MAP_TILE_URL=file:///usr/local/data/pmtiles/20240415.pmtiles

SPATIAL_DATABASE_URI=rtree:///?strict=false&index_alt_files=0
SPATIAL_DATABASE_SOURCE=/usr/local/data/sfomuseum-data-whosonfirst

cli:
	rm -rf bin/*
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/parse cmd/parse/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/server cmd/server/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/convert cmd/convert/main.go

server:
	go run -mod $(GOMOD) -ldflags="$(LDFLAGS)" \
		cmd/server/main.go \
		-verbose \
		-map-provider $(MAP_PROVIDER) \
		-map-tile-uri $(MAP_TILE_URL) 

server-intersects:
	go run -mod $(GOMOD) -ldflags="$(LDFLAGS)" \
		cmd/server/main.go \
		-map-provider $(MAP_PROVIDER) \
		-map-tile-uri $(MAP_TILE_URL) \
		-enable-intersects \
		-spatial-database-uri '$(SPATIAL_DATABASE_URI)' \
		-spatial-database-source 'repo://#$(SPATIAL_DATABASE_SOURCE)'

convert-intersects:
	go run -mod $(GOMOD) -ldflags="$(LDFLAGS)" \
		cmd/convert/main.go \
		-enable-intersects \
		-spatial-database-uri '$(SPATIAL_DATABASE_URI)' \
		-spatial-database-source 'repo://#$(SPATIAL_DATABASE_SOURCE)' \
		fixtures/marc034-intersects.csv
