package bmat

import (
	"image"
	"image/color"
)

func FromImage(im image.Image, backColor color.Color) (m *Mat) {
	offset := im.Bounds().Min
	w := im.Bounds().Dx()
	h := im.Bounds().Dy()
	m = New(w, h)

	var b uint8
	for row := 0; row < h; row++ {
		b = 0

		for x := 0; x < w; x++ {
			b <<= 1

			if im.At(x+offset.X, row+offset.Y) != backColor {
				b += 1
			}

			if x&7 == 7 {
				m.SetByte(row, x/8, b)
				b = 0
			}
		}

		if w&7 != 0 {
			m.SetByte(row, w/8, b<<(m.Cols()*8-w))
		}
	}
	return
}

func (m *Mat) ToImage() (im *image.Gray) {
	penColor := color.Gray{Y: 255}
	im = image.NewGray(image.Rect(0, 0, m.Width(), m.Height()))

	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			b := m.GetByte(row, col)

			for dx := 0; dx < 8; dx++ {
				x := col*8 + dx

				if x >= m.width {
					break
				}

				if b&128 > 0 {
					im.SetGray(x, row, penColor)
				}
				b <<= 1
			}

		}
	}
	return
}
