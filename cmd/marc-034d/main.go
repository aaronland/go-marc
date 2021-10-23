package main

import (
	"context"
	"fmt"
	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-ping/v2"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-tangramjs"
	"github.com/aaronland/go-marc/http"
	"github.com/aaronland/go-marc/templates/html"
	"github.com/sfomuseum/go-flags/flagset"
	tiles_http "github.com/tilezen/go-tilepacks/http"
	"github.com/tilezen/go-tilepacks/tilepack"
	"html/template"
	"log"
	gohttp "net/http"
	"os"
	"strings"
)

func main() {

	fs := flagset.NewFlagSet("marc-034")

	server_uri := fs.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")

	nextzen_api_key := fs.String("nextzen-api-key", "nextzen-xxxxxx", "A valid Nextzen API key")
	nextzen_style_url := fs.String("nextzen-style-url", "/tangram/refill-style.zip", "A valid Nextzen style URL")

	tilepack_db := fs.String("nextzen-tilepack-database", "", "The path to a valid MBTiles database (tilepack) containing Nextzen MVT tiles.")

	tilepack_uri := fs.String("nextzen-tilepack-uri", "/tilezen/vector/v1/512/all/{z}/{x}/{y}.mvt", "The relative URI to serve Nextzen MVT tiles from a MBTiles database (tilepack).")

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

	if *tilepack_db != "" {
		tangramjs_opts.NextzenOptions.TileURL = *tilepack_uri
	}

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

	if *tilepack_db != "" {

		tiles_reader, err := tilepack.NewMbtilesReader(*tilepack_db)

		if err != nil {
			log.Fatalf("Failed to load tilepack, %v", err)
		}

		u := strings.TrimLeft(*tilepack_uri, "/")
		p := strings.Split(u, "/")
		path_tiles := fmt.Sprintf("/%s/", p[0])

		tiles_handler := tiles_http.MbtilesHandler(tiles_reader)
		mux.Handle(path_tiles, tiles_handler)
	}

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
