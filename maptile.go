package maptile

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
)

var TileURL string
var Subdomains []string

type Coords struct {
	X, Y float64
	Zoom int
}

func (c Coords) AsInt() Coords {
	return Coords{math.Floor(c.X), math.Floor(c.Y), c.Zoom}
}

func (c Coords) Url() string {
	s := Subdomains[rand.Intn(len(Subdomains)-1)]
	url := TileURL
	url = strings.Replace(url, "{x}", fmt.Sprintf("%d", int(c.AsInt().X)), -1)
	url = strings.Replace(url, "{y}", fmt.Sprintf("%d", int(c.AsInt().Y)), -1)
	url = strings.Replace(url, "{z}", fmt.Sprintf("%d", c.Zoom), -1)
	url = strings.Replace(url, "{s}", s, -1)

	return url
}

type LatLng struct {
	Lat, Lng float64
}

func (ll LatLng) Tile(zoom int) Coords {
	var c Coords
	c.Zoom = zoom
	c.X = (((ll.Lng + 180) / 360) * math.Pow(2, float64(zoom)))
	c.Y = ((1 - math.Log(math.Tan(deg2rad(ll.Lat))+1/math.Cos(deg2rad(ll.Lat)))/math.Pi) / 2 * math.Pow(2, float64(zoom)))
	return c
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
