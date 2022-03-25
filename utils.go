package bmp

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

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

func toPadded(value int, padding int) []byte {
	b := make([]byte, padding)
	binary.LittleEndian.PutUint16(b, uint16(value))
	return b
}

func checkIndexes(pixels [][]Pixel, x, y int) {
	width := len(pixels)
	height := len(pixels[0])
	if x > width || x > height {
		msg := fmt.Sprintf("out of bounds: invalid values (x:%d, y:%d) for size (%d, %d)", x, y, width, height)
		panic(msg)
	}
}
