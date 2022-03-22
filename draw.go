package bmp

var colorR = 0
var colorG = 0
var colorB = 0

// SetColor sets the color used for drawing
func SetColor(r, g, b int) {
	colorR = clampRGB(r)
	colorG = clampRGB(g)
	colorB = clampRGB(b)
}

// DrawPoint sets the pixel with (r, g, b)
func DrawPoint(pixels [][]Pixel, x, y int) {
	pixels[x][y].R = colorR
	pixels[x][y].G = colorG
	pixels[x][y].B = colorB
}

// DrawRect draw a filled rectangle at (x, y) with (w, h) width and height
func DrawRect(pixels [][]Pixel, x, y, w, h int) [][]Pixel {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			pixels[x+i][y+j].R = colorR
			pixels[x+i][y+j].G = colorG
			pixels[x+i][y+j].B = colorB
		}
	}
	return pixels
}
