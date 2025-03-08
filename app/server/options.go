package server

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	ServerURI              string
	MARC034Column          string
	MapProvider            string
	MapTileURI             string
	InitialView            string
	LeafletStyle           string
	LeafletPointStyle      string
	ProtomapsTheme         string
	EnableIntersects       bool
	SpatialDatabaseURI     string
	SpatialDatabaseSources map[string][]string
	Verbose                bool
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "MARC")

	if err != nil {
		return nil, fmt.Errorf("Failed to assign flags from environment variables, %w", err)
	}

	opts := &RunOptions{
		ServerURI:              server_uri,
		MARC034Column:          marc034_column,
		MapProvider:            map_provider,
		MapTileURI:             map_tile_uri,
		LeafletStyle:           leaflet_style,
		LeafletPointStyle:      leaflet_point_style,
		ProtomapsTheme:         protomaps_theme,
		EnableIntersects:       enable_intersects,
		SpatialDatabaseURI:     spatial_database_uri,
		SpatialDatabaseSources: spatial_database_sources.AsMap(),
		Verbose:                verbose,
	}

	return opts, nil
}
