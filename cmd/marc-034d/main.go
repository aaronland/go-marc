// marc-034d is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).
package main

import (
	"context"
	"fmt"
	"log"
	gohttp "net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-marc/http"
	"github.com/aaronland/go-marc/static/www"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-http-protomaps"
)

const leaflet_osm_tile_url = "https://tile.openstreetmap.org/{z}/{x}/{y}.png"
const protomaps_api_tile_url string = "https://api.protomaps.com/tiles/v3/{z}/{x}/{y}.mvt?key={key}"

func main() {

	var server_uri string
	var map_provider string
	var map_tile_uri string
	var protomaps_theme string

	var style string

	fs := flagset.NewFlagSet("marc-034")

	fs.StringVar(&map_provider, "map-provider", "leaflet", "Valid options are: leaflet, protomaps")
	fs.StringVar(&map_tile_uri, "map-tile-uri", leaflet_osm_tile_url, "A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs.")
	fs.StringVar(&protomaps_theme, "protomaps-theme", "white", "A valid Protomaps theme label.")

	fs.StringVar(&style, "style", "", "A custom Leaflet style definition for geometries. This may either be a JSON-encoded string or a path on disk.")

	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI")

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

	bbox_handler, err := http.BboxHandler()

	if err != nil {
		log.Fatalf("Failed to create bbox handler, %v", err)
	}

	mux.Handle("/bbox", bbox_handler)

	// START OF put me in a function or something...

	map_cfg := &http.MapConfig{
		Provider: map_provider,
		TileURL:  map_tile_uri,
		// Style:           style,
	}

	if map_provider == "protomaps" {

		u, err := url.Parse(map_tile_uri)

		if err != nil {
			log.Fatalf("Failed to parse Protomaps tile URL, %v", err)
		}

		switch u.Scheme {
		case "file":

			mux_url, mux_handler, err := protomaps.FileHandlerFromPath(u.Path, "")

			if err != nil {
				log.Fatalf("Failed to determine absolute path for '%s', %v", map_tile_uri, err)
			}

			mux.Handle(mux_url, mux_handler)
			map_cfg.TileURL = mux_url

		case "api":
			key := u.Host
			map_cfg.TileURL = strings.Replace(protomaps_api_tile_url, "{key}", key, 1)
		}

		map_cfg.Protomaps = &http.ProtomapsConfig{
			Theme: protomaps_theme,
		}
	}

	// END OF put me in a function or something...

	map_cfg_handler := http.MapConfigHandler(map_cfg)

	mux.Handle("/map.json", map_cfg_handler)

	//

	www_fs := gohttp.FS(www.FS)
	www_handler := gohttp.FileServer(www_fs)

	mux.Handle("/", www_handler)

	s, err := server.NewServer(ctx, server_uri)

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
