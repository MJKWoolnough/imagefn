package imagefn

import (
	"image"
	"image/color"
)

type invert struct {
	image.Image
}

type invertSet struct {
	invert
	setter
}

func Invert(i image.Image) image.Image {
	in := invert{
		Image: i,
	}
	if s, ok := i.(setter); ok {
		return invertSet{
			invert: in,
			setter: s,
		}
	}
	return in
}

func (i invert) At(x, y int) color.Color {
	r, g, b, a := i.Image.At(x, y).RGBA()
	return i.Image.ColorModel().Convert(color.RGBA64{
		R: uint16(a - r),
		G: uint16(a - g),
		B: uint16(a - b),
		A: uint16(a),
	})
}

func (i invert) SubImage(r image.Rectangle) image.Image {
	return Invert(SubImage(i.Image, r))
}

func (i invertSet) Set(x, y int, c color.Color) {
	r, g, b, a := c.RGBA()
	i.setter.Set(x, y, color.RGBA64{
		R: uint16(a - r),
		G: uint16(a - g),
		B: uint16(a - b),
		A: uint16(a),
	})
}
