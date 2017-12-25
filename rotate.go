package imagefn

import (
	"image"
	"image/color"
)

type rotate90 struct {
	image.Image
	ax, dy int
}

type rotate90Set struct {
	rotate90
	setter
}

func Rotate90(i image.Image) image.Image {
	switch i := i.(type) {
	case *rotate90:
		return Rotate180(i.Image)
	case *rotate90Set:
		return Rotate180(i.Image)
	case *rotate180:
		return Rotate270(i.Image)
	case *rotate180Set:
		return Rotate270(i.Image)
	case *rotate270:
		return i.Image
	case *rotate270Set:
		return i.Image
	}
	b := i.Bounds()
	r := rotate90{
		Image: i,
		ax:    b.Max.Y - b.Min.X, // x = My - mx - x
		dy:    b.Min.X - b.Min.Y, // y = mx - my + y
	}
	if s, ok := i.(setter); ok {
		return &rotate90Set{
			rotate90: r,
			setter:   s,
		}
	}
	return &r
}

func (r rotate90) At(x, y int) color.Color {
	return r.Image.At(r.dy+y, r.ax-x-1)
}

func (r rotate90) Bounds() image.Rectangle {
	b := r.Image.Bounds()
	b.Min.X, b.Min.Y = b.Min.Y, b.Min.X
	b.Max.X, b.Max.Y = b.Max.Y, b.Max.X
	return b
}

func (s rotate90) SubImage(r image.Rectangle) image.Image {
	return s.Image
}

func (r rotate90Set) Set(x, y int, c color.Color) {
	r.setter.Set(r.dy+y, r.ax-x-1, c)
}

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
	case *rotate90:
		return Rotate270(i.Image)
	case *rotate90Set:
		return Rotate270(i.Image)
	case *rotate180:
		return i.Image
	case *rotate180Set:
		return i.Image
	case *rotate90:
		return Rotate90(i.Image)
	case *rotate90Set:
		return Rotate90(i.Image)
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
	return r.Image.At(r.dx-x-1, r.dy-y-1)
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
	r.setter.Set(r.dx-x-1, r.dy-y-1, c)
}

type rotate270 struct {
	image.Image
	dx, ay int
}

type rotate270Set struct {
	rotate270
	setter
}

func Rotate270(i image.Image) image.Image {
	switch i := i.(type) {
	case *rotate90:
		return i.Image
	case *rotate90Set:
		return i.Image
	case *rotate180:
		return Rotate90(i.Image)
	case *rotate180Set:
		return Rotate90(i.Image)
	case *rotate270:
		return Rotate180(i.Image)
	case *rotate270Set:
		return Rotate180(i.Image)
	}
	b := i.Bounds()
	r := rotate270{
		Image: i,
		dx:    b.Min.X - b.Min.Y, // y = mx - my + x
		ay:    b.Max.X - b.Min.X, // x = Mx - mx - y
	}
	if s, ok := i.(setter); ok {
		return &rotate270Set{
			rotate270: r,
			setter:    s,
		}
	}
	return &r
}

func (r rotate270) At(x, y int) color.Color {
	return r.Image.At(r.ay-y-1, r.dx+x)
}

func (r rotate270) Bounds() image.Rectangle {
	b := r.Image.Bounds()
	b.Min.X, b.Min.Y = b.Min.Y, b.Min.X
	b.Max.X, b.Max.Y = b.Max.Y, b.Max.X
	return b
}

func (s rotate270) SubImage(r image.Rectangle) image.Image {
	return s.Image
}

func (r rotate270Set) Set(x, y int, c color.Color) {
	r.setter.Set(r.ay-y-1, r.dx+x, c)
}
