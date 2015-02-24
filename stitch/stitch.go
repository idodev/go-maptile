package stitch

import (
	"github.com/idodev/maptile"
	"math"
	"sync"
)

var wg = new(sync.WaitGroup)

type abc struct {
	X, Y float64
	Zoom int
}

type Stitch struct {
	NorthEast maptile.MapTile
	SouthWest maptile.MapTile
}

func (s Stitch) GetTilesX() []int {
	minX := math.Min(float64(s.SouthWest.FloorX()), float64(s.NorthEast.FloorX()))
	maxX := math.Max(float64(s.SouthWest.FloorX()), float64(s.NorthEast.FloorX()))

	var xTiles []int
	for x := minX; x <= maxX; x++ {
		xTiles = append(xTiles, int(x))
	}
	return xTiles
}

func (s Stitch) GetTilesY() []int {
	minY := math.Min(float64(s.SouthWest.FloorY()), float64(s.NorthEast.FloorY()))
	maxY := math.Max(float64(s.SouthWest.FloorY()), float64(s.NorthEast.FloorY()))

	var yTiles []int
	for y := minY; y <= maxY; y++ {
		yTiles = append(yTiles, int(y))
	}
	return yTiles
}

func (s Stitch) SaveAllImages() bool {

	xTiles := s.GetTilesX()
	yTiles := s.GetTilesY()

	for x := xTiles[0]; x <= xTiles[len(xTiles)-1]; x++ {
		for y := yTiles[0]; y <= yTiles[len(yTiles)-1]; y++ {
			wg.Add(1)
			go loadImageAsync(float64(x), float64(y), s.SouthWest.Zoom)
		}
	}

	// Waiting for all goroutines to finish (otherwise they die as main routine dies)
	wg.Wait()

	return true
}

func New(NE, SW maptile.MapTile) Stitch {
	s := Stitch{NE, SW}
	return s
}

func loadImageAsync(x, y float64, z int) {
	defer wg.Done()

	tile := maptile.New(x, y, z)
	tile.SaveImage()
}
