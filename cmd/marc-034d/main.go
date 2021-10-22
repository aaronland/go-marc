package main

import (
	"context"
	"flag"
	_ "fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-ping/v2"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/aaronland/go-marc/http"
	"github.com/aaronland/go-marc/templates/html"
	"html/template"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")

	nextzen_api_key := flag.String("nextzen-api-key", "mapzen-xxxxxx", "A valid Nextzen API key")
	nextzen_style_url := flag.String("nextzen-style-url", "/tangram/refill-style.zip", "A valid Nextzen style URL")

	flag.Parse()

	ctx := context.Background()

	t := template.New("marc")

	t, err := t.ParseFS(html.FS, "*.html")

	if err != nil {
		log.Fatalf("Failed to parse templates, %v", err)
	}

	mux := gohttp.NewServeMux()

	err = bootstrap.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append Bootstrap asset handlers, %v", err)
	}

	err = tangramjs.AppendAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append Tangram asset handlers, %v", err)
	}

	err = http.AppendStaticAssetHandlers(mux)

	if err != nil {
		log.Fatalf("Failed to append static asset handlers, %v", err)
	}

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	tangramjs_opts := tangramjs.DefaultTangramJSOptions()
	tangramjs_opts.NextzenOptions.APIKey = *nextzen_api_key
	tangramjs_opts.NextzenOptions.StyleURL = *nextzen_style_url

	www_handler, err := http.WWWHandler(t)

	if err != nil {
		log.Fatal(err)
	}

	www_handler = tangramjs.AppendResourcesHandler(www_handler, tangramjs_opts)
	www_handler = bootstrap.AppendResourcesHandler(www_handler, bootstrap_opts)

	bbox_handler, err := http.BboxHandler()

	if err != nil {
		log.Fatal(err)
	}

	ping_handler, err := ping.PingPongHandler()

	if err != nil {
		log.Fatal(err)
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
