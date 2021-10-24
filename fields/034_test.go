package fields

import (
	"bytes"
	_ "fmt"
	"strconv"
	"strings"
	"testing"
)

func TestParse034(t *testing.T) {

	tests := map[string]string{
		"1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000": "-82.91666666666667,26.833333333333332,-82.33333333333333,27.5",
	}

	for raw, expected_bbox := range tests {

		p, err := Parse034(raw)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", raw, err)
		}

		if p.String() != raw {
			t.Fatalf("Invalid stringification for '%s': '%s'", raw, p.String())
		}

		b, err := p.Bound()

		if err != nil {
			t.Fatalf("Failed to derive bounding box for '%s', %v", raw, err)
		}

		bbox := []string{
			strconv.FormatFloat(b.Left(), 'f', -1, 64),
			strconv.FormatFloat(b.Bottom(), 'f', -1, 64),
			strconv.FormatFloat(b.Right(), 'f', -1, 64),
			strconv.FormatFloat(b.Top(), 'f', -1, 64),
		}

		str_bbox := strings.Join(bbox, ",")

		if str_bbox != expected_bbox {
			t.Fatalf("Unexpected bounding box for '%s', got '%s' but expected '%s'", raw, str_bbox, expected_bbox)
		}
	}
}

func TestParse034Coordinate(t *testing.T) {

	tests := map[string]string{
		"W0825500 W": "-82.91666666666667 W",
		"W0822000 E": "-82.33333333333333 W",
		"N0273000 N": "27.5 N",
		"N0265000 S": "26.833333333333332 N",
	}

	for raw, expected := range tests {

		parts := strings.Split(raw, " ")
		coord := parts[0]
		hemisphere := parts[1]

		c, err := Parse034Coordinate(coord, hemisphere)

		if err != nil {
			t.Fatalf("Failed to parse '%s' (%s), %v", coord, hemisphere, err)
		}

		str_c := c.String()

		if str_c != expected {
			t.Fatalf("Unexpected stringification for '%s' (%s), got '%s' but expected '%s'", coord, hemisphere, str_c, expected)
		}
	}

}

func TestParse034AsGeoJSON(t *testing.T) {

	tests := map[string]string{
		"1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000": `{"type":"Feature","bbox":[-82.91666666666667,26.833333333333332,-82.33333333333333,27.5],"geometry":{"type":"Polygon","coordinates":[[[-82.91666666666667,26.833333333333332],[-82.33333333333333,26.833333333333332],[-82.33333333333333,27.5],[-82.91666666666667,27.5],[-82.91666666666667,26.833333333333332]]]},"properties":{"marc:034":"1#$aa$b80000$dW0825500$eW0822000$fN0273000$gN0265000"}}`,
	}

	for raw, expected_json := range tests {

		p, err := Parse034(raw)

		if err != nil {
			t.Fatalf("Failed to parse '%s', %v", raw, err)
		}

		f, err := p.AsGeoJSON()

		if err != nil {
			t.Fatalf("Failed to generate GeoJSON for '%s', %v", raw, err)
		}

		enc, err := f.MarshalJSON()

		if err != nil {
			t.Fatalf("Failed to marshal GeoJSON for '%s', %v", raw, err)
		}

		if !bytes.Equal(enc, []byte(expected_json)) {
			t.Fatalf("Failed to generate expected GeoJSON output for '%s', %v", raw, err)
		}

	}
}
