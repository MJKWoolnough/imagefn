package imagefn // import "vimagination.zapto.org/imagefn"

import (
	"image"
	"image/color"
)

type empty struct {
	color.Model
	Min image.Point
}

func newEmpty(i image.Image) image.Image {
	e := empty{
		Model: i.ColorModel(),
		Min:   i.Bounds().Min,
	}
	if _, ok := i.(setter); ok {
		return &emptySet{
			empty: e,
		}
	}
	return &e
}

func (e empty) ColorModel() color.Model {
	return e.Model
}

func (e empty) Bounds() image.Rectangle {
	return image.Rectangle{e.Min, e.Min}
}

func (e empty) At(x, y int) color.Color {
	return e.Model.Convert(color.Alpha{})
}

type emptySet struct {
	empty
}

func (emptySet) Set(_, _ int, _ color.Color) {}
