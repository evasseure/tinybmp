package tinybmp

import (
	"encoding/binary"
)

// BMP represent the BMP header of the file
type BMP struct {
	ID     string
	size   int
	offset int
}

func (bmp BMP) asBytes() []byte {
	bmpBytes := make([]byte, 0)
	bmpBytes = append(bmpBytes, []byte(bmp.ID)...)
	bmpBytes = append(bmpBytes, toPadded(bmp.size, 4)...)
	bmpBytes = append(bmpBytes, toPadded(0, 4)...)
	bmpBytes = append(bmpBytes, toPadded(bmp.offset, 4)...)
	return bmpBytes
}

func (bmp *BMP) fromBytes(bytes []byte) {
	bmp.ID = string(bytes[:2])
	bmp.size = int(binary.LittleEndian.Uint16(bytes[2:6]))
	bmp.offset = int(binary.LittleEndian.Uint16(bytes[10:14]))
}
