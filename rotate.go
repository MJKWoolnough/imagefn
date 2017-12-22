package imagefn

import (
	"image"
	"image/color"
)

type rotate180 struct {
	image.Image
	dx, dy int
}

type rotate180Set struct {
	rotate180
	setter
}

func Rotate180(i image.Image) image.Image {
	switch i := i.(type) {
	case *flipX:
		return FlipY(i.Image)
	case *flipXSet:
		return FlipY(i.Image)
	case *flipY:
		return FlipX(i.Image)
	case *flipYSet:
		return FlipX(i.Image)
	case *rotate180:
		return i.Image
	case *rotate180Set:
		return i.Image
	}
	b := i.Bounds()
	r := rotate180{
		Image: i,
		dx:    b.Max.X + b.Min.X,
		dy:    b.Max.Y + b.Min.Y,
	}
	if s, ok := i.(setter); ok {
		return &rotate180Set{
			rotate180: r,
			setter:    s,
		}
	}
	return &r
}

func (r rotate180) At(x, y int) color.Color {
	return r.Image.At(r.dx-x, r.dy-y)
}

func (s rotate180) SubImage(r image.Rectangle) image.Image {
	return SubImage(s.Image, image.Rect(
		s.dx-r.Min.X,
		s.dy-r.Min.Y,
		s.dx-r.Max.X,
		s.dy-r.Max.Y,
	))
}

func (r rotate180Set) Set(x, y int, c color.Color) {
	r.setter.Set(r.dx-x, r.dy-y, c)
}
