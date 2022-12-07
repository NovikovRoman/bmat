package bmat

import (
	"reflect"
	"testing"
)

func TestMat_And(t *testing.T) {
	m1_8x1 := &Mat{
		width:      8,
		widthBytes: 1,
		height:     1,
		data:       []uint8{0b00001111},
	}
	m2_8x1 := &Mat{
		width:      8,
		widthBytes: 1,
		height:     1,
		data:       []uint8{0b00000111},
	}
	m_8x2 := &Mat{
		width:      8,
		widthBytes: 1,
		height:     2,
		data:       []uint8{0b11110000, 0b00001111},
	}
	m_12x2 := &Mat{
		width:      12,
		widthBytes: 2,
		height:     2,
		data:       []uint8{0b11001100, 0b11000000, 0b00110011, 0b00110000},
	}
	tests := []struct {
		name    string
		m1      *Mat
		m2      *Mat
		logic   string
		wantRes *Mat
	}{
		{
			name:  "NOT",
			m1:    m1_8x1,
			m2:    nil,
			logic: "not",
			wantRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     1,
				data:       []uint8{0b11110000},
			},
		},
		{
			name:  "одинаковые AND",
			m1:    m1_8x1,
			m2:    m2_8x1,
			logic: "and",
			wantRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     1,
				data:       []uint8{0b00000111},
			},
		},
		{
			name:  "одинаковые OR",
			m1:    m1_8x1,
			m2:    m2_8x1,
			logic: "or",
			wantRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     1,
				data:       []uint8{0b00001111},
			},
		},
		{
			name:  "одинаковые XOR",
			m1:    m1_8x1,
			m2:    m2_8x1,
			logic: "xor",
			wantRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     1,
				data:       []uint8{0b00001000},
			},
		},
		{
			name:  "разная высота AND",
			m1:    m1_8x1,
			m2:    m_8x2,
			logic: "and",
			wantRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     2,
				data:       []uint8{0, 0},
			},
		},
		{
			name:  "разная высота OR",
			m1:    m1_8x1,
			m2:    m_8x2,
			logic: "or",
			wantRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     2,
				data:       []uint8{0b11111111, 0b00001111},
			},
		},
		{
			name:  "разная высота XOR",
			m1:    m1_8x1,
			m2:    m_8x2,
			logic: "xor",
			wantRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     2,
				data:       []uint8{0b11111111, 0b00001111},
			},
		},
		{
			name:  "разные ширина и высота AND",
			m1:    m1_8x1,
			m2:    m_12x2,
			logic: "and",
			wantRes: &Mat{
				width:      12,
				widthBytes: 2,
				height:     2,
				data:       []uint8{0b00001100, 0, 0, 0},
			},
		},
		{
			name:  "разные ширина и высота OR",
			m1:    m1_8x1,
			m2:    m_12x2,
			logic: "or",
			wantRes: &Mat{
				width:      12,
				widthBytes: 2,
				height:     2,
				data:       []uint8{0b11001111, 0b11000000, 0b00110011, 0b00110000},
			},
		},
		{
			name:  "разные ширина и высота XOR",
			m1:    m1_8x1,
			m2:    m_12x2,
			logic: "xor",
			wantRes: &Mat{
				width:      12,
				widthBytes: 2,
				height:     2,
				data:       []uint8{0b11000011, 0b11000000, 0b00110011, 0b00110000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.logic {
			case "and":
				if gotRes := tt.m1.And(tt.m2); !reflect.DeepEqual(gotRes, tt.wantRes) {
					t.Errorf("Mat.And() = %v, want %v", gotRes, tt.wantRes)
				}
			case "or":
				if gotRes := tt.m1.Or(tt.m2); !reflect.DeepEqual(gotRes, tt.wantRes) {
					t.Errorf("Mat.Or() = %v, want %v", gotRes, tt.wantRes)
				}
			case "xor":
				if gotRes := tt.m1.Xor(tt.m2); !reflect.DeepEqual(gotRes, tt.wantRes) {
					t.Errorf("Mat.Xor() = %v, want %v", gotRes, tt.wantRes)
				}

			case "not":
				if gotRes := tt.m1.Not(); !reflect.DeepEqual(gotRes, tt.wantRes) {
					t.Errorf("Mat.Not() = %v, want %v", gotRes, tt.wantRes)
				}

			default:
				t.Errorf("Unknown logic %s", tt.logic)
			}
		})
	}
}

var (
	testLogicMat = &Mat{
		width:      22,
		height:     6,
		widthBytes: 3,
		data: []uint8{
			0, 0, 0,
			// Последние два бита в каждой строке не учитываются, тк длина 22.
			0b01101110, 0b00001110, 0b00101000,
			0b11001110, 0b00101010, 0b11000000,
			0b01001110, 0b00001110, 0b10010000,
			0b10010000, 0b01001110, 0b11001100,
			0, 0, 0,
		},
	}
)

func TestMat_AndByCoord(t *testing.T) {
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
			mat1: testLogicMat.Clone(),
			x:    3,
			y:    2,
			mat2: &Mat{
				width:      8,
				height:     4,
				widthBytes: 1,
				data: []uint8{
					0b01101110,
					0b11001110,
					0b01001110,
					0b10010000,
				},
			},
			wantMRes: &Mat{
				width:      testLogicMat.width,
				height:     testLogicMat.height,
				widthBytes: testLogicMat.widthBytes,
				data: []uint8{
					0, 0, 0,
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					0b01101110, 0b00001110, 0b00101000,
					0b11001100, 0b00001010, 0b11000000,
					0b01001000, 0b00001110, 0b10010000,
					0b10000000, 0b01001110, 0b11001100,
					0, 0, 0,
				},
			},
		},
		{
			name: "x:-3 y:-2",
			mat1: testLogicMat.Clone(),
			x:    -3,
			y:    -2,
			mat2: &Mat{
				width:      8,
				height:     4,
				widthBytes: 1,
				data: []uint8{
					0b01101110,
					0b11001110,
					0b01001110,
					0b10010000,
				},
			},
			wantMRes: &Mat{
				width:      testLogicMat.width,
				height:     testLogicMat.height,
				widthBytes: testLogicMat.widthBytes,
				data: []uint8{
					0, 0, 0,
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					0b00000110, 0b00001110, 0b00101000,
					0b11001110, 0b00101010, 0b11000000,
					0b01001110, 0b00001110, 0b10010000,
					0b10010000, 0b01001110, 0b11001100,
					0, 0, 0,
				},
			},
		},
		/*

		 01110111
		 10000111
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMRes := tt.mat1.AndByCoord(tt.mat2, tt.x, tt.y); !reflect.DeepEqual(gotMRes, tt.wantMRes) {
				t.Errorf("Mat.AndByCoord() = %v, want %v", gotMRes, tt.wantMRes)
			}
		})
	}
}

func TestMat_OrByCoord(t *testing.T) {
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
			mat1: testLogicMat.Clone(),
			x:    3,
			y:    2,
			mat2: &Mat{
				width:      8,
				height:     4,
				widthBytes: 1,
				data: []uint8{
					0b01101110,
					0b11001110,
					0b01001110,
					0b10010000,
				},
			},
			wantRes: &Mat{
				width:      testLogicMat.width,
				height:     testLogicMat.height,
				widthBytes: testLogicMat.widthBytes,
				data: []uint8{
					0, 0, 0,
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					0b01101110, 0b00001110, 0b00101000,
					0b11001111, 0b11101010, 0b11000000,
					0b01011111, 0b11001110, 0b10010000,
					0b10011001, 0b11001110, 0b11001100,
					0b00010010, 0, 0,
				},
			},
		},
		{
			name: "x:-3 y:-2",
			mat1: testLogicMat.Clone(),
			x:    -3,
			y:    -2,
			mat2: &Mat{
				width:      8,
				height:     4,
				widthBytes: 1,
				data: []uint8{
					0b01101110,
					0b11001110,
					0b01001110,
					0b10010000,
				},
			},
			wantRes: &Mat{
				width:      testLogicMat.width,
				height:     testLogicMat.height,
				widthBytes: testLogicMat.widthBytes,
				data: []uint8{
					0b01110000, 0, 0,
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					0b11101110, 0b00001110, 0b00101000,
					0b11001110, 0b00101010, 0b11000000,
					0b01001110, 0b00001110, 0b10010000,
					0b10010000, 0b01001110, 0b11001100,
					0, 0, 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMRes := tt.mat1.OrByCoord(tt.mat2, tt.x, tt.y); !reflect.DeepEqual(gotMRes, tt.wantRes) {
				t.Errorf("Mat.OrByCoord() = %v, want %v", gotMRes, tt.wantRes)
			}
		})
	}
}

func TestMat_XorByCoord(t *testing.T) {
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
			mat1: testLogicMat.Clone(),
			x:    3,
			y:    2,
			mat2: &Mat{
				width:      8,
				height:     4,
				widthBytes: 1,
				data: []uint8{
					0b01101110,
					0b11001110,
					0b01001110,
					0b10010000,
				},
			},
			wantRes: &Mat{
				width:      testLogicMat.width,
				height:     testLogicMat.height,
				widthBytes: testLogicMat.widthBytes,
				data: []uint8{
					0, 0, 0,
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					0b01101110, 0b00001110, 0b00101000,
					0b11000011, 0b11101010, 0b11000000,
					0b01010111, 0b11001110, 0b10010000,
					0b10011001, 0b10001110, 0b11001100,
					0b00010010, 0, 0,
				},
			},
		},
		{
			name: "x:-3 y:-2",
			mat1: testLogicMat.Clone(),
			x:    -3,
			y:    -2,
			mat2: &Mat{
				width:      8,
				height:     4,
				widthBytes: 1,
				data: []uint8{
					0b01101110,
					0b11001110,
					0b01001110,
					0b10010000,
				},
			},
			wantRes: &Mat{
				width:      testLogicMat.width,
				height:     testLogicMat.height,
				widthBytes: testLogicMat.widthBytes,
				data: []uint8{
					0b01110000, 0, 0,
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					0b11101110, 0b00001110, 0b00101000,
					0b11001110, 0b00101010, 0b11000000,
					0b01001110, 0b00001110, 0b10010000,
					0b10010000, 0b01001110, 0b11001100,
					0, 0, 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMRes := tt.mat1.XorByCoord(tt.mat2, tt.x, tt.y); !reflect.DeepEqual(gotMRes, tt.wantRes) {
				t.Errorf("Mat.XorByCoord() = %v, want %v", gotMRes, tt.wantRes)
			}
		})
	}
}
