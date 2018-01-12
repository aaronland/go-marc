# go-whosonfirst-bbox

Go package for parsing bounding box strings.

## Install

You will need to have both `Go` and the `make` programs installed on your computer. Assuming you do just type:

```
make bin
```

All of this package's dependencies are bundled with the code in the `vendor` directory.

## Usage

### Simple

```
package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-bbox/parser"
	"log"
)

func main() {

	var bbox = flag.String("bbox", "", "A bounding box conforming to its scheme.")
	var scheme = flag.String("scheme", "cardinal", "A scheme describing how the bounding box is formatted. Valid options include: cardinal,marc.")
	var order = flag.String("order", "swne", "The order in which (cardinal) coordinates are defined. Valid options are: swne,wsen,nwse.")
	var separator = flag.String("separator", ",", "The string character used to seperate (cardinal) coordinates.")

	flag.Parse()

	p, err := parser.NewParser()

	if err != nil {
		log.Fatal(err)
	}

	p.Scheme = *scheme
	p.Order = *order
	p.Separator = *separator

	bb, err := p.Parse(*bbox)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(bb)
}
```

## Schemes

### cardinal

Bounding boxes defined as four cardinal coorindates.

### marc

[MARC 034 - Coded Cartographic Mathematical Data](https://www.loc.gov/marc/bibliographic/bd034.html) encoded bounding boxes.

## Order

The order in which cardinal coordinates are defined.

#### swne (South, West, North, East)

min Y, min X, max Y, max X

#### wsen (West, South, East North)

min X, min Y, max X, max Y

#### nwse (North, West, South, East)

max Y, min X, min Y, max X

## Utilities

### parse-bbox

Parse a bounding box string on the command line.

```
./bin/parse-bbox -h
Usage of ./bin/parse-bbox:
  -bbox string
    	A bounding box conforming to its scheme.
  -order string
    	The order in which (cardinal) coordinates are defined. Valid options are: swne,swen,nwse,nwes. (default "swne")
  -scheme string
    	A scheme describing how the bounding box is formatted. Valid options include: cardinal,marc. (default "cardinal")
  -separator string
    	The string character used to seperate (cardinal) coordinates. (default ",")
```

_As of this writing all results are returned as cardinal coordinates using the SWNE scheme._

#### Example

Parse a (swlon,swlat,nelon,nelat) bounding box:

```
./bin/parse-bbox -bbox '-120.683022988, 51.173603058, -59.6136767005, 83.3362128' -order wsen
51.173603 -120.683023 83.336213 -59.613677
```

Parsing a MARC bounding box:

```
./bin/parse-bbox -bbox '1#$aa$ba$dE1414646$eE1414646$f0315114$g0315114' -scheme marc
-32.083333 142.533333 32.083333 142.533333
```

### parse-bbox-server

Ask an HTTP pony to parse a bounding box string.

```
./bin/parse-bbox-server -h
Usage of ./bin/parse-bbox-server:
  -host string
    	The hostname to listen for requests on (default "localhost")
  -port int
    	The port number to listen for requests on (default 8080)
```

#### Example

```
./bin/parse-bbox-server
```

```
$> python
>>> import requests
>>> import json
>>> url = "http://localhost:8080"
>>> params = { "bbox": "1#$aa$ba$dE1414646$eE1414646$f0315114$g0315114", "scheme": "marc"}
>>> rsp = requests.get(url, params=params)
>>> data = json.loads(rsp.content)
>>> print data
{u'min_x': 142.53333333333333, u'min_y': -32.083333333333336, u'max_x': 142.53333333333333, u'max_y': 32.083333333333336}
```

## See also

* https://github.com/thisisaaronland/go-marc
