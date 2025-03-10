# convert

Process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.

```
$> ./bin/convert -h
Process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.
Usage:
	 ./bin/convert csv-file(N) csv-file(N)
  -enable-intersects
    	Enable intersecting geometry lookups for MARC034-derived bounding boxes.
  -marc-034-column string
    	The name of the CSV column where MARC 034 data is stored. (default "marc_034")
  -spatial-database-source value
    	Zero or more '{ITERATOR_URI}#{ITERATOR_SOURCE}' strings following the whosonfirst/go-whosonfirst-iterate/v2 URI syntax.
  -spatial-database-uri string
    	A registered whosonfirst/go-whosonfirst-spatial/database.SpatialDatabase URI.
  -to-file string
    	The path where your new CSV file should be created.
  -to-stdout
    	Output CSV data to STDOUT. (default true)
  -verbose
    	Enable verbose (debug) logging.
```

## Example

For example, given in an input CSV file that looks this:

```
$> cat fixtures/marc034.csv
id,marc_034,name
123,1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000,example
456,1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000,another example
```

Passing it to the `marc-034-convert` tool would yield:

```
$> ./bin/convert -to-stdout ./fixtures/marc034.csv
error,id,marc_034,max_x,max_y,min_x,min_y,name,valid,intersects
,123,1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000,180,84,-180,-70,example,1,
,456,1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000,-82.33333333333333,27.5,-82.91666666666667,26.833333333333332,another example,1,
```

## Intersecting (Who's On First) geometries

There is optional support for retrieving [Who's On First](https://whosonfirst.org) (WOF) records whose geometries intersect a bounding box derived from a MARC 034 record. For example:

```
$> ./bin/convert \
	-enable-intersects \
	-spatial-database-uri 'rtree:///?strict=false&index_alt_files=0' \
	-spatial-database-source 'repo://#/usr/local/data/sfomuseum-data-whosonfirst' \
	fixtures/marc034-intersects.csv
2025/03/10 11:53:25 INFO Indexing spatial database.
2025/03/10 11:53:53 INFO time to index paths (1) 28.397383917s
```

Would yield results like this:

```
error,id,intersects,marc_034,max_x,max_y,min_x,min_y,name,valid
,456,"wof:planet=0,wof:continent=102191575,wof:country=85633793,wof:region=85688651,wof:campus=102527435",1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000,-82.33333333333333,27.5,-82.91666666666667,26.833333333333332,another example,1
```

Intersecting WOF records are recorded in the `intersects` column as a (comma-separated) list of [machine tags](https://web.archive.org/web/20160420154054/https://www.flickr.com/groups/api/discuss/72157594497877875/) in the form of:

"wof:" + `{WHOSONFIRST_PLACETYPE}` + "=" + `{WHOSONFIRST_ID}`

Under the hood this is using the [whosonfirst/go-whosonfirst-spatial](https://github.com/whosonfirst/go-whosonfirst-spatial) package. That package is written in such a way as to be database-agnostic. It provides a default in-memory RTree-based spatial index but other (more performant) database implementations are defined in other packages.

That's the `-spatial-database-uri 'rtree:///?strict=false&index_alt_files=0'` part in the command above. That's also why it takes 27 seconds to index the [sfomuseum-data/sfomuseum-data-whosonfirst](https://github.com/sfomuseum-data/sfomuseum-data-whosonfirst) repository. There are package implementing the `go-whosonfirst-spatial` interfaces for the following databases:

* [whosonfirst/go-whosonfirst-spatial-sqlite](https://github.com/whosonfirst/go-whosonfirst-spatial-sqlite)
* [whosonfirst/go-whosonfirst-spatial-pmtiles](https://github.com/whosonfirst/go-whosonfirst-spatial-pmtiles)
* [whosonfirst/go-whosonfirst-spatial-duckdb](https://github.com/whosonfirst/go-whosonfirst-spatial-duckdb)

Support for these databases is _not_ bundled with this package. In order to use them you will need to clone the `cmd/marc-034d` tool and add the relevant. To that end the "guts" of that application have been moved in to an easy-to-use package (`app/server`) to save time-and-typing. For example here is how you would write a custom `marc-034d` tool to use a SQLite database (using the `go-whosonfirst-spatial-sqlite` package):

```
package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst-spatial-sqlite"
	"github.com/aaronland/go-marc/v3/app/convert"
)

func main() {

	ctx := context.Background()
	err := convert.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run tool, %v", err)
	}
}

```

Which would then be invoked like this:

```
$> ./bin/convert \
	-enable-intersects \
	-spatial-database-uri 'sqlite://mattn?dsn=/path/to/sqlite.db' \	
	fixtures/marc034-intersects.csv
```

_Note the absense of the `-spatial-database-source` flag because it is assumed the SQLite database has already been indexed (using the [whosonfirst/go-whosonfirst-database-sqlite](https://github.com/whosonfirst/go-whosonfirst-database-sqlite) package)._

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial