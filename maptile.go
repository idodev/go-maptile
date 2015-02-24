package maptile

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var UrlFormat string
var UrlSubdomains []string
var TileStore string

type MapTile struct {
	X, Y float64
	Zoom int
}

// methods

func (t MapTile) FloorX() int {
	return int(math.Floor(t.X))
}

func (t MapTile) FloorY() int {
	return int(math.Floor(t.Y))
}

func (t MapTile) Url() string {
	s := UrlSubdomains[rand.Intn(len(UrlSubdomains)-1)]
	url := UrlFormat
	url = strings.Replace(url, "{x}", fmt.Sprintf("%d", t.FloorX()), -1)
	url = strings.Replace(url, "{y}", fmt.Sprintf("%d", t.FloorY()), -1)
	url = strings.Replace(url, "{z}", fmt.Sprintf("%d", t.Zoom), -1)
	url = strings.Replace(url, "{s}", s, -1)

	return url
}
func (t MapTile) Filename() string {
	filename := "Tile-{x}-{y}-{z}.jpg"
	filename = strings.Replace(filename, "{x}", fmt.Sprintf("%d", t.FloorX()), -1)
	filename = strings.Replace(filename, "{y}", fmt.Sprintf("%d", t.FloorY()), -1)
	filename = strings.Replace(filename, "{z}", fmt.Sprintf("%d", t.Zoom), -1)

	return filename
}

func (t MapTile) GetImage() image.Image {
	res, err := http.Get(t.Url())

	if err != nil {
		fmt.Println(err)
		return nil
	}

	if res.StatusCode != 200 {
		fmt.Println(res.Status)
		return nil
	}
	defer res.Body.Close()
	img, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error decoding map tile")
		return nil
	}

	return img
}

func (t MapTile) SaveImage() bool {
	img := t.GetImage()

	toimg, err := os.Create(TileStore + t.Filename())
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer toimg.Close()

	var opt jpeg.Options
	opt.Quality = 80
	err = jpeg.Encode(toimg, img, &opt)

	return err == nil
}

// constructors

func New(X, Y float64, Zoom int) MapTile {
	return MapTile{X, Y, Zoom}
}

func FromLatLng(Lat, Lng float64, Zoom int) MapTile {
	var t MapTile
	t.Zoom = Zoom
	t.X = (((Lng + 180) / 360) * math.Pow(2, float64(Zoom)))
	t.Y = ((1 - math.Log(math.Tan(deg2rad(Lat))+1/math.Cos(deg2rad(Lat)))/math.Pi) / 2 * math.Pow(2, float64(Zoom)))
	return t
}

// utilities

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
