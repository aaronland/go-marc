// server is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).
package server

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	gohttp "net/http"

	"github.com/aaronland/go-http-maps/v2"
	"github.com/aaronland/go-http-server"
	"github.com/aaronland/go-marc/v3/http"
	"github.com/aaronland/go-marc/v3/static/www"
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

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	mux := gohttp.NewServeMux()

	maps_opts := &maps.AssignMapConfigHandlerOptions{
		MapProvider:       opts.MapProvider,
		MapTileURI:        opts.MapTileURI,
		InitialView:       opts.InitialView,
		LeafletStyle:      opts.LeafletStyle,
		LeafletPointStyle: opts.LeafletPointStyle,
		ProtomapsTheme:    opts.ProtomapsTheme,
	}

	err := maps.AssignMapConfigHandler(maps_opts, mux, "/map.json")

	if err != nil {
		return fmt.Errorf("Failed to assign map config handler, %w", err)
	}

	bbox_handler, err := http.BboxHandler()

	if err != nil {
		return fmt.Errorf("Failed to create bbox handler, %w", err)
	}

	mux.Handle("/bbox", bbox_handler)

	convert_opts := &http.ConvertHandlerOptions{
		MARC034Column: opts.MARC034Column,
	}

	// Note: We defer creating or registering the convert handler
	// until after we have (or have not) dealt with spatial intersects
	// stuff below.

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

		// See notes above
		convert_opts.EnableIntersects = true
		convert_opts.SpatialDatabase = db

		intersects_opts := &http.IntersectsHandlerOptions{
			SpatialDatabase: db,
		}

		intersects_handler, err := http.IntersectsHandler(intersects_opts)

		if err != nil {
			return fmt.Errorf("Failed to create intersects handler, %w", err)
		}

		mux.Handle("/intersects", intersects_handler)
	}

	convert_handler, err := http.ConvertHandler(convert_opts)

	if err != nil {
		return fmt.Errorf("Failed to create convert handler, %w", err)
	}

	mux.Handle("/convert", convert_handler)

	www_fs := gohttp.FS(www.FS)
	www_handler := gohttp.FileServer(www_fs)

	mux.Handle("/", www_handler)

	s, err := server.NewServer(ctx, opts.ServerURI)

	if err != nil {
		return fmt.Errorf("Failed to create new server, %w", err)
	}

	slog.Info("Listening for requests", "address", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to serve requests, %w", err)
	}

	return nil
}
