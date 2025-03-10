// convert processes one or more CSV files containing MARC 034 data and append bounding box information to a new CSV document.
package convert

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/aaronland/go-marc/v2/csv"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	if opts.Verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	writers := make([]io.Writer, 0)

	if opts.ToFile != "" {

		fh, err := os.OpenFile(opts.ToFile, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			return fmt.Errorf("Failed to open %s for writing, %v", opts.ToFile, err)
		}

		defer fh.Close()

		writers = append(writers, fh)
	}

	if opts.ToStdout {
		writers = append(writers, os.Stdout)
	}

	if len(writers) == 0 {
		return fmt.Errorf("You must configure at least one output target")
	}

	mw := io.MultiWriter(writers...)

	convert_opts := &csv.Convert034Options{
		MARC034Column: opts.MARC034Column,
	}

	if opts.EnableIntersects {

		db, err := database.NewSpatialDatabase(ctx, opts.SpatialDatabaseURI)

		if err != nil {
			return fmt.Errorf("Failed to create new spatial database, %w", err)
		}

		if len(opts.SpatialDatabaseSources) > 0 {

			slog.Info("Indexing spatial database.")

			err = database.IndexDatabaseWithIterators(ctx, db, opts.SpatialDatabaseSources)

			if err != nil {
				return fmt.Errorf("Failed to index database, %w", err)
			}
		}

		convert_opts.EnableIntersects = true
		convert_opts.SpatialDatabase = db
	}

	for _, path := range opts.Files {

		slog.Info("Parse file", "path", path)
		var to_read io.ReadCloser

		switch path {
		case "-":
			to_read = os.Stdin
		default:

			r, err := os.Open(path)

			if err != nil {
				return fmt.Errorf("Failed to open %s, %v", path, err)
			}

			to_read = r
			defer to_read.Close()
		}

		err := csv.Convert034(ctx, to_read, mw, convert_opts)

		if err != nil {
			return fmt.Errorf("Failed to convert %s, %v", path, err)
		}
	}

	return nil
}
