package bmat

import (
	"fmt"
)

type Mat struct {
	// реальная ширина
	width int
	// по ширине выравнивается до байта
	data [][]uint8
}

func (m Mat) Width() int {
	return m.width
}

func (m *Mat) Height() int {
	return len(m.data)
}
func (m *Mat) Cols() int {
	if len(m.data) == 0 {
		return 0
	}
	return len(m.data[0])
}

func (m *Mat) Rows() int {
	return len(m.data)
}

func (m *Mat) Data() [][]uint8 {
	return m.data
}

func (m *Mat) DataToString() string {
	s := ""
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			s += fmt.Sprintf("%08b ", m.GetByte(row, col))
		}
		s += "\n"
	}
	return s
}

func (m *Mat) GetByte(row, col int) uint8 {
	if m.Cols() <= col || m.Rows() <= row {
		return 0
	}
	return m.data[row][col]
}

func (m *Mat) SetByte(row, col int, b uint8) bool {
	if col < 0 || row < 0 || m.Cols() <= col || m.Rows() <= row {
		return false
	}
	m.data[row][col] = b
	return true
}

func (m *Mat) Clone() (cMat *Mat) {
	cMat = &Mat{
		width: m.Width(),
		data:  make([][]uint8, m.Rows()),
	}

	for i := range m.data {
		cMat.data[i] = append(cMat.data[i], m.data[i]...)
	}
	return
}

func New(width, height int) (m *Mat) {
	m = &Mat{
		width: width,
	}

	rows := width / 8
	if width%8 > 0 {
		rows++
	}

	m.data = make([][]uint8, height)
	for i := range m.data {
		m.data[i] = make([]uint8, rows)
	}
	return
}

func (m *Mat) CountBits() (n int) {
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			n += Bits(m.GetByte(row, col))
		}
	}
	return
}

func (m *Mat) Area(x, y, width, height int) (mRes *Mat) {
	mRes = New(width, height)

	offsetRow := y
	if y < 0 {
		height -= -1 * y
		y = 0
	}

	offsetCol := x / 8
	var offset int
	if x < 0 {
		offset = (-1 * x) & 7
	} else {
		offset = x & 7
	}

	startCol := x / 8
	endCol := startCol + (width-1)/8 + 1
	for row := y; row < y+height; row++ {
		if row >= m.Rows() {
			break
		}

		for col := startCol; col < endCol; col++ {
			if col < 0 {
				continue
			}

			if x >= 0 {
				if col >= m.Cols() {
					break
				}

				mRes.data[row-offsetRow][col-offsetCol] = m.data[row][col] << offset
				if offset > 0 && col+1 < m.Cols() {
					mRes.data[row-offsetRow][col-offsetCol] |= m.data[row][col+1] >> (8 - offset)
				}
				continue
			}

			if col == 0 {
				mRes.data[row-offsetRow][col-offsetCol] = m.data[row][col] >> offset

			} else {
				mRes.data[row-offsetRow][col-offsetCol] = m.data[row][col-1] << (8 - offset)

				if offset > 0 && col < m.Cols() {
					mRes.data[row-offsetRow][col-offsetCol] |= m.data[row][col] >> offset
				}
			}
		}
	}
	return
}

func (m *Mat) TopMargin() (margin int) {
	return m.vertMargin(false)
}

func (m *Mat) BottomMargin() (margin int) {
	return m.vertMargin(true)
}

func (m *Mat) vertMargin(bottom bool) (margin int) {
	for row := 0; row < m.Rows(); row++ {
		r := row
		if bottom {
			r = m.Rows() - row - 1
		}

		for col := 0; col < m.Cols(); col++ {
			if m.data[r][col] != 0 {
				return r
			}
		}
	}

	if !bottom {
		return m.Rows() - 1
	}
	return
}

func (m *Mat) LeftMargin() (margin int) {
	var col int
	for col = 0; col < m.Cols(); col++ {
		if !m.emptyColumn(col) {
			break
		}
	}
	margin = col * 8
	if margin >= m.width {
		return m.width - 1
	}

	// поиск бита
	for x := uint8(128); x > 0; x /= 2 {
		for row := 0; row < m.Rows(); row++ {
			if m.data[row][col]&x > 0 {
				return
			}
		}
		margin++
	}
	return
}

func (m *Mat) RightMargin() (margin int) {
	var col int
	for col = m.Cols() - 1; col > -2; col-- {
		if col < 0 {
			return 0
		}

		if !m.emptyColumn(col) {
			break
		}
	}

	margin = (col + 1) * 8

	// поиск бита
	for x := uint8(1); x < 129; x *= 2 {
		for row := 0; row < m.Rows(); row++ {
			if m.data[row][col]&x > 0 {
				margin--
				return
			}
		}
		margin--
	}
	return
}

func (m *Mat) emptyColumn(col int) bool {
	for row := 0; row < m.Rows(); row++ {
		if m.data[row][col] != 0 {
			return false
		}
	}
	return true
}

func Bits(b uint8) (n int) {
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
