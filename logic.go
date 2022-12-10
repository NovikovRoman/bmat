package bmat

import "math"

// Not возвращает матрицу с инвертированными битами.
func (m *Mat) Not() {
	cols := m.Cols()
	if m.width != cols*8 {
		cols--
		lastCol := cols
		mask := uint8(math.Pow(2, float64(m.width-cols*8)) - 1)
		for row := 0; row < m.Rows(); row++ {
			b := ^m.data[row][lastCol]
			m.data[row][lastCol] = b & ^mask | m.data[row][lastCol]&mask
		}
	}

	if cols > 0 {
		for row := 0; row < m.Rows(); row++ {
			for col := 0; col < cols; col++ {
				m.data[row][col] = ^m.data[row][col]
			}
		}
	}
}

// And накладывает матрицы от координаты (x,y) используя логическую операцию AND.
func (m *Mat) And(m2 *Mat, x, y int) {
	var nb uint
	if x > 0 {
		nb = uint(x)
	} else {
		nb = 8 - uint(-1*x%8)
	}
	tMat := m2.align(nb, true)
	m.eachByteByCoord(tMat, x, y, func(b1, b2 uint8) uint8 {
		return b1 & b2
	})
}

// Or накладывает матрицы от координаты (x,y) используя логическую операцию OR.
func (m *Mat) Or(m2 *Mat, x, y int) {
	var nb uint
	if x > 0 {
		nb = uint(x)
	} else {
		nb = 8 - uint(-1*x%8)
	}
	tMat := m2.align(nb, false)
	m.eachByteByCoord(tMat, x, y, func(b1, b2 uint8) uint8 {
		return b1 | b2
	})
}

// Xor накладывает матрицы от координаты (x,y) используя логическую операцию XOR.
func (m *Mat) Xor(m2 *Mat, x, y int) {
	var nb uint
	if x > 0 {
		nb = uint(x)
	} else {
		nb = 8 - uint(-1*x%8)
	}
	tMat := m2.align(nb, false)
	m.eachByteByCoord(tMat, x, y, func(b1, b2 uint8) uint8 {
		return b1 ^ b2
	})
}

func (m *Mat) eachByteByCoord(m2 *Mat, x, y int, fn func(b1, b2 uint8) uint8) {
	startCol := x / 8
	if x < 0 {
		startCol--
	}

	for row := y; row < y+m2.Rows(); row++ {
		if row < 0 {
			continue
		}

		if row >= m.Rows() {
			break
		}

		for col := startCol; col < startCol+m2.Cols(); col++ {
			if col < 0 {
				continue
			}

			if col >= m.Cols() {
				break
			}

			m.data[row][col] = fn(m.data[row][col], m2.data[row-y][col-startCol])
		}
	}
}

/*
func (m *Mat) eachByte(m2 *Mat, fn func(b1, b2 uint8) uint8) (mRes *Mat) {
	width := m.width
	rows := m.Rows()
	if m2.Rows() > m.Rows() {
		width = m2.width
		rows = m2.Rows()
	}

	cols := m.Cols()
	if m2.Cols() > m.Cols() {
		cols = m2.Cols()
	}

	mRes = New(width, cols)
	for row := 0; row < cols; row++ {
		for col := 0; col < rows; col++ {
			mRes.SetByte(row, col, fn(m.GetByte(row, col), m2.GetByte(row, col)))
		}
	}
	return mRes
}
*/
