package http

import (
	gohttp "net/http"
	"strconv"

	"github.com/aaronland/go-marc/v2/fields"
	"github.com/sfomuseum/go-csvdict/v2"
)

type ConvertHandlerOptions struct {
	Marc034Column string
	MinXColumn    string
	MinYColumn    string
	MaxXColumn    string
	MaxYColumn    string
}

// ConvertHandler returns a `net/http.Handler` instance for convert MARC 034 fields in a CSV file.
func ConvertHandler(opts *ConvertHandlerOptions) (gohttp.Handler, error) {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		if req.Method != gohttp.MethodPost {
			gohttp.Error(rsp, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}

		csv_r, err := csvdict.NewReader(req.Body)

		if err != nil {
			gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
			return
		}

		var csv_wr *csvdict.Writer

		rsp.Header().Set("Content-Type", "text/csv")
		rsp.Header().Set("Access-Control-Allow-Origin", "*")

		for row, err := range csv_r.Iterate() {

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			value, ok := row[opts.Marc034Column]

			if !ok {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			p, err := fields.Parse034(value)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			b, err := p.Bound()

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}

			row[opts.MinXColumn] = strconv.FormatFloat(b.Left(), 'f', -1, 64)
			row[opts.MinYColumn] = strconv.FormatFloat(b.Bottom(), 'f', -1, 64)
			row[opts.MaxXColumn] = strconv.FormatFloat(b.Right(), 'f', -1, 64)
			row[opts.MaxYColumn] = strconv.FormatFloat(b.Top(), 'f', -1, 64)

			if csv_wr == nil {

				wr, err := csvdict.NewWriter(rsp)

				if err != nil {
					gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
					return
				}

				csv_wr = wr
			}

			err = csv_wr.WriteRow(row)

			if err != nil {
				gohttp.Error(rsp, err.Error(), gohttp.StatusBadRequest)
				return
			}
		}

		csv_wr.Flush()
		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
