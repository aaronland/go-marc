// marc-034-convert is a command line tool to process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/aaronland/go-marc/v2/fields"
	"github.com/sfomuseum/go-csvdict/v2"
)

func main() {

	marc034_column := flag.String("marc-034-column", "marc_034", "The name of the CSV column where MARC 034 data is stored.")
	minx_column := flag.String("min-x-column", "min_x", "The name of the CSV column where the left-side coordinate (min x) of the bounding box should be stored.")
	miny_column := flag.String("min-y-column", "min_y", "The name of the CSV column where the bottom-side coordinate (min y) of the bounding box should be stored.")
	maxx_column := flag.String("max-x-column", "max_x", "The name of the CSV column where the right-side coordinate (max x) of the bounding box should be stored.")
	maxy_column := flag.String("max-y-column", "max_y", "The name of the CSV column where the top-side coordinate (max y) of the bounding box should be stored.")

	to_file := flag.String("to-file", "", "The path where your new CSV file should be created.")
	to_stdout := flag.Bool("to-stdout", false, "Output CSV data to STDOUT.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Process one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s csv-file(N) csv-file(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

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

	var csv_wr *csvdict.Writer

	for _, path := range flag.Args() {

		r, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open %s, %v", path, err)
		}

		defer r.Close()

		csv_r, err := csvdict.NewReader(r)

		if err != nil {
			log.Fatalf("Failed to create CSV reader for %s, %v", path, err)
		}

		for row, err := range csv_r.Iterate() {

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Failed to read row for %s, %v", path, err)
			}

			value, ok := row[*marc034_column]

			if !ok {
				log.Fatalf("Row missing '%s' column in %s", *marc034_column, path)
			}

			p, err := fields.Parse034(value)

			if err != nil {
				log.Fatalf("Failed to parse '%s' in %s, %v", value, path, err)
			}

			b, err := p.Bound()

			if err != nil {
				log.Fatalf("Failed to derive bounds for '%s' (parsed) in %s, %v", value, path, err)
			}

			row[*minx_column] = strconv.FormatFloat(b.Left(), 'f', -1, 64)
			row[*miny_column] = strconv.FormatFloat(b.Bottom(), 'f', -1, 64)
			row[*maxx_column] = strconv.FormatFloat(b.Right(), 'f', -1, 64)
			row[*maxy_column] = strconv.FormatFloat(b.Top(), 'f', -1, 64)

			if csv_wr == nil {

				wr, err := csvdict.NewWriter(mw)

				if err != nil {
					log.Fatalf("Failed to create new CSV writer, %v", err)
				}

				csv_wr = wr
			}

			err = csv_wr.WriteRow(row)

			if err != nil {
				log.Fatalf("Failed to write row for %s, %v", path, err)
			}

			csv_wr.Flush()
		}
	}
}
