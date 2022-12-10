package bmat

import (
	"reflect"
	"testing"
)

var (
	testLogicM_8x1 = &Mat{
		width: 8,
		data:  [][]uint8{{0b00001111}},
	}
	testLogicM_8x2 = &Mat{
		width: 8,
		data:  [][]uint8{{0b11110000}, {0b00001111}},
	}
	testLogicM_12x2 = &Mat{
		width: 12,
		data:  [][]uint8{{0b11001100, 0b11000000}, {0b00110011, 0b00110000}},
	}
	testLogicM_22x6 = &Mat{
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
)

func TestMat_Not(t *testing.T) {
	tests := []struct {
		name    string
		mat     *Mat
		wantRes *Mat
	}{
		{
			name: "8x1",
			mat:  testLogicM_8x1,
			wantRes: &Mat{
				width: 8,
				data:  [][]uint8{{0b11110000}},
			},
		},
		{
			name: "8x2",
			mat:  testLogicM_8x2,
			wantRes: &Mat{
				width: 8,
				data:  [][]uint8{{0b00001111}, {0b11110000}},
			},
		},
		{
			name: "12x2",
			mat:  testLogicM_12x2,
			wantRes: &Mat{
				width: 12,
				// последние 4 бита каждой строки не учавствуют в операции, тк реальная длина 12
				data: [][]uint8{{0b00110011, 0b00110000}, {0b11001100, 0b11000000}},
			},
		},
		{
			name: "12x2",
			mat: &Mat{
				width: 12,
				data:  [][]uint8{{0b11001100, 0b11001001}, {0b00110011, 0b00110110}},
			},
			wantRes: &Mat{
				width: 12,
				// последние 4 бита каждой строки не учавствуют в операции, тк реальная длина 12
				data: [][]uint8{{0b00110011, 0b00111001}, {0b11001100, 0b11000110}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := tt.mat.Clone()
			if gotRes.Not(); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Mat.Not() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestMat_And(t *testing.T) {
	tests := []struct {
		name     string
		mat1     *Mat
		mat2     *Mat
		x        int
		y        int
		wantMRes *Mat
	}{
		{
			name: "x:3 y:2",
			mat1: testLogicM_22x6.Clone(),
			x:    3,
			y:    2,
			mat2: &Mat{
				width: 8,
				data: [][]uint8{
					{0b01101110},
					{0b11001110},
					{0b01001110},
					{0b10010000},
				},
			},
			wantMRes: &Mat{
				width: testLogicM_22x6.width,
				data: [][]uint8{
					{0, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0b01101110, 0b00001110, 0b00101000},
					{0b11001100, 0b00001010, 0b11000000},
					{0b01001000, 0b00001110, 0b10010000},
					{0b10000000, 0b01001110, 0b11001100},
					{0, 0, 0},
				},
			},
		},
		{
			name: "x:-3 y:-2",
			mat1: testLogicM_22x6.Clone(),
			x:    -3,
			y:    -2,
			mat2: &Mat{
				width: 8,
				data: [][]uint8{
					{0b01101110},
					{0b11001110},
					{0b01001110},
					{0b10010000},
				},
			},
			wantMRes: &Mat{
				width: testLogicM_22x6.width,
				data: [][]uint8{
					{0, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0b00000110, 0b00001110, 0b00101000},
					{0b11001110, 0b00101010, 0b11000000},
					{0b01001110, 0b00001110, 0b10010000},
					{0b10010000, 0b01001110, 0b11001100},
					{0, 0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mat1.And(tt.mat2, tt.x, tt.y); !reflect.DeepEqual(tt.mat1, tt.wantMRes) {
				t.Errorf("Mat.AndByCoord() = %v, want %v", tt.mat1, tt.wantMRes)
			}
		})
	}
}

func TestMat_Or(t *testing.T) {
	tests := []struct {
		name    string
		mat1    *Mat
		mat2    *Mat
		x       int
		y       int
		wantRes *Mat
	}{
		{
			name: "x:3 y:2",
			mat1: testLogicM_22x6.Clone(),
			x:    3,
			y:    2,
			mat2: &Mat{
				width: 8,
				data: [][]uint8{
					{0b01101110},
					{0b11001110},
					{0b01001110},
					{0b10010000},
				},
			},
			wantRes: &Mat{
				width: testLogicM_22x6.width,
				data: [][]uint8{
					{0, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0b01101110, 0b00001110, 0b00101000},
					{0b11001111, 0b11101010, 0b11000000},
					{0b01011111, 0b11001110, 0b10010000},
					{0b10011001, 0b11001110, 0b11001100},
					{0b00010010, 0, 0},
				},
			},
		},
		{
			name: "x:-3 y:-2",
			mat1: testLogicM_22x6.Clone(),
			x:    -3,
			y:    -2,
			mat2: &Mat{
				width: 8,
				data: [][]uint8{
					{0b01101110},
					{0b11001110},
					{0b01001110},
					{0b10010000},
				},
			},
			wantRes: &Mat{
				width: testLogicM_22x6.width,
				data: [][]uint8{
					{0b01110000, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0b11101110, 0b00001110, 0b00101000},
					{0b11001110, 0b00101010, 0b11000000},
					{0b01001110, 0b00001110, 0b10010000},
					{0b10010000, 0b01001110, 0b11001100},
					{0, 0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mat1.Or(tt.mat2, tt.x, tt.y); !reflect.DeepEqual(tt.mat1, tt.wantRes) {
				t.Errorf("Mat.OrByCoord() = %v, want %v", tt.mat1, tt.wantRes)
			}
		})
	}
}

func TestMat_Xor(t *testing.T) {
	tests := []struct {
		name    string
		mat1    *Mat
		mat2    *Mat
		x       int
		y       int
		wantRes *Mat
	}{
		{
			name: "x:3 y:2",
			mat1: testLogicM_22x6.Clone(),
			x:    3,
			y:    2,
			mat2: &Mat{
				width: 8,
				data: [][]uint8{
					{0b01101110},
					{0b11001110},
					{0b01001110},
					{0b10010000},
				},
			},
			wantRes: &Mat{
				width: testLogicM_22x6.width,
				data: [][]uint8{
					{0, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0b01101110, 0b00001110, 0b00101000},
					{0b11000011, 0b11101010, 0b11000000},
					{0b01010111, 0b11001110, 0b10010000},
					{0b10011001, 0b10001110, 0b11001100},
					{0b00010010, 0, 0},
				},
			},
		},
		{
			name: "x:-3 y:-2",
			mat1: testLogicM_22x6.Clone(),
			x:    -3,
			y:    -2,
			mat2: &Mat{
				width: 8,
				data: [][]uint8{
					{0b01101110},
					{0b11001110},
					{0b01001110},
					{0b10010000},
				},
			},
			wantRes: &Mat{
				width: testLogicM_22x6.width,
				data: [][]uint8{
					{0b01110000, 0, 0},
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					{0b11101110, 0b00001110, 0b00101000},
					{0b11001110, 0b00101010, 0b11000000},
					{0b01001110, 0b00001110, 0b10010000},
					{0b10010000, 0b01001110, 0b11001100},
					{0, 0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mat1.Xor(tt.mat2, tt.x, tt.y); !reflect.DeepEqual(tt.mat1, tt.wantRes) {
				t.Errorf("Mat.XorByCoord() = %v, want %v", tt.mat1, tt.wantRes)
			}
		})
	}
}
