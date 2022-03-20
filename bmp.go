package bmp

// BMP ...
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
