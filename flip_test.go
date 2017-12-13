package imagefn

import (
	"image"
	"image/color"
	"testing"
)

func testImage(a, b image.Image) bool {
	ab := a.Bounds()
	bb := b.Bounds()
	dx := ab.Dx()
	dy := ab.Dy()
	if dx != bb.Dx() || dy != bb.Dy() {
		return false
	}
	for j := 0; j < dy; j++ {
		for i := 0; i < dx; i++ {
			ar, ag, ab, aa := a.At(ab.Min.X+i, ab.Min.Y+j).RGBA()
			br, bg, bb, ba := b.At(bb.Min.X+i, bb.Min.Y+j).RGBA()
			if ar != br || ag != bg || ab != bb || aa != ba {
				return false
			}
		}
	}
	return true
}

func newGray(w, h int, pix ...uint8) *image.Gray {
	g := image.NewGray(image.Rect(0, 0, w, h))
Loop:
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			g.SetGray(x, y, color.Gray{pix[0]})
			pix = pix[1:]
			if len(pix) == 0 {
				break Loop
			}
		}
	}
	return g
}

func TestFlipX(t *testing.T) {
	for n, test := range [...]struct {
		In, Out  *image.Gray
		SubImage image.Rectangle
	}{
		{
			newGray(1, 1, 1),
			newGray(1, 1, 1),
			image.Rect(0, 0, 1, 1),
		},
		{
			newGray(1, 2, 0, 1),
			newGray(1, 2, 0, 1),
			image.Rect(0, 1, 1, 2),
		},
		{
			newGray(2, 1, 0, 1),
			newGray(2, 1, 1, 0),
			image.Rect(0, 0, 1, 1),
		},
		{
			newGray(3, 3, 0, 1, 2, 3, 4, 5, 6, 7, 8),
			newGray(3, 3, 2, 1, 0, 5, 4, 3, 8, 7, 6),
			image.Rect(1, 1, 3, 3),
		},
	} {
		fx := FlipX{test.In}
		if !testImage(&fx, test.Out) {
			t.Errorf("test %d: images do not match", n+1)
		}
		if !testImage(fx.SubImage(test.SubImage), test.Out.SubImage(test.SubImage)) {
			t.Errorf("test %d: subimages do not match", n+1)
		}
	}
}

func TestFlipY(t *testing.T) {
	for n, test := range []struct {
		In, Out  *image.Gray
		SubImage image.Rectangle
	}{
		{
			newGray(1, 1, 1),
			newGray(1, 1, 1),
			image.Rect(0, 0, 1, 1),
		},
		{
			newGray(2, 1, 0, 1),
			newGray(2, 1, 0, 1),
			image.Rect(0, 1, 2, 1),
		},
		{
			newGray(1, 2, 0, 1),
			newGray(1, 2, 1, 0),
			image.Rect(0, 0, 1, 1),
		},
		{
			newGray(3, 3, 0, 1, 2, 3, 4, 5, 6, 7, 8),
			newGray(3, 3, 6, 7, 8, 3, 4, 5, 0, 1, 2),
			image.Rect(1, 1, 3, 3),
		},
	} {
		fy := FlipY{test.In}
		if !testImage(&fy, test.Out) {
			t.Errorf("test %d: images do not match", n+1)
		}
		if !testImage(fy.SubImage(test.SubImage), test.Out.SubImage(test.SubImage)) {
			t.Errorf("test %d: subimages do not match", n+1)
		}
	}
}
