package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	// "strings"
	"encoding/binary"
	"os"
)

type pixel struct {
	r int
	g int
	b int
}

func createEmpty(w, h int) [][]pixel {
	array := make([][]pixel, w)
	for x := 0; x < w; x++ {
		array[x] = make([]pixel, h)
	}
	return array
}

func stringToHex(content string) string {
	return hex.EncodeToString([]byte(content))
}

func clampRGB(value int) int {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return value
}

func randomizeColors(pixels [][]pixel) [][]pixel {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for y := 0; y < len(pixels[0]); y++ {
		for x := 0; x < len(pixels); x++ {
			// offset := (r1.Intn(20)*2 - 20)
			offset := r1.Intn(20) - 10
			// fmt.Println(offset)
			if x == 0 && y == 0 {
				pixels[0][0].b = r1.Intn(255)
				pixels[0][0].g = r1.Intn(255)
				pixels[0][0].r = r1.Intn(255)
			} else if x == 0 {
				pixels[x][y].b = clampRGB((pixels[x][y-1].b) + offset)
				pixels[x][y].g = clampRGB((pixels[x][y-1].g) + offset)
				pixels[x][y].r = clampRGB((pixels[x][y-1].r) + offset)
			} else if y == 0 {
				pixels[x][y].b = clampRGB((pixels[x-1][y].b) + offset)
				pixels[x][y].g = clampRGB((pixels[x-1][y].g) + offset)
				pixels[x][y].r = clampRGB((pixels[x-1][y].r) + offset)
			} else {
				pixels[x][y].b = clampRGB((pixels[x-1][y].b+pixels[x][y-1].b)/2 + offset)
				pixels[x][y].g = clampRGB((pixels[x-1][y].g+pixels[x][y-1].g)/2 + offset)
				pixels[x][y].r = clampRGB((pixels[x-1][y].r+pixels[x][y-1].r)/2 + offset)
			}
		}
	}
	return pixels
}

// BMP ...
type BMP struct {
	ID     string
	size   int
	offset int
}

// DIB ...
type DIB struct {
	bytes           int
	width           int
	height          int
	colorPlanes     int
	bitsPerPixel    int
	compression     int
	pixelsSize      int
	resolutionX     int
	resolutionY     int
	paletteColors   int
	importantColors int
}

func toPadded(value int, padding int) []byte {
	b := make([]byte, padding)
	binary.LittleEndian.PutUint16(b, uint16(value))
	return b
}

func main() {
	var width int = 350
	var height int = 200

	pixelData := createEmpty(width, height)
	pixelData = randomizeColors(pixelData)
	// pixelData = drawRect(pixelData, 10, 10, 50, 100)

	pixels := make([]byte, 0)
	padding := 0
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			pixels = append(pixels, byte(pixelData[x][y].b))
			pixels = append(pixels, byte(pixelData[x][y].g))
			pixels = append(pixels, byte(pixelData[x][y].r))
		}

		len := (width * 3) % 4
		// Adds the padding at the end since it's only RGB
		// https://en.wikipedia.org/wiki/BMP_file_format#:~:text=optional%20color%20list.-,Pixel%20storage,-%5Bedit%5D
		if len != 0 {
			for i := 0; i < 4-len; i++ {
				pixels = append(pixels, 0)
				padding++
			}
		}
	}

	// fmt.Println(pixels)

	rawBitmapDataSize := (4 * width * height) + padding
	dib := DIB{40, width, height, 1, 24, 0, rawBitmapDataSize, 2835, 2835, 0, 0}

	dibBytes := make([]byte, 0)
	dibBytes = append(dibBytes, toPadded(dib.bytes, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.width, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.height, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.colorPlanes, 2)...)
	dibBytes = append(dibBytes, toPadded(dib.bitsPerPixel, 2)...)
	dibBytes = append(dibBytes, toPadded(dib.compression, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.pixelsSize, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.resolutionX, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.resolutionX, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.paletteColors, 4)...)
	dibBytes = append(dibBytes, toPadded(dib.importantColors, 4)...)

	fmt.Println(dibBytes)

	bmp := BMP{"BM", 54 + rawBitmapDataSize, 54}
	bmpBytes := make([]byte, 0)
	bmpBytes = append(bmpBytes, []byte(bmp.ID)...)
	bmpBytes = append(bmpBytes, toPadded(bmp.size, 4)...)
	bmpBytes = append(bmpBytes, toPadded(0, 4)...)
	bmpBytes = append(bmpBytes, toPadded(bmp.offset, 4)...)

	content := hex.EncodeToString(bmpBytes) + hex.EncodeToString(dibBytes) + hex.EncodeToString(pixels)
	decoded, _ := hex.DecodeString(content)
	os.WriteFile("./out.bmp", decoded, 0644)
}
