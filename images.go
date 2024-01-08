package main

import (
	"golang.org/x/tour/pic"
	"image"
    "image/color"
)

type Image struct{
	width int
	height int
}

func pixelGen(x int) uint8 {
	return uint8(x * 2)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{pixelGen(x), pixelGen(y), 255, 255}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}
