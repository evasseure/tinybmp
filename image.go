package bmp

import (
	"encoding/hex"
	"os"
)

// Pixel ...
type Pixel struct {
	R int
	G int
	B int
}

// Image ...
type Image struct {
	Width  int
	Height int
	Pixels [][]Pixel
}

// NewImage creates a new Image with width and height (w, h)
func NewImage(w, h int) Image {
	image := Image{}
	image.Width = w
	image.Height = h
	image.createEmpty(w, h)

	return image
}

func (image *Image) createEmpty(w, h int) {
	array := make([][]Pixel, w)
	for x := 0; x < w; x++ {
		array[x] = make([]Pixel, h)
	}
	image.Pixels = array
}

func (image *Image) asBytes() ([]byte, int) {
	pixels := make([]byte, 0)
	padding := 0
	for y := image.Height - 1; y >= 0; y-- {
		for x := 0; x < image.Width; x++ {
			pixels = append(pixels, byte(image.Pixels[x][y].B))
			pixels = append(pixels, byte(image.Pixels[x][y].G))
			pixels = append(pixels, byte(image.Pixels[x][y].R))
		}

		len := (image.Width * 3) % 4
		// Adds the padding at the end since it's only RGB
		// https://en.wikipedia.org/wiki/BMP_file_format#:~:text=optional%20color%20list.-,Pixel%20storage,-%5Bedit%5D
		if len != 0 {
			for i := 0; i < 4-len; i++ {
				pixels = append(pixels, 0)
				padding++
			}
		}
	}

	return pixels, padding
}

// Open ...
func (image *Image) Open(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	bmp := BMP{}
	bmp.fromBytes(data)

	if bmp.ID != "BM" {
		panic("Invalid BMP ID:" + bmp.ID)
	}

	dib := DIB{}
	dib.fromBytes(data)

	array := make([][]Pixel, dib.width)
	for x := 0; x < dib.width; x++ {
		array[x] = make([]Pixel, dib.height)
	}

	for i, x, y := 0, 0, dib.height-1; y >= 0; {

		rgb := data[54+i : 54+i+3]

		// Little endian
		array[x][y].R = int(rgb[2])
		array[x][y].G = int(rgb[1])
		array[x][y].B = int(rgb[0])
		i += 3
		x++

		if x >= dib.width {
			x = 0
			y--
			for i%4 != 0 {
				i++
			}
		}
	}

	image.Pixels = array
	image.Width = dib.width
	image.Height = dib.height
}

// Save save the image in a file
func (image Image) Save(filename string) {
	pixels, padding := image.asBytes()

	rawBitmapDataSize := (4 * image.Width * image.Height) + padding
	dib := DIB{40, image.Width, image.Height, 1, 24, 0, rawBitmapDataSize, 2835, 2835, 0, 0}
	dibBytes := dib.asBytes()

	bmp := BMP{"BM", 54 + rawBitmapDataSize, 54}
	bmpBytes := bmp.asBytes()

	content := hex.EncodeToString(bmpBytes) + hex.EncodeToString(dibBytes) + hex.EncodeToString(pixels)
	decoded, _ := hex.DecodeString(content)
	os.WriteFile(filename, decoded, 0644)
}
