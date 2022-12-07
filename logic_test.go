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
		name     string
		m1       *Mat
		m2       *Mat
		logic    string
		wantMRes *Mat
	}{
		{
			name:  "NOT",
			m1:    m1_8x1,
			m2:    nil,
			logic: "not",
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
			wantMRes: &Mat{
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
				if gotMRes := tt.m1.And(tt.m2); !reflect.DeepEqual(gotMRes, tt.wantMRes) {
					t.Errorf("Mat.And() = %v, want %v", gotMRes, tt.wantMRes)
				}
			case "or":
				if gotMRes := tt.m1.Or(tt.m2); !reflect.DeepEqual(gotMRes, tt.wantMRes) {
					t.Errorf("Mat.Or() = %v, want %v", gotMRes, tt.wantMRes)
				}
			case "xor":
				if gotMRes := tt.m1.Xor(tt.m2); !reflect.DeepEqual(gotMRes, tt.wantMRes) {
					t.Errorf("Mat.Xor() = %v, want %v", gotMRes, tt.wantMRes)
				}

			case "not":
				if gotMRes := tt.m1.Not(); !reflect.DeepEqual(gotMRes, tt.wantMRes) {
					t.Errorf("Mat.Not() = %v, want %v", gotMRes, tt.wantMRes)
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
			name: "x:3 y:2 and",
			mat1: testLogicMat,
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
			name: "x:-3 y:-2 and",
			mat1: testLogicMat,
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
					0b11001100, 0b00001010, 0b11000000,
					0b01001000, 0b00001110, 0b10010000,
					0b10000000, 0b01001110, 0b11001100,
					0, 0, 0,
				},
			},
		},

		/*
			01101110 00001110 00101000
			11001110 00101010 11000000
			01001110 00001110 10010000
			10010000 01001110 11001100
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
