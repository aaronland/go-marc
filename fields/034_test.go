package fields

import (
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
	t.Skip()
}
