package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	x, y  int
	color Color
}

type Color struct {
	r, g, b, a uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x, i.y) // 最初は0, 0
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.color.r, i.color.g, i.color.b, i.color.a}
}

func main() {
	c := Color{60, 60, 0, 190}
	m := Image{200, 150, c}
	pic.ShowImage(m)
}
