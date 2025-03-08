package server

import (
	"flag"
	"fmt"
	"os"

	"github.com/aaronland/go-http-maps/v2"
	"github.com/sfomuseum/go-flags/flagset"
	spatial_flags "github.com/whosonfirst/go-whosonfirst-spatial/flags"
)

var verbose bool
var server_uri string
var marc034_column string

var enable_intersects bool
var spatial_database_uri string
var spatial_database_sources spatial_flags.MultiCSVIteratorURIFlag

var map_provider string
var map_tile_uri string
var protomaps_theme string
var initial_view string
var leaflet_style string
var leaflet_point_style string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("marc-034")

	fs.StringVar(&map_provider, "map-provider", "leaflet", "Valid options are: leaflet, protomaps")
	fs.StringVar(&map_tile_uri, "map-tile-uri", maps.LEAFLET_OSM_TILE_URL, "A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs.")
	fs.StringVar(&protomaps_theme, "protomaps-theme", "white", "A valid Protomaps theme label.")
	fs.StringVar(&leaflet_style, "leaflet-style", "", "A custom Leaflet style definition for geometries. This may either be a JSON-encoded string or a path on disk.")
	fs.StringVar(&leaflet_point_style, "leaflet-point-style", "", "A custom Leaflet style definition for point geometries. This may either be a JSON-encoded string or a path on disk.")
	fs.StringVar(&initial_view, "initial-view", "", "A comma-separated string indicating the map's initial view. Valid options are: 'LON,LAT', 'LON,LAT,ZOOM' or 'MINX,MINY,MAXX,MAXY'.")

	fs.BoolVar(&enable_intersects, "enable-intersects", false, "...")
	fs.StringVar(&spatial_database_uri, "spatial-database-uri", "", "...")
	fs.Var(&spatial_database_sources, "spatial-database-source", "Zero or more...")

	fs.StringVar(&marc034_column, "marc034-column", "marc_034", "The name of the CSV column where MARC 034 data is stored.")
	fs.StringVar(&server_uri, "server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "marc-034d is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fs.PrintDefaults()
	}

	return fs
}
