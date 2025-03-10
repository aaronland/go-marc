// marc-034-convert is a command line tool to process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aaronland/go-marc/v2/csv"
)

func main() {

	marc034_column := flag.String("marc-034-column", "marc_034", "The name of the CSV column where MARC 034 data is stored.")

	to_file := flag.String("to-file", "", "The path where your new CSV file should be created.")
	to_stdout := flag.Bool("to-stdout", false, "Output CSV data to STDOUT.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s csv-file(N) csv-file(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	writers := make([]io.Writer, 0)

	if *to_file != "" {

		fh, err := os.OpenFile(*to_file, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			log.Fatalf("Failed to open %s for writing, %v", *to_file, err)
		}

		defer fh.Close()

		writers = append(writers, fh)
	}

	if *to_stdout {
		writers = append(writers, os.Stdout)
	}

	if len(writers) == 0 {
		log.Fatalf("You must configure at least one output target")
	}

	mw := io.MultiWriter(writers...)

	opts := &csv.Convert034Options{
		MARC034Column: *marc034_column,
	}

	for _, path := range flag.Args() {

		r, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open %s, %v", path, err)
		}

		defer r.Close()

		err = csv.Convert034(ctx, r, mw, opts)

		if err != nil {
			log.Fatalf("Failed to convert %s, %v", path, err)
		}
	}
}
