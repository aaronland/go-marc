package main

import (
	"flag"
	"github.com/thisisaaronland/go-marc/fields"
	"log"
)

func main() {

	var f = flag.String("f", "1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000", "A valid MARC 034 string")

	flag.Parse()

	log.Println(*f)

	p, err := fields.Parse034(*f)

	if err != nil {
		log.Fatal(err)
	}

	// log.Println(p)

	/*
		c, err := fields.Parse034Coordinate(p.Subfields["$d"].Value)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(c)
	*/

	bbox, err := p.BoundingBox()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(bbox)

}
