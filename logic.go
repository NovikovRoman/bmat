package bmat

func (m *Mat) Not() (mRes *Mat) {
	mRes = New(m.width, m.height)
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.widthBytes; x++ {
			mRes.SetByte(x, y, ^m.GetByte(x, y))
		}
	}
	return
}

func (m *Mat) And(m2 *Mat) (mRes *Mat) {
	return m.eachByte(m2, func(b1, b2 uint8) uint8 {
		return b1 & b2
	})
}

func (m *Mat) Or(m2 *Mat) (mRes *Mat) {
	return m.eachByte(m2, func(b1, b2 uint8) uint8 {
		return b1 | b2
	})
}

func (m *Mat) Xor(m2 *Mat) (mRes *Mat) {
	return m.eachByte(m2, func(b1, b2 uint8) uint8 {
		return b1 ^ b2
	})
}

func (m *Mat) eachByte(m2 *Mat, fn func(b1, b2 uint8) uint8) (mRes *Mat) {
	width := m.width
	widthBytes := m.widthBytes
	if m2.widthBytes > m.widthBytes {
		width = m2.width
		widthBytes = m2.widthBytes
	}

	height := m.height
	if m2.height > m.height {
		height = m2.height
	}

	mRes = New(width, height)
	num := 0
	for y := 0; y < height; y++ {
		for x := 0; x < widthBytes; x++ {
			mRes.data[num] = fn(m.GetByte(x, y), m2.GetByte(x, y))
			num++
		}
	}

	return mRes
}
