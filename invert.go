package imagefn

import (
	"image"
	"image/color"
)

type Invert struct {
	image.Image
}

func (i Invert) At(x, y int) color.Color {
	r, g, b, a := i.Image.At(x, y).RGBA()
	return i.Image.ColorModel().Convert(color.RGBA64{
		R: 0xFFFF - r,
		G: 0xFFFF - g,
		B: 0xFFFF - b,
		A: a,
	})
}
