package bmat

// Not возвращает матрицу с инвертированными битами.
func (m *Mat) Not() (mRes *Mat) {
	mRes = New(m.width, m.height)
	for row := 0; row < m.height; row++ {
		for col := 0; col < m.widthBytes; col++ {
			mRes.SetByte(row, col, ^m.GetByte(row, col))
		}
	}
	return
}

// And накладывает матрицы от координаты (0,0) используя логическую операцию AND.
// Возвращает матрицу с самой большой шириной и высотой.
func (m *Mat) And(m2 *Mat) (mRes *Mat) {
	return m.eachByte(m2, func(b1, b2 uint8) uint8 {
		return b1 & b2
	})
}

// Or накладывает матрицы от координаты (0,0) используя логическую операцию OR.
// Возвращает матрицу с самой большой шириной и высотой.
func (m *Mat) Or(m2 *Mat) (mRes *Mat) {
	return m.eachByte(m2, func(b1, b2 uint8) uint8 {
		return b1 | b2
	})
}

// Xor накладывает матрицы от координаты (0,0) используя логическую операцию XOR.
// Возвращает матрицу с самой большой шириной и высотой.
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
	for row := 0; row < height; row++ {
		for col := 0; col < widthBytes; col++ {
			mRes.data[num] = fn(m.GetByte(row, col), m2.GetByte(row, col))
			num++
		}
	}

	return mRes
}
