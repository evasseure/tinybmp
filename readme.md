# TinyBMP

A simple **BMP creation** and edition go library.  
ℹ️ For now this is only a toy project.

## Features

- creation of BMP files from scratch,
- simple primitives drawing,
  - point,
  - line,
  - rectangle,
  - circle,
- set color,
- reading of the simplest type of BMP possible (so it probably won't work with a random .bmp you found).

## Usage

Creating and editing a BMP file:

```go
import (
	"github.com/evasseure/tinybmp"
)

img := tinybmp.NewImage(800, 600)

tinybmp.SetColor(255, 255, 255)
tinybmp.DrawFilledRect(img.Pixels, 0, 0, 800, 600)

tinybmp.SetColor(0, 0, 255)
tinybmp.DrawRect(img.Pixels, 10, 10, 800 - 20, 600 - 20)

img.Save("white_square.bmp")
```

Opening and editing a BMP file:

```go
import (
	"github.com/evasseure/tinybmp"
)

img := tinybmp.Open("./image.bmp")

tinybmp.SetColor(0, 255, 0)
tinybmp.DrawPoint(img.Pixels, 10, 10)

img.Save("edited.bmp")
```

## "Documentation"

### Image

```golang
type Image struct {
	Width  int
	Height int
	Pixels [][]Pixel
}

func (image Image) Save(filename string)
    Save save the image in a file

```

### Image creation

```golang
func NewImage(w, h int) Image
    NewImage creates a new Image with width and height (w, h)

func Open(filename string) Image
    Open a file at path "filename" and returns an Image
```

### Drawing functions

```golang
func SetColor(r, g, b int)
    SetColor sets the color used for drawing

func DrawFilledRect(pixels [][]Pixel, x, y, w, h int) [][]Pixel
    DrawFilledRect draw a filled rectangle at (x, y) with (w, h) width and height

func DrawLine(pixels [][]Pixel, x1, y1, x2, y2 int) [][]Pixel
    DrawLine draw a line from (x1, y1) to (x2, y2)

func DrawPoint(pixels [][]Pixel, x, y int)
    DrawPoint sets the pixel with (r, g, b)

func DrawRect(pixels [][]Pixel, x, y, w, h int) [][]Pixel
    DrawRect draw a empty rectangle at (x, y) with (w, h) width and height
```
