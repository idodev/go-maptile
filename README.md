# maptile golang package

A simple utility package to translate latatitude/longitude values to X/Y tile values with a given zoom level.

This can also be used to generate tile urls.

## Usage Example

```go
  package main

  import (
    "fmt"
    "github.com/idodev/maptile"
  )

  func main() {
    // specify tile url format
    maptile.TileURL = "http://otile{s}.mqcdn.com/tiles/1.0.0/map/{z}/{x}/{y}.png"
    maptile.Subdomains = []string{"0", "1", "2", "3"}
    // create a latlng reference
    latlng := maptile.LatLng{Lat: -5.123, Lng: 40.234}
    // get the tile from the latlng at a specified zoom
    tile := latlng.Tile(14)
    // get the url from this tile
    url := tile.Url()
    // output to console
    fmt.Println(url)
  }
```

Toby Foord

@tobyfoord

http://idodev.co.uk
