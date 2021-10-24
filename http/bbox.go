package http

import (
	"encoding/json"
	"github.com/aaronland/go-marc/fields"
	"github.com/whosonfirst/go-sanitize"
	_ "log"
	gohttp "net/http"
	"strings"
)

type GeoJSONCoordinate []float64

type GeoJSONRing []GeoJSONCoordinate

type GeoJSONPolygon []GeoJSONRing

type GeoJSONGeometry struct {
	Type        string         `json:"type"`
	Coordinates GeoJSONPolygon `json:"coordinates"`
}

type GeoJSONProperties map[string]string

type GeoJSONBoundingBox []float64

type GeoJSONFeature struct {
	Type        string             `json:"type"`
	Geometry    GeoJSONGeometry    `json:"geometry"`
	Properties  GeoJSONProperties  `json:"properties"`
	BoundingBox GeoJSONBoundingBox `json:"bbox"`
}

type BboxResponse struct {
	MinX float64 `json:"min_x"`
	MinY float64 `json:"min_y"`
	MaxX float64 `json:"max_x"`
	MaxY float64 `json:"max_y"`
}

// BboxHandler returns a `net/http.Handler` instance for convert MARC 034 fields in to GeoJSON Features.
func BboxHandler() (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		query := req.URL.Query()

		marc_raw := query.Get("034")
		marc_raw = strings.Trim(marc_raw, " ")

		if marc_raw == "" {
			gohttp.Error(rsp, "Missing or invalid bounding box", gohttp.StatusBadRequest)
			return
		}

		opts := sanitize.DefaultOptions()
		marc_clean, err := sanitize.SanitizeString(marc_raw, opts)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		// log.Println("RAW", marc_raw)
		// log.Println("CLEAN", marc_clean)

		parsed, err := fields.Parse034(marc_clean)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		bbox, err := parsed.Bound()

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		min_x := bbox.Left()
		min_y := bbox.Bottom()
		max_x := bbox.Right()
		max_y := bbox.Top()

		sw := GeoJSONCoordinate{min_x, min_y}
		nw := GeoJSONCoordinate{min_x, max_y}
		ne := GeoJSONCoordinate{max_x, max_y}
		se := GeoJSONCoordinate{max_x, min_y}

		geojson_ring := GeoJSONRing{sw, nw, ne, se, sw}
		geojson_polygon := GeoJSONPolygon{geojson_ring}

		geojson_geometry := GeoJSONGeometry{
			Type:        "Polygon",
			Coordinates: geojson_polygon,
		}

		geojson_bbox := GeoJSONBoundingBox{min_x, min_y, max_x, max_y}

		geojson_properties := GeoJSONProperties{
			"marc:034": marc_clean,
		}

		geojson_feature := GeoJSONFeature{
			Type:        "Feature",
			Geometry:    geojson_geometry,
			Properties:  geojson_properties,
			BoundingBox: geojson_bbox,
		}

		enc, err := json.Marshal(geojson_feature)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "application/json")
		rsp.Header().Set("Access-Control-Allow-Origin", "*")

		rsp.Write(enc)

	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
