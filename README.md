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
    maptile.UrlFormat = "http://otile{s}.mqcdn.com/tiles/1.0.0/map/{z}/{x}/{y}.png"
    maptile.UrlSubdomains = []string{"1", "2", "3", "4"}

    // find a tile from latitude, longitude & zoom
    myTile := maptile.FromLatLng(50.263195, -5.051041, 12)

    // get the url from this tile
    url := myTile.Url()

    // output to console
    fmt.Println(url)

    // set the storage location for tiles
    maptile.TileStore = "C:\\tiles\\"

    // fetch and save the tile image
    myTile.SaveImage()

  }
```

Toby Foord

@tobyfoord

http://idodev.co.uk
