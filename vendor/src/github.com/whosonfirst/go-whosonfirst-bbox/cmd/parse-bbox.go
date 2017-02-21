package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-bbox/parser"
	"log"
)

func main() {

	var bbox = flag.String("bbox", "", "...")
	var scheme = flag.String("scheme", "cardinal", "...")
	var order = flag.String("order", "swne", "...")
	var separator = flag.String("separator", ",", "...")

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
