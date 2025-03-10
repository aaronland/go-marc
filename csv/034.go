package csv

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"

	"github.com/aaronland/go-marc/v3/fields"
	"github.com/paulmach/orb/geojson"
	"github.com/sfomuseum/go-csvdict/v2"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spatial/query"
)

type Convert034Options struct {
	MARC034Column    string
	EnableIntersects bool
	SpatialDatabase  database.SpatialDatabase
}

// Convert034 will process 'r' as though it were a CSV data containing a MARC 034 data appending `min_x`, `min_y`
// `max_x` and `max_y` columns (derived from the MARC 034 data. It was also append a `valid` column indicating
// whether the MARC 034 data could be parsed as well as an `error` column which will be populated in the event
// that MARC 034 data could not be parsed.
func Convert034(ctx context.Context, r io.Reader, wr io.Writer, opts *Convert034Options) error {

	defaults := map[string]string{
		"min_x":      "",
		"min_y":      "",
		"max_x":      "",
		"max_y":      "",
		"valid":      "0",
		"error":      "",
		"intersects": "",
	}

	csv_r, err := csvdict.NewReader(r)

	if err != nil {
		slog.Error("Failed to create new CSV reader", "error", err)
		return err
	}

	var csv_wr *csvdict.Writer
	var intersects_fn query.SpatialFunction

	if opts.EnableIntersects {

		fn, err := query.NewSpatialFunction(ctx, "intersects://")

		if err != nil {
			return err
		}

		intersects_fn = fn
	}

	for row, err := range csv_r.Iterate() {

		if err != nil {
			slog.Error("Failed to iterate row", "error", err)
			return err
		}

		slog.Debug("Process", "row", row)

		for k, v := range defaults {
			row[k] = v
		}

		value, ok := row[opts.MARC034Column]

		if !ok {
			slog.Error("Row is missing MARC 034 column", "column", opts.MARC034Column)
			row["error"] = "Missing MARC 034"
			continue
		}

		p, err := fields.Parse034(value)

		if err != nil {
			slog.Error("Failed to parse MARC 034 value", "value", value, "error", err)
			row["error"] = err.Error()
			continue
		}

		b, err := p.Bound()

		if err != nil {
			slog.Error("Failed to derive bounds for MARC 034 value", "value", value, "error", err)
			row["error"] = err.Error()
			continue
		}

		row["min_x"] = strconv.FormatFloat(b.Left(), 'f', -1, 64)
		row["min_y"] = strconv.FormatFloat(b.Bottom(), 'f', -1, 64)
		row["max_x"] = strconv.FormatFloat(b.Right(), 'f', -1, 64)
		row["max_y"] = strconv.FormatFloat(b.Top(), 'f', -1, 64)
		row["valid"] = "1"

		if opts.EnableIntersects {

			orb_geom := b.ToPolygon()
			geojson_geom := geojson.NewGeometry(orb_geom)

			intersects_query := &query.SpatialQuery{
				Geometry: geojson_geom,
				Sort: []string{
					"placetype://",
					"name://",
					"inception://",
				},
			}

			intersects_rsp, err := query.ExecuteQuery(ctx, opts.SpatialDatabase, intersects_fn, intersects_query)

			if err != nil {
				slog.Error("Failed to execute intersects query", "error", err)
				row["error"] = err.Error()
			} else {

				results := intersects_rsp.Results()
				ids := make([]string, len(results))

				for idx, pl := range results {
					ids[idx] = fmt.Sprintf("wof:%s=%s", pl.Placetype(), pl.Id())
				}

				row["intersects"] = strings.Join(ids, ",")
			}
		}

		if csv_wr == nil {

			new_wr, err := csvdict.NewWriter(wr)

			if err != nil {
				slog.Error("Failed to create CSV writer", "error", err)
				return err
			}

			csv_wr = new_wr
		}

		err = csv_wr.WriteRow(row)

		if err != nil {
			slog.Error("Failed to write CSV row", "error", err)
			return err
		}
	}

	if csv_wr != nil {
		csv_wr.Flush()
	}

	return nil
}
