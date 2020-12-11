package main

import (
	"image"
	"image/color"
)

func (g *grid) ColorModel() color.Model {
	return color.GrayModel
}

func (g *grid) Bounds() image.Rectangle {
	return image.Rect(0, 0, g.width, g.height)
}

func (g *grid) At(x, y int) color.Color {
	switch g.data[y*g.width+x] {
	case cellFloor:
		return color.Gray{0}
	case cellOccupied:
		return color.Gray{255}
	case cellEmpty:
		return color.Gray{128}
	}
	return color.Black
}
