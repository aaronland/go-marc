GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

MAP_PROVIDER=protomaps
MAP_TILE_URL=file:///usr/local/data/pmtiles/20240415.pmtiles

SPATIAL_DATABASE_URI=rtree:///?strict=false&index_alt_files=0
SPATIAL_DATABASE_SOURCE=/usr/local/data/sfomuseum-data-whosonfirst

cli:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/marc-034 cmd/marc-034/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/marc-034d cmd/marc-034d/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/marc-034-convert cmd/marc-034-convert/main.go

debug:
	go run -mod $(GOMOD) -ldflags="$(LDFLAGS)" \
		cmd/marc-034d/main.go \
		-verbose \
		-map-provider $(MAP_PROVIDER) \
		-map-tile-uri $(MAP_TILE_URL) 

debug-intersects:
	go run -mod $(GOMOD) -ldflags="$(LDFLAGS)" \
		cmd/marc-034d/main.go \
		-verbose \
		-map-provider $(MAP_PROVIDER) \
		-map-tile-uri $(MAP_TILE_URL) \
		-enable-intersects \
		-spatial-database-uri '$(SPATIAL_DATABASE_URI)' \
		-spatial-database-source 'repo://#$(SPATIAL_DATABASE_SOURCE)'
