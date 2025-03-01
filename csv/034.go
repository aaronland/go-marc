package csv

import (
	"io"
	"log/slog"
	"strconv"

	"github.com/aaronland/go-marc/v2/fields"
	"github.com/sfomuseum/go-csvdict/v2"
)

func Convert034(r io.Reader, wr io.Writer, marc034_column string) error {

	csv_r, err := csvdict.NewReader(r)

	if err != nil {
		slog.Error("Failed to create new CSV reader", "error", err)
		return err
	}

	var csv_wr *csvdict.Writer

	for row, err := range csv_r.Iterate() {

		if err != nil {
			slog.Error("Failed to iterate row", "error", err)
			return err
		}

		slog.Debug("Process", "row", row)

		row["min_x"] = ""
		row["min_y"] = ""
		row["max_x"] = ""
		row["max_y"] = ""
		row["valid"] = "0"
		row["error"] = ""

		value, ok := row[marc034_column]

		if !ok {
			slog.Error("Row is missing MARC 034 column", "column", marc034_column)
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
