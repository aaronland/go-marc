# go-marc

Go package for working with MARC records.

## Important

Not all of MARC. Probably not ever. Just the `034` field so far. If you are looking for a general-purpose library for working with MARC records I'd recommend looking at [miku/marc21](https://github.com/miku/marc21).

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/aaronland/go-marc.svg)](https://pkg.go.dev/github.com/aaronland/go-marc)

## Tools

```
$> make cli

rm -rf bin/*
go build -mod vendor -ldflags="-s -w" -o bin/parse cmd/parse/main.go
go build -mod vendor -ldflags="-s -w" -o bin/server cmd/server/main.go
go build -mod vendor -ldflags="-s -w" -o bin/convert cmd/convert/main.go
```

### parse

Parse one or more MARC 034 strings and emit a (S, W, N, E) bounding box for each.

Documentation for the `parse` tool can be found in [cmd/parse/README.md](cmd/parse/README.md).

### convert

Process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.

Documentation for the `parse` tool can be found in [cmd/convert/README.md](cmd/convert/README.md).

### server

A web application for converting MARC 034 strings in to bounding boxes and/or batch processing CSV files uploaded to the server.

Documentation for the `server` tool can be found in [cmd/server/README.md](cmd/server/README.md).

## See also

* https://www.loc.gov/marc/bibliographic/bd034.html
