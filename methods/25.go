package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	h   int
	w   int
	img [][]color.RGBA
}

func NewImage(h,w int) Image {
	img := make([][]color.RGBA, h*w)
	for i := 0; i < h; i++ {
		img[i] = make([]color.RGBA, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			img[i][j] = color.RGBA{uint8(i*j/2), uint8(i*j/2), 255, 255}
		}
	}

	return Image{h, w, img}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	return i.img[x][y]
}

func main() {
	m := NewImage(100,100)
	pic.ShowImage(m)
}
