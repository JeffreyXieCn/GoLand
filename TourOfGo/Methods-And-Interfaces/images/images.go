package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
}

//type Image interface {
//   ColorModel() color.Model
//   Bounds() Rectangle
//   At(x, y int) color.Color
//}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 200)
}

func (img Image) At(x, y int) color.Color {
	//v := uint8(x | y)
	// v:= uint8((x + y) / 2)
	// v:= uint8(x * y)
	// v:= uint8(x & y)
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	m2 := Image{}
	pic.ShowImage(m2)
}
