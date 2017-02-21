package bbox

import (
	"errors"
	"fmt"
)

type COORD interface {
	Latitude() float64
	Longitude() float64
}

type BBOX interface {
	South() float64
	North() float64
	East() float64
	West() float64
	MinX() float64
	MinY() float64
	MaxX() float64
	MaxY() float64
	SouthWest() COORD
	NorthEast() COORD
}

func NewCoordinate(x float64, y float64) (*Coordinate, error) {

	c := Coordinate{
		latitude:  y,
		longitude: x,
	}

	return &c, nil
}

type Coordinate struct {
	COORD
	latitude  float64
	longitude float64
	// order string
}

func (c *Coordinate) Latitude() float64 {
	return c.latitude
}

func (c *Coordinate) Longitude() float64 {
	return c.longitude
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%0.6f %0.6f", c.Latitude(), c.Longitude())
}

func NewBoundingBox(minx float64, miny float64, maxx float64, maxy float64) (*BoundingBox, error) {

	if miny > 90.0 || miny < -90.0 {
		return nil, errors.New("E_IMPOSSIBLE_LATITUDE (MINY)")
	}

	if maxy > 90.0 || maxy < -90.0 {
		return nil, errors.New("E_IMPOSSIBLE_LATITUDE (MAXY)")
	}

	if minx > 180.0 || minx < -180.0 {
		return nil, errors.New("E_IMPOSSIBLE_LONGITUDE (MINX)")
	}

	if maxx > 180.0 || maxx < -180.0 {
		return nil, errors.New("E_IMPOSSIBLE_LONGITUDE (MAXX)")
	}

	if miny > maxy {
		return nil, errors.New("E_INVALID_LATITUDE (MINY > MAXY)")
	}

	if minx > maxx {
		return nil, errors.New("E_INVALID_LATITUDE (MINX > MAXX)")
	}

	sw, err := NewCoordinate(minx, miny)

	if err != nil {
		return nil, err
	}

	ne, err := NewCoordinate(maxx, maxy)

	if err != nil {
		return nil, err
	}

	bb := BoundingBox{
		sw: sw,
		ne: ne,
	}

	return &bb, nil
}

type BoundingBox struct {
	BBOX
	sw *Coordinate
	ne *Coordinate
}

func (b *BoundingBox) South() float64 {
	return b.sw.Latitude()
}

func (b *BoundingBox) West() float64 {
	return b.sw.Longitude()
}

func (b *BoundingBox) North() float64 {
	return b.ne.Latitude()
}

func (b *BoundingBox) East() float64 {
	return b.ne.Longitude()
}

func (b *BoundingBox) MinX() float64 {
	return b.West()
}

func (b *BoundingBox) MaxX() float64 {
	return b.East()
}

func (b *BoundingBox) MinY() float64 {
	return b.South()
}

func (b *BoundingBox) MaxY() float64 {
	return b.North()
}

func (b *BoundingBox) SouthWest() COORD {

	c := Coordinate{
		latitude:  b.South(),
		longitude: b.West(),
	}

	return &c
}

func (b *BoundingBox) NorthEast() COORD {

	c := Coordinate{
		latitude:  b.North(),
		longitude: b.East(),
	}

	return &c
}

func (b *BoundingBox) String() string {
	return fmt.Sprintf("%s %s", b.sw.String(), b.ne.String())
}
