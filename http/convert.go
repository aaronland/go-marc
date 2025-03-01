package http

import (
	gohttp "net/http"

	"github.com/aaronland/go-marc/v2/csv"
)

type ConvertHandlerOptions struct {
	Marc034Column string
}

// ConvertHandler returns a `net/http.Handler` instance for converting one or more MARC 034 fields in a CSV file
// and returning a new CSV file containing both the original data (MARC 034 and other) as well as the min_x, min_y,
// max_x and max_y values for each MARC 034 value.
func ConvertHandler(opts *ConvertHandlerOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		if req.Method != gohttp.MethodPost {
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}

		defer req.Body.Close()

		rsp.Header().Set("Content-Type", "text/csv")

		err := csv.Convert034(req.Body, rsp, opts.Marc034Column)

		if err != nil {
			gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
			return
		}
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
