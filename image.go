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
	width  int
	height int
	Pixels [][]Pixel
}

// NewImage creates a new Image with width and height (w, h)
func NewImage(w, h int) Image {
	image := Image{}
	image.width = w
	image.height = h
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
	for y := image.height - 1; y >= 0; y-- {
		for x := 0; x < image.width; x++ {
			pixels = append(pixels, byte(image.Pixels[x][y].B))
			pixels = append(pixels, byte(image.Pixels[x][y].G))
			pixels = append(pixels, byte(image.Pixels[x][y].R))
		}

		len := (image.width * 3) % 4
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

// Save ...
func (image Image) Save(filename string) {
	pixels, padding := image.asBytes()

	rawBitmapDataSize := (4 * image.width * image.height) + padding
	dib := DIB{40, image.width, image.height, 1, 24, 0, rawBitmapDataSize, 2835, 2835, 0, 0}
	dibBytes := dib.asBytes()

	bmp := BMP{"BM", 54 + rawBitmapDataSize, 54}
	bmpBytes := bmp.asBytes()

	content := hex.EncodeToString(bmpBytes) + hex.EncodeToString(dibBytes) + hex.EncodeToString(pixels)
	decoded, _ := hex.DecodeString(content)
	os.WriteFile(filename, decoded, 0644)
}
