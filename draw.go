package bmp

// DrawPoint sets the pixel with (r, g, b)
func DrawPoint(pixel *Pixel, r, g, b int) {
	pixel.R = clampRGB(r)
	pixel.G = clampRGB(g)
	pixel.B = clampRGB(b)
}

// DrawRect IT IS MISSING THE COLOR
func DrawRect(pixels [][]Pixel, x, y, w, h int) [][]Pixel {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			pixels[x+i][y+j].B = 255
			pixels[x+i][y+j].G = 0
			pixels[x+i][y+j].R = 0
		}
	}
	return pixels
}
