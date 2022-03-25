package tinybmp

import (
	"encoding/binary"
	"encoding/hex"
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
