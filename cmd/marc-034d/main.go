// marc-034d is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).
package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	gohttp "net/http"
	"os"
	
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-ping/v2"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-marc/http"
	"github.com/aaronland/go-marc/templates/html"
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

	t := template.New("marc")

	t, err = t.ParseFS(html.FS, "*.html")

	if err != nil {
		log.Fatalf("Failed to parse templates, %v", err)
	}

	mux := gohttp.NewServeMux()

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()	
	err = bootstrap.AppendAssetHandlers(mux, bootstrap_opts)

	if err != nil {
		log.Fatalf("Failed to append Bootstrap asset handlers, %v", err)
	}

	err = http.AppendStaticAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append static asset handlers, %v", err)
	}

	www_handler, err := http.MARC034Handler(t)

	if err != nil {
		log.Fatalf("Failed to create MARC034 handler, %v", err)
	}

	www_handler = bootstrap.AppendResourcesHandler(www_handler, bootstrap_opts)

	bbox_handler, err := http.BboxHandler()

	if err != nil {
		log.Fatal(err)
	}

	ping_handler, err := ping.PingPongHandler()

	if err != nil {
		log.Fatalf("Failed to create ping handler, %v", err)
	}

	mux.Handle("/", www_handler)
	mux.Handle("/bbox", bbox_handler)
	mux.Handle("/ping", ping_handler)

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
