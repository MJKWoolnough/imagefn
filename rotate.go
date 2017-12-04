package imagefn

import (
	"image"
	"image/color"
)

type Rotate180 struct {
	image.Image
}

func (r Rotate180) At(x, y int) color.Color {
	b := r.Image.Bounds()
	return r.Image.At(b.Max.X+b.Min.X-x, b.Max.Y+b.Min.Y-y)
}

func (s Rotate180) SubImage(r image.Rectangle) image.Image {
	b := s.Bounds()
	return SubImage(s.Image, image.Rect(
		b.Max.X+b.Min.X-r.Min.X,
		b.Max.X+b.Min.X-r.Max.X,
		b.Max.Y+b.Min.Y-r.Min.Y,
		b.Max.Y+b.Min.Y-r.Max.Y,
	))
}
