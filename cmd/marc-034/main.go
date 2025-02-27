// marc-034 parses one or more MARC 034 strings and emit a (S, W, N, E) bounding box for each.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/aaronland/go-marc/fields"	
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse one or more MARC 034 strings and emit a (S, W, N, E) bounding box for each.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s MARC034(N) MARC034(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, raw := range flag.Args() {

		p, err := fields.Parse034(raw)

		if err != nil {
			log.Fatalf("Failed to parse '%s', %v", raw, err)
		}

		b, err := p.Bound()

		if err != nil {
			log.Fatalf("Failed to derive bounds for '%s', %v", raw, err)
		}

		bbox := []string{
			strconv.FormatFloat(b.Bottom(), 'f', -1, 64),
			strconv.FormatFloat(b.Left(), 'f', -1, 64),
			strconv.FormatFloat(b.Top(), 'f', -1, 64),
			strconv.FormatFloat(b.Right(), 'f', -1, 64),
		}

		str_bbox := strings.Join(bbox, ",")
		fmt.Println(str_bbox)
	}
}
