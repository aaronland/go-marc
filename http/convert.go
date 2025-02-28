package http

import (
	"fmt"
	"log/slog"
	gohttp "net/http"
	"strconv"
	"time"

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

		csv_r, err := csvdict.NewReader(req.Body)

		if err != nil {
			slog.Error("Failed to create new CSV reader", "error", err)
			gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
			return
		}

		rsp.Header().Set("Content-Type", "text/csv")		

		now := time.Now()
		ts := now.Unix()

		disposition := fmt.Sprintf("attachment; filename='marc034-bbox-%d.csv'", ts)
		rsp.Header().Set("Content-Disposition", disposition)

		slog.Debug("Process", "disposition", disposition)

		var csv_wr *csvdict.Writer

		for row, err := range csv_r.Iterate() {

			if err != nil {
				slog.Error("Failed to iterate row", "error", err)
				gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
				return
			}

			slog.Debug("Process", "row", row)

			value, ok := row[opts.Marc034Column]

			if !ok {
				slog.Error("Row is missing MARC 034 column", "column", opts.Marc034Column)
				gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
				return
			}

			p, err := fields.Parse034(value)

			if err != nil {
				slog.Error("Failed to parse MARC 034 value", "value", value, "error", err)
				gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
				return
			}

			b, err := p.Bound()

			if err != nil {
				slog.Error("Failed to derive bounds for MARC 034 value", "value", value, "error", err)
				gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
				return
			}

			row[opts.MinXColumn] = strconv.FormatFloat(b.Left(), 'f', -1, 64)
			row[opts.MinYColumn] = strconv.FormatFloat(b.Bottom(), 'f', -1, 64)
			row[opts.MaxXColumn] = strconv.FormatFloat(b.Right(), 'f', -1, 64)
			row[opts.MaxYColumn] = strconv.FormatFloat(b.Top(), 'f', -1, 64)

			if csv_wr == nil {

				wr, err := csvdict.NewWriter(rsp)

				if err != nil {
					slog.Error("Failed to create CSV writer", "error", err)
					gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
					return
				}

				csv_wr = wr
			}

			err = csv_wr.WriteRow(row)

			if err != nil {
				slog.Error("Failed to write CSV row", "error", err)
				gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
				return
			}
		}

		if csv_wr != nil {
			csv_wr.Flush()
		}

		return
	}

	h := gohttp.HandlerFunc(fn)
	return h, nil
}
