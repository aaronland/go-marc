package main

import (
	"flag"
	"fmt"
	"github.com/thisisaaronland/go-marc/http"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")

	flag.Parse()

	www_handler, err := http.WWWHandler()

	if err != nil {
		log.Fatal(err)
	}

	/*
	static_handler, err := http.StaticHandler()

	if err != nil {
		log.Fatal(err)
	}
	*/
	
	bbox_handler, err := http.BboxHandler()

	if err != nil {
		log.Fatal(err)
	}

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()

	mux.Handle("/", www_handler)
	// mux.Handle("/javascript/", static_handler)
	// mux.Handle("/css/", static_handler)

	mux.Handle("/bbox", bbox_handler)
	mux.Handle("/ping", ping_handler)

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("listening on %s\n", address)

	err = gohttp.ListenAndServe(address, mux)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
