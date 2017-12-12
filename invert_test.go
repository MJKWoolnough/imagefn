package imagefn

import (
	"image"
	"testing"
)

func TestInvert(t *testing.T) {
	for n, test := range [...]struct {
		In, Out *image.Gray
	}{
		{
			newGray(3, 3, 0, 1, 2, 3, 4, 5, 6, 7, 8),
			newGray(3, 3, 255, 254, 253, 252, 251, 250, 249, 248, 247),
		},
	} {
		if !testImage(&Invert{test.In}, test.Out) {
			t.Errorf("test %d: images do not match", n+1)
		}
	}
}
