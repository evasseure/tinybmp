package bmp

import (
	"math"
)

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
	checkIndexes(pixels, x, y)
	pixels[x][y].R = colorR
	pixels[x][y].G = colorG
	pixels[x][y].B = colorB
}

// DrawRect draw a empty rectangle at (x, y) with (w, h) width and height
func DrawRect(pixels [][]Pixel, x, y, w, h int) [][]Pixel {
	pixels = DrawLine(pixels, x, y, x+w, y)
	pixels = DrawLine(pixels, x+w, y, x+w, y+h)
	pixels = DrawLine(pixels, x+w, y+h, x, y+h)
	pixels = DrawLine(pixels, x, y+h, x, y)
	return pixels
}

// DrawFilledRect draw a filled rectangle at (x, y) with (w, h) width and height
func DrawFilledRect(pixels [][]Pixel, x, y, w, h int) [][]Pixel {
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			pixels[x+i][y+j].R = colorR
			pixels[x+i][y+j].G = colorG
			pixels[x+i][y+j].B = colorB
		}
	}
	return pixels
}

// DrawLine draw a line from (x1, y1) to (x2, y2)
// Taken from https://jstutorial.medium.com/how-to-code-your-first-algorithm-draw-a-line-ca121f9a1395
func DrawLine(pixels [][]Pixel, x1, y1, x2, y2 int) [][]Pixel {
	var x, y, xe, ye int

	// Calculate line deltas
	dx := x2 - x1
	dy := y2 - y1

	// Create a positive copy of deltas (makes iterating easier)
	dx1 := math.Abs(float64(dx))
	dy1 := math.Abs(float64(dy))

	// Calculate error intervals for both axis
	px := 2*dy1 - dx1
	py := 2*dx1 - dy1

	// The line is X-axis dominant
	if dy1 <= dx1 {
		// Line is drawn left to right
		if dx >= 0 {
			x = x1
			y = y1
			xe = x2
		} else { // Line is drawn right to left (swap ends)
			x = x2
			y = y2
			xe = x1
		}
		DrawPoint(pixels, x, y)

		// Rasterize the line
		for i := 0; x < xe; i++ {
			x = x + 1
			// Deal with octants...
			if px < 0 {
				px = px + 2*dy1
			} else {
				if (dx < 0 && dy < 0) || (dx > 0 && dy > 0) {
					y = y + 1
				} else {
					y = y - 1
				}
				px = px + 2*(dy1-dx1)
			}
			// Draw pixel from line span at
			// currently rasterized position
			DrawPoint(pixels, x, y)
		}
	} else { // The line is Y-axis dominant
		// Line is drawn bottom to top
		if dy >= 0 {
			x = x1
			y = y1
			ye = y2
		} else { // Line is drawn top to bottom
			x = x2
			y = y2
			ye = y1
		}
		DrawPoint(pixels, x, y)
		// Rasterize the line
		for i := 0; y < ye; i++ {
			y = y + 1
			// Deal with octants...
			if py <= 0 {
				py = py + 2*dx1
			} else {
				if (dx < 0 && dy < 0) || (dx > 0 && dy > 0) {
					x = x + 1
				} else {
					x = x - 1
				}
				py = py + 2*(dx1-dy1)
			}
			// Draw pixel from line span at
			// currently rasterized position
			DrawPoint(pixels, x, y)
		}
	}

	return pixels
}
