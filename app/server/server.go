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
	"github.com/aaronland/go-marc/v2/http"
	"github.com/aaronland/go-marc/v2/static/www"
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
		return fmt.Errorf("Failed to assign map config handler, %v", err)
	}

	bbox_handler, err := http.BboxHandler()

	if err != nil {
		return fmt.Errorf("Failed to create bbox handler, %v", err)
	}

	mux.Handle("/bbox", bbox_handler)

	convert_opts := &http.ConvertHandlerOptions{
		Marc034Column: opts.MARC034Column,
	}

	convert_handler, err := http.ConvertHandler(convert_opts)

	if err != nil {
		return fmt.Errorf("Failed to create convert handler, %v", err)
	}

	mux.Handle("/convert", convert_handler)

	www_fs := gohttp.FS(www.FS)
	www_handler := gohttp.FileServer(www_fs)

	mux.Handle("/", www_handler)

	s, err := server.NewServer(ctx, opts.ServerURI)

	if err != nil {
		return fmt.Errorf("Failed to create new server, %v", err)
	}

	slog.Info("Listening for requests", "address", s.Address())

	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		return fmt.Errorf("Failed to serve requests, %v", err)
	}

	return nil
}
