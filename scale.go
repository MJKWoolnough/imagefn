package imagefn

import (
	"image"
	"image/color"
)

type Scale struct {
	image.Image
	ScaleX, ScaleY float64
}

func (s *Scale) Bounds() image.Rectangle {
	b := s.Image.Bounds()
	return image.Rect(
		b.Min.X,
		b.Min.Y,
		b.Min.X+int(float64(b.Dx())*s.ScaleX),
		b.Min.Y+int(float64(b.Dy())*s.ScaleY),
	)
}

func (s *Scale) At(x, y int) color.Color {
	b := s.Image.Bounds()
	return s.Image.At(
		b.Min.X+int(float64(x-b.Min.X)*s.ScaleX),
		b.Min.Y+int(float64(y-b.Min.Y)*s.ScaleY),
	)
}
