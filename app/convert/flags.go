package convert

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
)

var marc034_column string
var to_file string
var to_stdout bool

var enable_intersects bool
var spatial_database_uri string
var spatial_database_sources spatial_flags.MultiCSVIteratorURIFlag

var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("convert")

	fs.StringVar(&marc034_column, "marc-034-column", "marc_034", "The name of the CSV column where MARC 034 data is stored.")
	fs.StringVar(&to_file, "to-file", "", "The path where your new CSV file should be created.")
	fs.BoolVar(&to_stdout, "to-stdout", true, "Output CSV data to STDOUT.")

	fs.BoolVar(&enable_intersects, "enable-intersects", false, "Enable intersecting geometry lookups for MARC034-derived bounding boxes.")
	fs.StringVar(&spatial_database_uri, "spatial-database-uri", "", "A registered whosonfirst/go-whosonfirst-spatial/database.SpatialDatabase URI.")
	fs.Var(&spatial_database_sources, "spatial-database-source", "Zero or more '{ITERATOR_URI}#{ITERATOR_SOURCE}' strings following the whosonfirst/go-whosonfirst-iterate/v2 URI syntax.")

	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s csv-file(N) csv-file(N)\n", os.Args[0])
		fs.PrintDefaults()
	}

	return fs
}
