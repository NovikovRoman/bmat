package bmat

import (
	"reflect"
	"testing"
)

func Test_ByteShift(t *testing.T) {
	m := &Mat{
		width: 22,
		data: [][]uint8{
			{0, 0, 0},
			// Последние два бита в каждой строке не учитываются, тк длина 22.
			{0b01101110, 0b00001110, 0b00101000},
			{0b11001110, 0b00101010, 0b11000000},
			{0b01001110, 0b00001110, 0b10010000},
			{0b10010000, 0b01001110, 0b11001100},
			{0, 0, 0},
		},
	}
	tests := []struct {
		name     string
		numBytes int
		wantMat  *Mat
	}{
		{
			name:     "0",
			numBytes: 0,
			wantMat:  m,
		},
		{
			name:     "<-1",
			numBytes: -1,
			wantMat: &Mat{
				width: m.width,
				data: [][]uint8{
					{0, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0b00001110, 0b00101000, 0},
					{0b00101010, 0b11000000, 0},
					{0b00001110, 0b10010000, 0},
					{0b01001110, 0b11001100, 0},
					{0, 0, 0},
				},
			},
		},
		{
			name:     "->1",
			numBytes: 1,
			wantMat: &Mat{
				width: m.width,
				data: [][]uint8{
					{0, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0, 0b01101110, 0b00001110},
					{0, 0b11001110, 0b00101010},
					{0, 0b01001110, 0b00001110},
					{0, 0b10010000, 0b01001110},
					{0, 0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMat := m.Clone()
			if gotMat.ByteShift(tt.numBytes); !reflect.DeepEqual(gotMat, tt.wantMat) {
				t.Errorf("align() = %v, want %v", gotMat, tt.wantMat)
			}
		})
	}
}

func Test_align(t *testing.T) {
	m := &Mat{
		width: 22,
		data: [][]uint8{
			{0, 0, 0},
			// Последние два бита в каждой строке не учитываются, тк длина 22.
			{0b01101110, 0b00001110, 0b00101000},
			{0b11001110, 0b00101010, 0b11000000},
			{0b01001110, 0b00001110, 0b10010000},
			{0b10010000, 0b01001110, 0b11001100},
			{0, 0, 0},
		},
	}
	tests := []struct {
		name    string
		x       uint
		filling bool
		wantMat *Mat
	}{
		{
			name:    "x 0",
			x:       0,
			filling: false,
			wantMat: m,
		},
		{
			name:    "x 8",
			x:       0,
			filling: false,
			wantMat: m,
		},
		{
			name:    "x 3",
			x:       3,
			filling: false,
			wantMat: &Mat{
				width: m.width + 3,
				data: [][]uint8{
					{0, 0, 0, 0},
					// 0-6 бит в последнем байте каждой строки не учитываются, тк длина 25.
					{0b00001101, 0b11000001, 0b11000101, 0},
					{0b00011001, 0b11000101, 0b01011000, 0},
					{0b00001001, 0b11000001, 0b11010010, 0},
					{0b00010010, 0b00001001, 0b11011001, 0b10000000},
					{0, 0, 0, 0},
				},
			},
		},
		{
			name:    "x 3 filling",
			x:       3,
			filling: true,
			wantMat: &Mat{
				width: m.width + 3,
				data: [][]uint8{
					// 6,5 бит в последнем байте каждой строки не учитываются,
					// тк длина исходной матрицы была 22 + 3 бита смещение.
					// 0-4 бит в последнем байте каждой строки заполняются единицами.
					{0b11100000, 0, 0, 0b00011111},
					{0b11101101, 0b11000001, 0b11000101, 0b00011111},
					{0b11111001, 0b11000101, 0b01011000, 0b00011111},
					{0b11101001, 0b11000001, 0b11010010, 0b00011111},
					{0b11110010, 0b00001001, 0b11011001, 0b10011111},
					{0b11100000, 0, 0, 0b00011111},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMat := m.align(tt.x, tt.filling); !reflect.DeepEqual(gotMat, tt.wantMat) {
				t.Errorf("align() = %v, want %v", gotMat, tt.wantMat)
			}
		})
	}
}
