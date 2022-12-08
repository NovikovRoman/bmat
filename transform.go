package bmat

import (
	"math"
)

// ByteShift сдвигает матрицу по байтам по горизонтали.
func (m *Mat) ByteShift(numBytes int) (aMat *Mat) {
	if numBytes == 0 {
		return m.Clone()
	}

	var add []uint8
	if numBytes > 0 {
		add = make([]uint8, numBytes)

	} else {
		add = make([]uint8, -1*numBytes)
	}

	aMat = New(m.width, m.height)
	aMat.data = []uint8{}
	for row := 0; row < m.height; row++ {
		var (
			startIdx int
			endIdx   int
		)
		if numBytes > 0 {
			startIdx = row * m.widthBytes
			endIdx = startIdx + m.widthBytes - len(add)
			aMat.data = append(aMat.data, add...)
			aMat.data = append(aMat.data, m.data[startIdx:endIdx]...)
			continue
		}

		startIdx = -1*numBytes + row*m.widthBytes
		endIdx = startIdx + m.widthBytes - 1
		aMat.data = append(aMat.data, m.data[startIdx:endIdx]...)
		aMat.data = append(aMat.data, add...)
	}
	return
}

// align выравнивает края до байта. filling - заполнить края единицами.
// От x берется остаток от деления на 8.
//
// Например: x = 3 значит каждая строка будет смещена на 3 бита вправо.
// Ширина в байтах увеличится на 1 байт, а реальная ширина на 3.
func align(m *Mat, x uint, filling bool) (aMat *Mat) {
	x = x % 8
	if x == 0 {
		return m.Clone()
	}

	aMat = New(m.width+int(x), m.height)
	mask := uint8(math.Pow(2, float64(x)) - 1)

	var b uint8
	for row := 0; row < m.height; row++ {
		for col := 0; col < m.widthBytes; col++ {
			if col == 0 {
				b = 0
				if filling {
					b = 255 << (8 - x)
				}

				aMat.SetByte(row, col, b|m.GetByte(row, col)>>x)
				continue
			}

			b = m.GetByte(row, col-1) & mask << (8 - x)
			aMat.SetByte(row, col, b|m.GetByte(row, col)>>x)
		}

		b = 0
		if filling {
			b = uint8(math.Pow(2, float64(8-x)) - 1)
		}
		aMat.SetByte(row, aMat.widthBytes-1, b|((m.GetByte(row, m.widthBytes-1)&mask)<<(8-x)))
	}
	return
}
