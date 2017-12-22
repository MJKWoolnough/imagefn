package imagefn

import (
	"image"
	"image/color"
)

type flipX struct {
	image.Image
	dx int
}

type flipXSet struct {
	flipX
	setter
}

func FlipX(i image.Image) image.Image {
	switch i := i.(type) {
	case *flipX:
		return i.Image
	case *flipXSet:
		return i.Image
	case *flipY:
		return Rotate180(i.Image)
	case *flipYSet:
		return Rotate180(i.Image)
	case *rotate180:
		return FlipY(i.Image)
	case *rotate180Set:
		return FlipY(i.Image)
	}
	b := i.Bounds()
	f := flipX{
		Image: i,
		dx:    b.Max.X + b.Min.X,
	}
	if s, ok := i.(setter); ok {
		return &flipXSet{
			flipX:  f,
			setter: s,
		}
	}
	return &f
}

func (f *flipX) At(x, y int) color.Color {
	return f.Image.At(f.dx-x-1, y)
}

func (f *flipX) SubImage(r image.Rectangle) image.Image {
	r.Min.X = f.dx - r.Min.X
	r.Max.X = f.dx - r.Max.X
	return FlipX(SubImage(f.Image, r.Canon()))
}

func (f *flipXSet) Set(x, y int, c color.Color) {
	f.setter.Set(f.dx-x-1, y, c)
}

type flipY struct {
	image.Image
	dy int
}

type flipYSet struct {
	flipY
	setter
}

func FlipY(i image.Image) image.Image {
	switch i := i.(type) {
	case *flipX:
		return Rotate180(i.Image)
	case *flipXSet:
		return Rotate180(i.Image)
	case *flipY:
		return i.Image
	case *flipYSet:
		return i.Image
	case *rotate180:
		return FlipX(i.Image)
	case *rotate180Set:
		return FlipX(i.Image)
	}
	b := i.Bounds()
	f := flipY{
		Image: i,
		dy:    b.Max.Y + b.Min.Y,
	}
	if s, ok := i.(setter); ok {
		return &flipYSet{
			flipY:  f,
			setter: s,
		}
	}
	return &f
}

func (f *flipY) At(x, y int) color.Color {
	return f.Image.At(x, f.dy-y-1)
}

func (f *flipY) SubImage(r image.Rectangle) image.Image {
	r.Min.Y = f.dy - r.Min.Y
	r.Max.Y = f.dy - r.Max.Y
	return FlipY(SubImage(f.Image, r.Canon()))
}

func (f *flipYSet) Set(x, y int, c color.Color) {
	f.setter.Set(x, f.dy-y-1, c)
}
