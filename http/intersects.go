package http

import (
	"encoding/json"
	_ "log/slog"
	gohttp "net/http"

	"github.com/aaronland/go-http-sanitize"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spatial/query"
	"github.com/whosonfirst/go-whosonfirst-spr-geojson"
)

type IntersectsHandlerOptions struct {
	SpatialDatabase database.SpatialDatabase
	EnableGeoJSON   bool
}

const GEOJSON string = "application/geo+json"

// IntersectsHandler returns a `net/http.Handler` instance for convert MARC 034 fields in to GeoJSON Features.
func IntersectsHandler(opts *IntersectsHandlerOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		ctx := req.Context()

		if req.Method != "POST" {
			gohttp.Error(rsp, "Unsupported method", gohttp.StatusMethodNotAllowed)
			return
		}

		intersects_fn, err := query.NewSpatialFunction(ctx, "intersects://")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		var intersects_query *query.SpatialQuery

		dec := json.NewDecoder(req.Body)
		err = dec.Decode(&intersects_query)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		accept, err := sanitize.HeaderString(req, "Accept")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		if accept == GEOJSON && !opts.EnableGeoJSON {
			gohttp.Error(rsp, "GeoJSON output is not supported", gohttp.StatusBadRequest)
			return
		}

		intersects_rsp, err := query.ExecuteQuery(ctx, opts.SpatialDatabase, intersects_fn, intersects_query)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		if opts.EnableGeoJSON && accept == GEOJSON {

			opts := &geojson.AsFeatureCollectionOptions{
				Reader: opts.SpatialDatabase,
				Writer: rsp,
			}

			err := geojson.AsFeatureCollection(ctx, intersects_rsp, opts)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
				return
			}
		}

		enc := json.NewEncoder(rsp)
		err = enc.Encode(intersects_rsp)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		return
	}

	intersects_handler := gohttp.HandlerFunc(fn)
	return intersects_handler, nil
}
