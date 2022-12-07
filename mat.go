package bmat

type Mat struct {
	// реальная ширина
	width int
	// ширина в байтах
	widthBytes int
	height     int
	// по ширине выравнивается до байта
	data []uint8
}

func (m *Mat) Width() int {
	return m.width
}

func (m *Mat) WidthBytes() int {
	return m.widthBytes
}

func (m *Mat) Height() int {
	return m.height
}

func (m *Mat) GetByte(row, col int) uint8 {
	if m.widthBytes <= col || m.height <= row {
		return 0
	}
	return m.data[row*m.widthBytes+col]
}

func (m *Mat) SetByte(row, col int, b uint8) bool {
	if m.widthBytes <= col || m.height <= row {
		return false
	}
	m.data[row*m.widthBytes+col] = b
	return true
}

func (m *Mat) Clone() (cMat *Mat) {
	cMat = &Mat{
		height:     m.height,
		width:      m.width,
		widthBytes: m.widthBytes,
		data:       m.data[:],
	}
	return
}

func New(width, height int) (m *Mat) {
	m = &Mat{
		width:  width,
		height: height,
	}

	m.widthBytes = width / 8
	if width%8 > 0 {
		m.widthBytes++
	}

	m.data = make([]uint8, m.widthBytes*height)
	return
}

func (m *Mat) CountBits() (n int) {
	for _, b := range m.data {
		n += bits(b)
	}
	return
}

func (m *Mat) Area(x, y, width, height int) (mRes *Mat) {
	mRes = New(width, height)

	num := 0
	if y < 0 {
		dy := -1 * y
		num += dy * mRes.widthBytes
		height -= dy
		y = 0
	}

	offset := x & 7
	if x < 0 {
		offset = (-1 * x) & 7
	}

	startCol := getCol(x)
	endCol := startCol + getCol(width-1) + 1
	for row := y; row < y+height; row++ {
		if row >= m.height {
			break
		}

		for col := startCol; col < endCol; col++ {
			if col < 0 {
				num++
				continue
			}

			var b uint8
			if x >= 0 {
				b = m.GetByte(row, col)
				mRes.data[num] = b << offset
				if offset > 0 && col+1 < m.widthBytes {
					mRes.data[num] |= m.GetByte(row, col+1) >> (8 - offset)
				}
				num++
				continue
			}

			if col == 0 {
				b = m.GetByte(row, col)
				mRes.data[num] = b >> offset

			} else {
				b = m.GetByte(row, col-1)
				mRes.data[num] = b << (8 - offset)

				if offset > 0 {
					mRes.data[num] |= m.GetByte(row, col) >> offset
				}
			}

			num++
		}
	}
	return
}

func (m *Mat) TopMargin() (margin int) {
	return m.horizontalMargin(false)
}

func (m *Mat) BottomMargin() (margin int) {
	return m.horizontalMargin(true)
}

func (m *Mat) horizontalMargin(bottom bool) (margin int) {
	if m.CountBits() == 0 { // пустая матрица
		if bottom {
			return 0
		}
		return m.height - 1
	}

	for row := 0; row < m.height; row++ {
		r := row
		if bottom {
			r = m.height - row - 1
		}

		for col := 0; col < m.widthBytes; col++ {
			if m.GetByte(r, col) != 0 {
				return r
			}
		}
	}

	return
}

func (m *Mat) LeftMargin() (margin int) {
	if m.CountBits() == 0 { // пустая матрица
		return m.width - 1
	}

	var col int
loop:
	for col = 0; col < m.widthBytes; col++ {
		for row := 0; row < m.height; row++ {
			if m.GetByte(row, col) != 0 {
				break loop
			}
		}
	}

	margin = col * 8

	// поиск бита
	for x := uint8(128); x > 0; x /= 2 {
		for row := 0; row < m.height; row++ {
			if m.GetByte(row, col)&x > 0 {
				return
			}
		}
		margin++
	}
	return
}

func (m *Mat) RightMargin() (margin int) {
	if m.CountBits() == 0 { // пустая матрица
		return 0
	}

	var col int
loop:
	for col = m.width - 1; col > -1; col-- {
		for row := 0; row < m.height; row++ {
			if m.GetByte(row, col) != 0 {
				break loop
			}
		}
	}

	margin = (col + 1) * 8

	// поиск бита
	for x := uint8(1); x < 129; x *= 2 {
		for row := 0; row < m.height; row++ {
			if m.GetByte(row, col)&x > 0 {
				margin--
				return
			}
		}
		margin--
	}
	return
}

func getCol(x int) int {
	return x / 8
}

func bits(b uint8) (n int) {
	// Единственный случай, когда ответ 0.
	if b == 0 {
		return 0
	}
	// Единственный случай, когда ответ 8.
	if b == 0xFF {
		return 8
	}
	// Считаем число бит по модулю 7.
	n = int((0x010101 * uint32(b) & 0x249249) % 7)
	// Гарантированно имеем 7 единичных битов.
	if n == 0 {
		return 7
	}

	// Случай, когда в числе от 1 до 6 единичных битов.
	return
}
