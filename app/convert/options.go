package convert

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	MARC034Column          string
	ToFile                 string
	ToStdout               bool
	Files                  []string
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
		MARC034Column:          marc034_column,
		ToFile:                 to_file,
		ToStdout:               to_stdout,
		Files:                  fs.Args(),
		EnableIntersects:       enable_intersects,
		SpatialDatabaseURI:     spatial_database_uri,
		SpatialDatabaseSources: spatial_database_sources.AsMap(),
		Verbose:                verbose,
	}

	return opts, nil
}
