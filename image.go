package bmat

import (
	"image"
	"image/color"
)

func FromImage(im image.Image, backColor color.Color) (m *Mat) {
	h := im.Bounds().Dy()
	m = New(im.Bounds().Dx(), h)

	for y := 0; y < im.Bounds().Dy(); y++ {
		b := uint8(0)

		for x := 0; x < im.Bounds().Dx(); x++ {
			b = b << 1

			if im.At(x, y) != backColor {
				b += 1
			}

			if x&7 == 7 {
				m.data[y*m.widthBytes+x/8] = b
				b = 0
			}
		}

		if im.Bounds().Dx()&7 != 0 {
			m.data[y*m.widthBytes+im.Bounds().Dx()/8] = b << (m.widthBytes*8 - im.Bounds().Dx())
		}
	}

	return
}
