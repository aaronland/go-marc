package convert

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
)

var marc034_column string
var to_file string
var to_stdout bool
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("convert")

	fs.StringVar(&marc034_column, "marc-034-column", "marc_034", "The name of the CSV column where MARC 034 data is stored.")
	fs.StringVar(&to_file, "to-file", "", "The path where your new CSV file should be created.")
	fs.BoolVar(&to_stdout, "to-stdout", true, "Output CSV data to STDOUT.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s csv-file(N) csv-file(N)\n", os.Args[0])
		fs.PrintDefaults()
	}

	return fs
}
