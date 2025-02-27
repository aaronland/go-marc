package http

import (
	gohttp "net/http"

	"github.com/aaronland/go-http-sanitize"
	"github.com/aaronland/go-marc/fields"
)

// BboxHandler returns a `net/http.Handler` instance for convert MARC 034 fields in to GeoJSON Features.
func BboxHandler() (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		marc_clean, err := sanitize.GetString(req, "034")

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		parsed, err := fields.Parse034(marc_clean)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		f, err := parsed.AsGeoJSON()

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		enc, err := f.MarshalJSON()

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
