package bmp

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

func (dib DIB) asBytes() []byte {
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
	return dibBytes
}
