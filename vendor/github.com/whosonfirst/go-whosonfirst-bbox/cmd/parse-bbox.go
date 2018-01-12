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
