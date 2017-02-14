package main

import (
	"github.com/thisisaaronland/go-marc/fields"
	"log"
)

func main() {

	raw := "1#$aa$b22000000$dW1800000$eE1800000$fN0840000$gS0700000"
	log.Println(raw)
	
	p, err := fields.Parse034(raw)

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
