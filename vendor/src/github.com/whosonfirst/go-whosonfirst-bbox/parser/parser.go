package parser

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-bbox"	
	"github.com/thisisaaronland/go-marc/fields"
	"strconv"
	"strings"
)

type Parser struct {
	Scheme    string
	Order     string
	Separator string
}

func NewParser() (*Parser, error) {

	p := Parser{
		Scheme:    "cardinal",
		Order:     "swne",
		Separator: ",",
	}

	return &p, nil
}

func (p *Parser) Parse(bbox string) (bbox.BBOX, error) {

	switch p.Scheme {

	case "cardinal":
		return p.ParseCardinal(bbox)
	case "marc":
		return p.ParseMARC(bbox)
	default:
		return nil, errors.New("Invalid or unsupported bounding box scheme")
	}
}

func (p *Parser) ParseCardinal(str_bbox string) (bbox.BBOX, error) {

	var minx float64
	var miny float64
	var maxx float64
	var maxy float64

	var str_miny string
	var str_minx string
	var str_maxy string
	var str_maxx string

	var err error

	parts := strings.Split(str_bbox, ",")

	if len(parts) != 4 {
		return nil, errors.New("Invalid bounding box")
	}

	switch p.Order {

	// SW, NE (lat,lon)
	
	case "swne":

		str_minx = parts[1]
		str_miny = parts[0]
		str_maxx = parts[3]		
		str_maxy = parts[2]

	// SW, NE (lon,lat)
	
	case "wsen":

		str_minx = parts[0]
		str_miny = parts[1]
		str_maxx = parts[2]
		str_maxy = parts[3]

	// NW, SE (lat,lon)
	
	case "nwse":
	
		str_minx = parts[1]	
		str_miny = parts[2]
		str_maxx = parts[3]		
		str_maxy = parts[0]

	// NW, SE (lon,lat)
	
	case "wnes":

	     	str_minx = parts[1]
		str_miny = parts[2]
		str_maxx = parts[3]
		str_maxy = parts[0]
		
	default:
		return nil, errors.New("Unsupported or invalid ordering")
	}

	miny, err = strconv.ParseFloat(str_miny, 64)

	if err != nil {
		return nil, errors.New("Invalid SW latitude parameter")
	}

	minx, err = strconv.ParseFloat(str_minx, 64)

	if err != nil {
		return nil, errors.New("Invalid SW longitude parameter")
	}

	maxy, err = strconv.ParseFloat(str_maxy, 64)

	if err != nil {
		return nil, errors.New("Invalid NE latitude parameter")
	}

	maxx, err = strconv.ParseFloat(str_maxx, 64)

	if err != nil {
		return nil, errors.New("Invalid NE longitude parameter")
	}

	return bbox.NewBoundingBox(minx, miny, maxx, maxy)
}

func (p *Parser) ParseMARC(str_bbox string) (bbox.BBOX, error) {

	parsed, err := fields.Parse034(str_bbox)

	if err != nil {

		msg := fmt.Sprintf("Invalid 034 MARC string %s", err)
		return nil, errors.New(msg)
	}

	_bb, err := parsed.BoundingBox()

	if err != nil {
		return nil, errors.New("Failed to derive bounding box from 034 MARC string")
	}

	// this is to account for the fact that we don't have an interface{} thingy
	// to share across packages yet... (20170220/thisisaaronland)

	minx := _bb.SW.Longitude
	miny := _bb.SW.Latitude
	maxx := _bb.NE.Longitude
	maxy := _bb.NE.Latitude

	return bbox.NewBoundingBox(minx, miny, maxx, maxy)
}
