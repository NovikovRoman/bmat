package bmat

import (
	"math"
)

// ByteShift сдвигает матрицу по байтам по горизонтали. > 0 вправо, < 0 влево.
func (m *Mat) ByteShift(numBytes int) {
	if numBytes == 0 {
		return
	}

	dirLeft := false
	if numBytes < 0 {
		dirLeft = true
		numBytes *= -1
	}

	for row := 0; row < m.Rows(); row++ {
		if dirLeft {
			m.data[row] = m.data[row][numBytes:]
			m.data[row] = append(m.data[row], make([]uint8, numBytes)...)
		} else {
			r := m.data[row][:len(m.data[row])-1]
			m.data[row] = make([]uint8, numBytes)
			m.data[row] = append(m.data[row], r...)
		}
	}
}

// align выравнивает края до байта. filling - заполнить края единицами.
// От x берется остаток от деления на 8.
//
// Например: x = 3 значит каждая строка будет смещена на 3 бита вправо.
// Ширина в байтах увеличится на 1 байт, а реальная ширина на 3.
func (m *Mat) align(x uint, filling bool) (aMat *Mat) {
	x = x % 8
	if x == 0 {
		return m.Clone()
	}

	aMat = New(m.width+int(x), m.Height())
	mask := uint8(math.Pow(2, float64(x)) - 1)

	var b uint8
	for row := 0; row < m.Rows(); row++ {
		for col := 0; col < m.Cols(); col++ {
			if col == 0 {
				b = 0
				if filling {
					b = 255 << (8 - x)
				}
				aMat.data[row][col] = b | m.data[row][col]>>x
				continue
			}

			b = m.data[row][col-1] & mask << (8 - x)
			aMat.data[row][col] = b | m.data[row][col]>>x
		}

		b = 0
		if filling {
			b = uint8(math.Pow(2, float64(8-x)) - 1)
		}
		aMat.data[row][aMat.Cols()-1] = b | ((m.data[row][m.Cols()-1] & mask) << (8 - x))
	}
	return
}
