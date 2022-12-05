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
