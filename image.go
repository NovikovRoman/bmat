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

func (m *Mat) ToImage() (im *image.Gray) {
	backColor := color.Gray{Y: 0}
	penColor := color.Gray{Y: 255}
	im = image.NewGray(image.Rect(0, 0, m.width, m.height))
	for row := 0; row < m.height; row++ {
		for col := 0; col < m.widthBytes; col++ {
			b := m.GetByte(col, row)
			for x := 0; x < 8; x++ {
				if b&128 > 0 {
					im.SetGray(x, row, penColor)
				} else {
					im.SetGray(x, row, backColor)
				}
			}

		}
	}
	return
}
