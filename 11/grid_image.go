package main

import (
	"image"
	"image/color"
)

func (g *grid) ColorModel() color.Model {
	return color.RGBAModel
}

func (g *grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, g.width, g.height)
}

func (g *grid) At(x, y int) color.Color {
	switch g.data[y*g.width+x] {
	case cellFloor:
		return color.RGBA{R: 15, G: 56, B: 15, A: 255}
	case cellOccupied:
		return color.RGBA{R: 48, G: 98, B: 48, A: 255}
	case cellEmpty:
		return color.RGBA{R: 155, G: 188, B: 15, A: 255}
	}
	return color.Black
}
