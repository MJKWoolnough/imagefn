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
		R: uint16(a - r),
		G: uint16(a - g),
		B: uint16(a - b),
		A: uint16(a),
	})
}

func (i Invert) SubImage(r image.Rectangle) image.Image {
	return Invert{SubImage(i.Image, r)}
}
