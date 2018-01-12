package parser

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-bbox"	
	"github.com/thisisaaronland/go-marc/fields"
	_ "log"
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
		return nil, errors.New(fmt.Sprintf("Invalid or unsupported bounding box scheme '%s'", p.Scheme))
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
		return nil, errors.New(fmt.Sprintf("Invalid bounding box '%s'", str_bbox))
	}

	switch p.Order {

	// SW, NE (lat,lon)
	
	case "swne":

		str_minx = strings.Trim(parts[1], " ")
		str_miny = strings.Trim(parts[0], " ")
		str_maxx = strings.Trim(parts[3], " ")
		str_maxy = strings.Trim(parts[2], " ")

	// SW, NE (lon,lat)
	
	case "wsen":
	
		str_minx = strings.Trim(parts[0], " ")
		str_miny = strings.Trim(parts[1], " ")
		str_maxx = strings.Trim(parts[2], " ")
		str_maxy = strings.Trim(parts[3], " ")

	// NW, SE (lat,lon)
	
	case "nwse":
	
		str_minx = strings.Trim(parts[1], " ")	
		str_miny = strings.Trim(parts[2], " ")
		str_maxx = strings.Trim(parts[3], " ")		
		str_maxy = strings.Trim(parts[0], " ")
		
	default:
		return nil, errors.New("Unsupported or invalid ordering")
	}

	miny, err = strconv.ParseFloat(str_miny, 64)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid SW latitude parameter '%s'", miny))
	}

	minx, err = strconv.ParseFloat(str_minx, 64)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid SW longitude parameter '%s'", str_minx))
	}

	maxy, err = strconv.ParseFloat(str_maxy, 64)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid NE latitude parameter '%s'", str_maxy))
	}

	maxx, err = strconv.ParseFloat(str_maxx, 64)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid NE longitude parameter '%s'", str_maxx))
	}

	return bbox.NewBoundingBox(minx, miny, maxx, maxy)
}

func (p *Parser) ParseMARC(str_bbox string) (bbox.BBOX, error) {

	parsed, err := fields.Parse034(str_bbox)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid 034 MARC string '%s' : %s", str_bbox, err))
	}

	bb, err := parsed.BoundingBox()

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to determine bounding box for MARC string '%s' : %s", str_bbox, err))
	}

	return bb, nil
}
