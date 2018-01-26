package imagefn

import (
	"image"
	"image/color"
)

type scale struct {
	image.Image
	ScaleX, ScaleY float64
}

func Scale(i image.Image, xScale, yScale float64) image.Image {
	switch im := i.(type) {
	case *empty:
		return i
	case *scale:
		xScale *= im.ScaleX
		yScale *= im.ScaleY
		i = im.Image
	}
	if xScale < 0 {
		if yScale < 0 {
			i = FlipX(i)
		} else {
			i = Rotate180(i)
			yScale = -yScale
		}
		xScale = -xScale
	} else if yScale < 0 {
		i = FlipY(i)
		yScale = -yScale
	}
	if xScale == 0 || yScale == 0 {
		return newEmpty(i)
	}
	if xScale == 1 && yScale == 1 {
		return i
	}
	return &scale{
		Image:  i,
		ScaleX: xScale,
		ScaleY: yScale,
	}
}

func ScaleDimensions(i image.Image, x, y int) image.Image {
	b := i.Bounds()
	return Scale(i, float64(x)/float64(b.Dx()), float64(y)/float64(b.Dy()))
}

func (s *scale) Bounds() image.Rectangle {
	b := s.Image.Bounds()
	return image.Rect(
		b.Min.X,
		b.Min.Y,
		b.Min.X+int(float64(b.Dx())/s.ScaleX),
		b.Min.Y+int(float64(b.Dy())/s.ScaleY),
	)
}

func (s *scale) At(x, y int) color.Color {
	b := s.Image.Bounds()
	return s.Image.At(
		b.Min.X+int(float64(x-b.Min.X)/s.ScaleX),
		b.Min.Y+int(float64(y-b.Min.Y)/s.ScaleY),
	)
}

type Scaler interface {
}

type smoothScale struct {
	scale
	scaler Scaler
}

func SmoothScale(i image.Image, xScale, yScale float64, scale Scale) image.Image {
	switch s := Scale(i, xScale, yScale).(type) {
	case *scale:
		return smoothScale{s, scaler}
	default:
		return s
	}
}

func SmoothScaleDimensions(i image.Image, x, y int, scaler Scaler) image.Image {
	b := i.Bounds()
	return SmoothScale(i, float64(x)/float64(b.Dx()), float64(y)/float64(b.Dy()), scaler)
}

func (s *smoothScale) At(x, y int) color.Color {
	return nil
}
