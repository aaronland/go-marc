package http

// SUDO put me in a package

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	gohttp "net/http"
	"os"
	"strings"
)

// mapConfig defines common configuration details for maps.
type MapConfig struct {
	// A valid map provider label.
	Provider string `json:"provider"`
	// A valid Leaflet tile layer URI.
	TileURL string `json:"tile_url"`
	// Optional Protomaps configuration details
	Protomaps       *ProtomapsConfig `json:"protomaps,omitempty"`
	Style           *LeafletStyle    `json:"style,omitempty"`
	PointStyle      *LeafletStyle    `json:"point_style,omitempty"`
	LabelProperties []string         `json:"label_properties"`
}

// protomapsConfig defines configuration details for maps using Protomaps.
type ProtomapsConfig struct {
	// A valid Protomaps theme label
	Theme string `json:"theme"`
}

// LeafletStyle is a struct containing details for decorating GeoJSON features and markers
type LeafletStyle struct {
	Color       string  `json:"color,omitempty"`
	FillColor   string  `json:"fillColor,omitempty"`
	Weight      float64 `json:"weight,omitempty"`
	Opacity     float64 `json:"opacity,omitempty"`
	Radius      float64 `json:"radius,omitempty"`
	FillOpacity float64 `json:"fillOpacity,omitempty"`
}

// UnmarshalStyle derives a `LeafletStyle` instance from 'raw'. If 'raw' starts with "{" then it is treated as
// a JSON-encoded string, otherwise it is treated as a local path on disk.
func UnmarshalStyle(raw string) (*LeafletStyle, error) {

	raw = strings.TrimSpace(raw)

	if len(raw) == 0 {
		return nil, fmt.Errorf("Empty style definition")
	}

	if string(raw[0]) == "{" {
		return UnmarshalStyleFromString(raw)
	}

	r, err := os.Open(raw)

	if err != nil {
		return nil, err
	}

	defer r.Close()

	return UnmarshalStyleFromReader(r)
}

// UnmarshalStyleFromString derives a `LeafletStyle` instance from 'raw'.
func UnmarshalStyleFromString(raw string) (*LeafletStyle, error) {

	var s *LeafletStyle

	err := json.Unmarshal([]byte(raw), &s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

// UnmarshalStyleFromString derives a `LeafletStyle` instance from the body of 'r'.
func UnmarshalStyleFromReader(r io.Reader) (*LeafletStyle, error) {

	var s *LeafletStyle

	dec := json.NewDecoder(r)
	err := dec.Decode(&s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func MapConfigHandler(cfg *MapConfig) gohttp.Handler {

	fn := func(rsp gohttp.ResponseWriter, req *gohttp.Request) {

		rsp.Header().Set("Content-type", "application/json")

		enc := json.NewEncoder(rsp)
		err := enc.Encode(cfg)

		if err != nil {
			slog.Error("Failed to encode map config", "error", err)
			gohttp.Error(rsp, "Internal server error", gohttp.StatusInternalServerError)
		}

		return
	}

	return gohttp.HandlerFunc(fn)
}
