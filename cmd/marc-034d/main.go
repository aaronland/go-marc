// marc-034d is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).
package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	gohttp "net/http"
	"os"

	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-http-maps/v2"	
	"github.com/aaronland/go-marc/v2/http"
	"github.com/aaronland/go-marc/v2/static/www"
	"github.com/sfomuseum/go-flags/flagset"
)

const leaflet_osm_tile_url = "https://tile.openstreetmap.org/{z}/{x}/{y}.png"
const protomaps_api_tile_url string = "https://api.protomaps.com/tiles/v3/{z}/{x}/{y}.mvt?key={key}"

func main() {

	var verbose bool
	var server_uri string
	var marc034_column string
	
	var map_provider string
	var map_tile_uri string
	var protomaps_theme string
	var initial_view string
	var leaflet_style string
	var leaflet_point_style string	

	fs := flagset.NewFlagSet("marc-034")

	fs.StringVar(&map_provider, "map-provider", "leaflet", "Valid options are: leaflet, protomaps")
	fs.StringVar(&map_tile_uri, "map-tile-uri", leaflet_osm_tile_url, "A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs.")
	fs.StringVar(&protomaps_theme, "protomaps-theme", "white", "A valid Protomaps theme label.")
	fs.StringVar(&leaflet_style, "leaflet-style", "", "A custom Leaflet style definition for geometries. This may either be a JSON-encoded string or a path on disk.")
	fs.StringVar(&leaflet_point_style, "leaflet-point-style", "", "A custom Leaflet style definition for point geometries. This may either be a JSON-encoded string or a path on disk.")	
	fs.StringVar(&initial_view, "initial-view", "", "A comma-separated string indicating the map's initial view. Valid options are: 'LON,LAT', 'LON,LAT,ZOOM' or 'MINX,MINY,MAXX,MAXY'.")
	
	fs.StringVar(&marc034_column, "marc034-column", "marc_034", "The name of the CSV column where MARC 034 data is stored.")
	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

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

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()

	mux := gohttp.NewServeMux()

	opts := &maps.AssignMapConfigHandlerOptions{
		MapProvider:       map_provider,
		MapTileURI:        map_tile_uri,
		InitialView:       initial_view,
		LeafletStyle:      leaflet_style,
		LeafletPointStyle: leaflet_point_style,
		ProtomapsTheme:    protomaps_theme,
	}

	err = maps.AssignMapConfigHandler(opts, mux, "/map.json")

	if err != nil {
		log.Fatalf("Failed to assign map config handler, %v", err)
	}
	
	bbox_handler, err := http.BboxHandler()

	if err != nil {
		log.Fatalf("Failed to create bbox handler, %v", err)
	}

	mux.Handle("/bbox", bbox_handler)

	convert_opts := &http.ConvertHandlerOptions{
		Marc034Column: marc034_column,
	}

	convert_handler, err := http.ConvertHandler(convert_opts)

	if err != nil {
		log.Fatalf("Failed to create convert handler, %v", err)
	}

	mux.Handle("/convert", convert_handler)

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
