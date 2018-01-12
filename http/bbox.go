package http

import (
	"encoding/json"
	"github.com/thisisaaronland/go-marc/fields"
	"github.com/whosonfirst/go-sanitize"
	_ "log"
	gohttp "net/http"
	"strings"
)

type BboxResponse struct {
     MinX	  float64 `json:"min_x"`
     MinY	  float64 `json:"min_y"`
     MaxX	  float64 `json:"max_x"`
     MaxY	  float64 `json:"max_y"`    
}


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
		
		bounds, err := fields.Parse034(marc_clean)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		bbox, err := bounds.BoundingBox()

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusInternalServerError)
			return
		}

		r := BboxResponse{
			MinX: bbox.MinX(),
			MinY: bbox.MinY(),
			MaxX: bbox.MaxX(),
			MaxY: bbox.MaxY(),
		}

		enc, err := json.Marshal(r)

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
