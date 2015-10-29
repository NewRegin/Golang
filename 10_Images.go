//10.Images 练习
//未作修改
package main

import (
	"code.google.com/p/go-tour/pic"
	"image"
	"image/color"
)

type Image struct {
	Width, Height int
}

// ColorModel returns the Image's color model.
func (im *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.Width, im.Height)
}

func (im *Image) At(x, y int) color.Color {
	return color.RGBA{128 + uint8(x), 128 + uint8(y), 255, 255}
}

func main() {
	m := &Image{100, 100}
	pic.ShowImage(m)
}
