// marc-034d is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).
package main

import (
	"context"
	"fmt"
	"log"
	gohttp "net/http"
	"os"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-marc/http"
	"github.com/aaronland/go-marc/static/www"
	"github.com/sfomuseum/go-flags/flagset"
)

func main() {

	fs := flagset.NewFlagSet("marc-034")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "marc-034d is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "MARC")

	if err != nil {
		log.Fatalf("Failed to assign flags from environment variables, %v", err)
	}

	ctx := context.Background()

	mux := gohttp.NewServeMux()

	www_fs := gohttp.FS(www.FS)
	www_handler := gohttp.FileServer(www_fs)

	bbox_handler, err := http.BboxHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", www_handler)
	mux.Handle("/bbox", bbox_handler)

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	log.Printf("listening on %s\n", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}

	os.Exit(0)
}
