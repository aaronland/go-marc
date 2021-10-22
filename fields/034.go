package fields

import (
	"errors"
	"fmt"
	"github.com/paulmach/orb"
	_ "log"
	"regexp"
	"strconv"
	"strings"
)

// https://www.loc.gov/marc/bibliographic/bd034.html

var scales map[string]string
var rings map[string]string
var subfields map[string]string

type Scale struct {
	Code string
}

func (s *Scale) String() string {
	return s.Code
}

type Ring struct {
	Code string
}

func (r *Ring) String() string {
	return r.Code
}

type Subfield struct {
	Code  string
	Value string
}

func (sf *Subfield) String() string {
	return fmt.Sprintf("%s%s", sf.Code, sf.Value)
}

type Parsed struct {
	Scale     *Scale
	Ring      *Ring
	Subfields map[string]*Subfield
}

func (p *Parsed) String() string {

	subfields := make([]string, 0)

	// sudo do this alphabetically...

	for _, s := range p.Subfields {
		subfields = append(subfields, s.String())
	}

	return p.Scale.String() + p.Ring.String() + strings.Join(subfields, "")
}

type Coord struct {
	DD         float64
	Hemisphere string
}

func NewScale(code string) (*Scale, error) {

	_, ok := scales[code]

	if !ok {
		return nil, errors.New("Invalid scale code")
	}

	scale := Scale{Code: code}
	return &scale, nil
}

func NewRing(code string) (*Ring, error) {

	_, ok := rings[code]

	if !ok {
		return nil, errors.New("Invalid ring code")
	}

	ring := Ring{Code: code}
	return &ring, nil
}

func NewSubfield(code string, value string) (*Subfield, error) {

	_, ok := subfields[code]

	if !ok {
		return nil, errors.New("Invalid subfield code")
	}

	sub := Subfield{Code: code, Value: value}
	return &sub, nil
}

func init() {

	scales = map[string]string{
		"0": "Scale indeterminable/No scale recorded",
		"1": "Single scale",
		"2": "Range of scales",
	}

	rings = map[string]string{
		"#": "Not applicable",
		"0": "Outer ring",
		"1": "Exclusion ring",
	}

	subfields = map[string]string{
		"$a": "Category of scale",
		"$b": "Constant ratio linear horizontal scale",
		"$c": "Constant ratio linear vertical scale",
		"$d": "Coordinates - westernmost longitude",
		"$e": "Coordinates - easternmost longitude",
		"$f": "Coordinates - northernmost latitude",
		"$g": "Coordinates - southernmost latitude",
		"$h": "Angular scale",
		"$j": "Declination - northern limit",
		"$k": "Declination - southern limit",
		"$m": "Right ascension - eastern limit",
		"$n": "Right ascension - western limit",
		"$p": "Equinox",
		"$r": "Distance from earth",
		"$s": "G-ring latitude",
		"$t": "G-ring longitude",
		"$x": "Beginning date",
		"$y": "Ending date",
		"$z": "Name of extraterrestrial body",
		"$0": "Authority record control number or standard number",
		"$2": "Source",
		"$3": "Materials specified",
		"$6": "Linkage",
		"$8": "Field link and sequence number",
	}

}

func Parse034(raw string) (*Parsed, error) {

	chars := strings.Split(raw, "")
	count := len(chars)

	if count < 3 {
		return nil, errors.New("E_INSUFFICIENT_034")
	}

	first := chars[0]
	second := chars[1]
	third := chars[2]

	if third != "$" {
		return nil, errors.New("Invalid 034 string")
	}

	scale, err := NewScale(first)

	if err != nil {
		return nil, err
	}

	ring, err := NewRing(second)

	if err != nil {
		return nil, err
	}

	subfields := make(map[string]*Subfield)

	idx := 2

	for {

		if idx >= count {
			break
		}

		buf := make([]string, 0)

		for {

			buf = append(buf, chars[idx])
			idx += 1

			if idx >= count {
				break
			}

			if chars[idx] == "$" {
				break
			}
		}

		code := strings.Join(buf[:2], "")
		value := strings.Join(buf[2:], "")

		sub, err := NewSubfield(code, value)

		if err != nil {
			return nil, err
		}

		subfields[code] = sub
	}

	p := Parsed{
		Scale:     scale,
		Ring:      ring,
		Subfields: subfields,
	}

	return &p, nil
}

func (p *Parsed) Bound() (*orb.Bound, error) {

	coord_w, err := Parse034Coordinate(p.Subfields["$d"].Value, "W")

	if err != nil {
		return nil, err
	}

	coord_e, err := Parse034Coordinate(p.Subfields["$e"].Value, "E")

	if err != nil {
		return nil, err
	}

	coord_n, err := Parse034Coordinate(p.Subfields["$f"].Value, "N")

	if err != nil {
		return nil, err
	}

	coord_s, err := Parse034Coordinate(p.Subfields["$g"].Value, "S")

	if err != nil {
		return nil, err
	}

	minx := coord_w.DD
	miny := coord_s.DD
	maxx := coord_e.DD
	maxy := coord_n.DD

	min := orb.Point{minx, miny}
	max := orb.Point{maxx, maxy}

	b := &orb.Bound{
		Min: min,
		Max: max,
	}

	return b, nil
}

func Parse034Coordinate(raw string, hemisphere string) (*Coord, error) {

	// log.Println("PARSE COORD ", raw)

	// hdddmmss (hemisphere-degrees-minutes-seconds)
	// hddd.dddddd (hemisphere-degrees.decimal degrees)
	// +-ddd.dddddd (hemisphere[+/-]-degrees.decimal degrees) (“+” for N and E, “-“ for S and W; the plus sign is optional)
	// hdddmm.mmmm (hemisphere-degrees-minutes.decimal minutes)
	// hdddmmss.sss (hemisphere-degrees-minutes-seconds.decimal seconds)

	re_hdms, err := regexp.Compile(`^(N|E|S|W)(\d{3})(\d{2})(\d{2})$`)

	if err != nil {
		return nil, err
	}

	if re_hdms.MatchString(raw) {

		m := re_hdms.FindStringSubmatch(raw)

		hem := m[1]
		deg := m[2]
		min := m[3]
		sec := m[4]

		dd, err := dms2dd(deg, min, sec)

		if err != nil {
			return nil, err
		}

		if (hem == "S") || (hem == "W") {
			dd = -dd
		}

		coord := Coord{DD: dd, Hemisphere: hem}
		return &coord, nil
	}

	re_dms, err := regexp.Compile(`^(\d{3})(\d{2})(\d{2})$`)

	if err != nil {
		return nil, err
	}

	if re_dms.MatchString(raw) {

		m := re_dms.FindStringSubmatch(raw)

		deg := m[1]
		min := m[2]
		sec := m[3]

		dd, err := dms2dd(deg, min, sec)

		if err != nil {
			return nil, err
		}

		if (hemisphere == "S") || (hemisphere == "W") {
			dd = -dd
		}

		coord := Coord{DD: dd, Hemisphere: ""}
		return &coord, nil
	}

	re_dd, _ := regexp.Compile(`^(\+|\-)(\d{1,3}\.\d+)$`)

	if err != nil {
		return nil, err
	}

	if re_dd.MatchString(raw) {

		m := re_dd.FindStringSubmatch(raw)

		dd, err := strconv.ParseFloat(m[2], 64)

		if err != nil {
			return nil, err
		}

		if m[1] == "-" {
			dd = -dd
		}

		coord := Coord{DD: dd, Hemisphere: ""}
		return &coord, nil
	}

	return nil, errors.New("Unsupported (or invalid) coordinate string")
}

func dms2dd(deg string, min string, sec string) (float64, error) {

	deg_fl, err := strconv.ParseFloat(deg, 64)

	if err != nil {
		return 0.0, err
	}

	min_fl, err := strconv.ParseFloat(min, 64)

	if err != nil {
		return 0.0, err
	}

	sec_fl, err := strconv.ParseFloat(sec, 64)

	if err != nil {
		return 0.0, err
	}

	var dd float64

	dd = float64(sec_fl) / float64(60)
	dd = dd + float64(min_fl)/float64(60)
	dd = dd + float64(deg_fl)

	return dd, nil
}
