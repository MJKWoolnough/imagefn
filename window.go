package imagefn

import (
	"image"
	"image/color"
)

type setter interface {
	Set(int, int, color.Color)
}

type subimage interface {
	SubImage(image.Rectangle) image.Image
}

func SubImage(i image.Image, r image.Rectangle) image.Image {
	if si, ok := i.(subimage); ok {
		return si.SubImage(r)
	}
	if s, ok := i.(setter); ok {
		return &windowSet{
			window: window{
				Image:     i,
				Rectangle: r,
			},
			setter: s,
		}
	}
	return &window{
		Image:     i,
		Rectangle: r,
	}
}

type window struct {
	image.Image
	Rectangle image.Rectangle
}

func (w *window) At(x, y int) color.Color {
	if !image.Pt(x, y).In(w.Rectangle) {
		return color.Transparent
	}
	return w.Image.At(x, y)
}

func (w *window) Bounds() image.Rectangle {
	return w.Rectangle
}

func (w *window) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(w.Rectangle)
	return &window{
		Image:     w.Image,
		Rectangle: r,
	}
}

type windowSet struct {
	window
	setter
}

func (w *windowSet) Set(x, y int, c color.Color) {
	if !image.Pt(x, y).In(w.Rectangle) {
		return
	}
	w.setter.Set(x, y, c)
}
