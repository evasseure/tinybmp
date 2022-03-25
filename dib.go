package bmp

import "encoding/binary"

// DIB represent the DIB header of the file
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

func (dib *DIB) fromBytes(bytes []byte) {
	dib.bytes = int(binary.LittleEndian.Uint16(bytes[14:18]))
	dib.width = int(binary.LittleEndian.Uint16(bytes[18:22]))
	dib.height = int(binary.LittleEndian.Uint16(bytes[22:26]))
	dib.colorPlanes = int(binary.LittleEndian.Uint16(bytes[26:28]))
	dib.bitsPerPixel = int(binary.LittleEndian.Uint16(bytes[28:30]))
	dib.compression = int(binary.LittleEndian.Uint16(bytes[30:34]))
	dib.pixelsSize = int(binary.LittleEndian.Uint16(bytes[34:38]))
	dib.resolutionX = int(binary.LittleEndian.Uint16(bytes[38:42]))
	dib.resolutionY = int(binary.LittleEndian.Uint16(bytes[42:46]))
	dib.paletteColors = int(binary.LittleEndian.Uint16(bytes[46:50]))
	dib.importantColors = int(binary.LittleEndian.Uint16(bytes[50:54]))
}
