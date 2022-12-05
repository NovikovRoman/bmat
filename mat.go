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

func (m *Mat) GetByte(col, row int) uint8 {
	if m.widthBytes <= col || m.height <= row {
		return 0
	}
	return m.data[row*m.widthBytes+col]
}

func (m *Mat) SetByte(col, row int, b uint8) bool {
	if m.widthBytes <= col || m.height <= row {
		return false
	}
	m.data[row*m.widthBytes+col] = b
	return true
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
	offset := x & 7
	startCol := getCol(x)
	endCol := startCol + getCol(width-1) + 1
	for row := y; row < y+height; row++ {
		if row >= m.height {
			break
		}

		for col := startCol; col < endCol; col++ {
			b := m.GetByte(col, row)
			mRes.data[num] = b << offset
			if offset > 0 && col+1 < m.widthBytes {
				mRes.data[num] |= m.GetByte(col+1, row) >> (8 - offset)
			}
			num++
		}
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
