package convert

import (
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	MARC034Column string
	ToFile        string
	ToStdout      bool
	Files         []string
	Verbose       bool
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "MARC")

	if err != nil {
		return nil, fmt.Errorf("Failed to assign flags from environment variables, %w", err)
	}

	opts := &RunOptions{
		MARC034Column: marc034_column,
		ToFile:        to_file,
		ToStdout:      to_stdout,
		Files:         fs.Args(),
		Verbose:       verbose,
	}

	return opts, nil
}
