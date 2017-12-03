package imagefn

import (
	"image"
	"image/color"
)

type FlipX struct {
	image.Image
}

func (f *FlipX) At(x, y int) color.Color {
	b := f.Bounds()
	return f.Image.At(b.Max.X+b.Min.X-x, y)
}

func (f *FlipX) SubImage(r image.Rectangle) image.Image {
	b := f.Bounds()
	r.Min.X = b.Max.X + b.Min.X - r.Min.X
	r.Max.X = b.Max.X + b.Min.X - r.Max.X
	return &FlipX{
		Image: SubImage(f.Image, r.Canon()),
	}
}

type FlipY struct {
	image.Image
}

func (f *FlipY) At(x, y int) color.Color {
	b := f.Bounds()
	return f.Image.At(x, b.Max.Y+b.Min.Y-y)
}

func (f *FlipY) SubImage(r image.Rectangle) image.Image {
	b := f.Bounds()
	r.Min.Y = b.Max.Y + b.Min.Y - r.Min.Y
	r.Max.Y = b.Max.Y + b.Min.Y - r.Max.Y
	return &FlipY{
		Image: SubImage(f.Image, r.Canon()),
	}
}
