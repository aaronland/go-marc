package main

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-marc/fields"
	"github.com/sfomuseum/go-csvdict"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	marc034_column := flag.String("marc-034-column", "marc_034", "...")
	minx_column := flag.String("min-x-column", "min_x", "...")
	miny_column := flag.String("min-y-column", "min_y", "...")
	maxx_column := flag.String("max-x-column", "max_x", "...")
	maxy_column := flag.String("max-y-column", "max_y", "...")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "...\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s csv-file(N) csv-file(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	writers := []io.Writer{
		os.Stdout,
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

		for {
			row, err := csv_r.Read()

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

				fieldnames := make([]string, 0)

				for k, _ := range row {
					fieldnames = append(fieldnames, k)
				}

				sort.Strings(fieldnames)

				wr, err := csvdict.NewWriter(mw, fieldnames)

				if err != nil {
					log.Fatalf("Failed to create new CSV writer, %v", err)
				}

				err = wr.WriteHeader()

				if err != nil {
					log.Fatalf("Failed to write CSV header, %v", err)
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
