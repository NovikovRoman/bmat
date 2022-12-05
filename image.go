package bmat

import (
	"image"
	"image/color"
)

func FromImage(im image.Image, backColor color.Color) (m *Mat) {
	p0 := im.Bounds().Min
	p1 := im.Bounds().Max

	w := p1.X - p0.X
	h := p1.Y - p0.Y
	m = New(w, h)

	for y := 0; y < h; y++ {
		b := uint8(0)

		for x := 0; x < w; x++ {
			b = b << 1

			if im.At(x+p0.X, y+p0.Y) != backColor {
				b += 1
			}

			if x&7 == 7 {
				m.data[y*m.widthBytes+x/8] = b
				b = 0
			}
		}

		if im.Bounds().Dx()&7 != 0 {
			m.data[y*m.widthBytes+w/8] = b << (m.widthBytes*8 - w)
		}
	}

	return
}

func (m *Mat) ToImage() (im *image.Gray) {
	penColor := color.Gray{Y: 255}
	im = image.NewGray(image.Rect(0, 0, m.width, m.height))

	for row := 0; row < m.height; row++ {
		for col := 0; col < m.widthBytes; col++ {
			b := m.GetByte(col, row)

			for dx := 0; dx < 8; dx++ {
				x := col*8 + dx

				if x >= m.width {
					break
				}

				if b&128 > 0 {
					im.SetGray(x, row, penColor)
				}
				b = b << 1
			}

		}
	}
	return
}
