package imagefn

import (
	"image"
	"image/color"
)

type empty struct {
	color.Model
	Min image.Point
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

func (empty) Set(_, _ int, _ color.Color) {}
