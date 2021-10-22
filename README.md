# go-marc

Go package for working with MARC records.

## Important

Not all of MARC. Probably not ever. Just the `034` field so far. If you are looking for a general-purpose library for working with MARC records I'd recommend looking at [miku/marc21](https://github.com/miku/marc21).

## Tools

```
$> make cli
go build -mod vendor -o bin/marc-034 cmd/marc-034/main.go
go build -mod vendor -o bin/marc-034d cmd/marc-034d/main.go
```

### marc-034

Convert a MARC 034 string in to a (S, W, N, E) bounding box.

```
$> ./bin/marc-034 -h
Usage of ./bin/marc-034:
  -f string
    	A valid MARC 034 string (default "1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000")
```

Currently this only supports `hdddmmss (hemisphere-degrees-minutes-seconds)` and `dddmmss (degrees-minutes-seconds)` notation. For example:

```
$> ./bin/marc-034
2017/02/13 22:23:38 1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000 <-- input (MARC 034)
2017/02/13 22:23:38 -70.000000, -180.000000 84.000000, 180.000000 <-- output (decimal WSG84)
```

### marc-034d

A web server for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON)

```
$> ./bin/marc-034d -h
Usage of ./bin/marc-034d:
  -nextzen-api-key string
    	A valid Nextzen API key (default "mapzen-xxxxxx")
  -nextzen-style-url string
    	A valid Nextzen style URL (default "/tangram/refill-style.zip")
  -server-uri string
    	A valid aaronland/go-http-server URI (default "http://localhost:8080")
```

For example:

```
$> ./bin/marc-034d -nextzen-api-key {APIKEY}

2018/01/12 09:12:44 listening on localhost:8080
```

The `marc-034d` server exposes the following endpoints:

#### / (or "root")

The `/` (or default) endpoint will display a handy web interface for converting MARC 034 records in to bounding boxes. For example, here's what it looks like querying for `1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000`:

![](docs/images/marc-034d-www-v2.png)

#### /bbox

The `/bbox` endpoint will return a bounding box for a MARC 034 field as GeoJSON.

```
$> curl -s 'http://localhost:8080/bbox?034=1%23%24aa$b22000000%24dW1800000%24eE1800000%24fN0840000%24gS0700000' | python -mjson.tool

{
    "bbox": [
        -180,
        -70,
        180,
        84
    ],
    "geometry": {
        "coordinates": [
            [
                [
                    -180,
                    -70
                ],
                [
                    -180,
                    84
                ],
                [
                    180,
                    84
                ],
                [
                    180,
                    -70
                ],
                [
                    -180,
                    -70
                ]
            ]
        ],
        "type": "Polygon"
    },
    "properties": {
        "marc:034": "1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000"
    },
    "type": "Feature"
}
```

_Note the way the `034` parameter is URL-encoded._

### Command-line flags and environment variables

Command line flags can be set also be set from environment variables. Environment variables for any given command line flag should be formatted as follows:

* Replace all `-` characters with `_`
* Upper case the flag name
* Prepend the string with `MARC_`

For example the equivalent environment variable for the `nextzen-api-key` flag would be `MARC_NEXTZEN_API_KEY`.

### Nextzen and Nextzen API keys

You can register for a Nextzen API key from [https://developers.nextzen.org/](https://developers.nextzen.org/).

## Docker

[Yes](Docker), for `marc-034d` at least.

```
$> docker build -t marc-034d .

$> docker run -it -p 8080:8080 marc-034d -server-uri http://0.0.0.0:8080 -nextzen-api-key {APIKEY} 
```

## See also

* https://www.loc.gov/marc/bibliographic/bd034.html
* https://github.com/aaronland/go-http-tangramjs
* https://github.com/aaronland/go-http-bootstrap
* https://github.com/aaronland/go-http-server
* https://developers.nextzen.org/
