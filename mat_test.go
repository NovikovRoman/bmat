package bmat

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want *Mat
	}{
		{
			name: "8x2",
			args: args{
				width:  8,
				height: 2,
			},
			want: &Mat{
				width:      8,
				widthBytes: 1,
				height:     2,
				data:       make([]uint8, 2),
			},
		},
		{
			name: "9x1",
			args: args{
				width:  9,
				height: 1,
			},
			want: &Mat{
				width:      9,
				height:     1,
				widthBytes: 2,
				data:       make([]uint8, 2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMat_CountBits(t *testing.T) {
	tests := []struct {
		name  string
		mat   *Mat
		wantN int
	}{
		{
			name: "8x2",
			mat: &Mat{
				width:      8,
				widthBytes: 1,
				height:     2,
				data:       []uint8{0b01010101, 0b10101010},
			},
			wantN: 8,
		},
		{
			name: "10x2",
			mat: &Mat{
				width:      10,
				widthBytes: 2,
				height:     2,
				data:       []uint8{0b01010101, 0b01000000, 0b10101010, 0b10000000},
			},
			wantN: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if gotN := tt.mat.CountBits(); gotN != tt.wantN {
				t.Errorf("Mat.CountBits() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestMat_Area(t *testing.T) {
	m := &Mat{
		width:      16,
		height:     2,
		widthBytes: 2,
		data: []uint8{
			0b10101100, 0b11001010,
			0b11100011, 0b10001110,
		},
	}
	type args struct {
		m      *Mat
		x      int
		y      int
		width  int
		height int
	}
	tests := []struct {
		name     string
		args     args
		wantMRes *Mat
	}{
		{
			name: "x:2 y:0 w:8 h:1",
			args: args{
				m:      m,
				x:      2,
				y:      0,
				width:  8,
				height: 1,
			},
			wantMRes: &Mat{
				width:      8,
				widthBytes: 1,
				height:     1,
				data: []uint8{
					0b10110011,
				},
			},
		},
		{
			name: "x:2 y:0 w:12 h:2",
			args: args{
				m:      m,
				x:      2,
				y:      0,
				width:  12,
				height: 2,
			},
			wantMRes: &Mat{
				width:      12,
				widthBytes: 2,
				height:     2,
				data: []uint8{
					0b10110011, 0b00101000,
					0b10001110, 0b00111000,
				},
			},
		},
		{
			name: "x:6 y:1 w:24 h:2",
			args: args{
				m:      m,
				x:      6,
				y:      1,
				width:  24,
				height: 2,
			},
			wantMRes: &Mat{
				width:      24,
				widthBytes: 3,
				height:     2,
				data: []uint8{
					0b11100011, 0b10000000, 0,
					0, 0, 0,
				},
			},
		},
		{
			name: "x:6 y:-2 w:24 h:4",
			args: args{
				m:      m,
				x:      6,
				y:      -2,
				width:  24,
				height: 4,
			},
			wantMRes: &Mat{
				width:      24,
				widthBytes: 3,
				height:     4,
				data: []uint8{
					0, 0, 0,
					0, 0, 0,
					0b00110010, 0b10000000, 0,
					0b11100011, 0b10000000, 0,
				},
			},
		},
		{
			name: "x:-2 y:0 w:24 h:2",
			args: args{
				m:      m,
				x:      -2,
				y:      0,
				width:  24,
				height: 2,
			},
			wantMRes: &Mat{
				width:      24,
				widthBytes: 3,
				height:     2,
				data: []uint8{
					0b00101011, 0b00110010, 0b10000000,
					0b00111000, 0b11100011, 0b10000000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMRes := tt.args.m.Area(tt.args.x, tt.args.y, tt.args.width, tt.args.height); !reflect.DeepEqual(gotMRes, tt.wantMRes) {
				t.Errorf("Mat.Area() = %v, want %v", gotMRes, tt.wantMRes)
			}
		})
	}
}

var (
	testMarginMat = &Mat{
		width:      22,
		height:     8,
		widthBytes: 3,
		data: []uint8{
			0, 0, 0,
			0, 0, 0,
			// Последние два бита в каждой строке не учитываются, тк длина 22.
			0b00000000, 0b00001110, 0b00000000,
			0b00000000, 0b00101010, 0b11000000,
			0b00000000, 0b00001110, 0b10010000,
			0b00000000, 0b00001110, 0b00000000,
			0, 0, 0,
			0, 0, 0,
		},
	}
	testEmptyMat = &Mat{
		width:      22,
		height:     5,
		widthBytes: 3,
		data: []uint8{
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
			0, 0, 0,
		},
	}
)

func TestMat_TopMargin(t *testing.T) {
	tests := []struct {
		name       string
		mat        *Mat
		wantMargin int
	}{
		{
			name:       "top 2",
			mat:        testMarginMat,
			wantMargin: 2,
		},
		{
			name:       "top 4 empty mat",
			mat:        testEmptyMat,
			wantMargin: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMargin := tt.mat.TopMargin(); gotMargin != tt.wantMargin {
				t.Errorf("Mat.TopMargin() = %v, want %v", gotMargin, tt.wantMargin)
			}
		})
	}
}

func TestMat_BottomMargin(t *testing.T) {
	tests := []struct {
		name       string
		mat        *Mat
		wantMargin int
	}{
		{
			name:       "bottom 5",
			mat:        testMarginMat,
			wantMargin: 5,
		},
		{
			name:       "bottom 0 empty mat",
			mat:        testEmptyMat,
			wantMargin: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMargin := tt.mat.BottomMargin(); gotMargin != tt.wantMargin {
				t.Errorf("Mat.BottomMargin() = %v, want %v", gotMargin, tt.wantMargin)
			}
		})
	}
}

func TestMat_LeftMargin(t *testing.T) {
	tests := []struct {
		name       string
		mat        *Mat
		wantMargin int
	}{
		{
			name:       "left 10",
			mat:        testMarginMat,
			wantMargin: 10,
		},
		{
			name:       "left 21 empty mat",
			mat:        testEmptyMat,
			wantMargin: 21,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMargin := tt.mat.LeftMargin(); gotMargin != tt.wantMargin {
				t.Errorf("Mat.LeftMargin() = %v, want %v", gotMargin, tt.wantMargin)
			}
		})
	}
}

func TestMat_RightMargin(t *testing.T) {
	tests := []struct {
		name       string
		mat        *Mat
		wantMargin int
	}{
		{
			name:       "right 19",
			mat:        testMarginMat,
			wantMargin: 19,
		},
		{
			name: "right 12",
			mat: &Mat{
				width:      22,
				height:     8,
				widthBytes: 3,
				data: []uint8{
					0, 0, 0,
					0, 0, 0,
					// Последние два бита в каждой строке не учитываются, тк длина 22.
					0b00000000, 0b00001000, 0b00000000,
					0b00000000, 0b01100000, 0b00000000,
					0b00000000, 0b01001000, 0b00000000,
					0b00000000, 0b00100000, 0b00000000,
					0, 0, 0,
					0, 0, 0,
				},
			},
			wantMargin: 12,
		},
		{
			name:       "right 0 empty mat",
			mat:        testEmptyMat,
			wantMargin: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMargin := tt.mat.RightMargin(); gotMargin != tt.wantMargin {
				t.Errorf("Mat.RightMargin() = %v, want %v", gotMargin, tt.wantMargin)
			}
		})
	}
}

func TestMat_Clone(t *testing.T) {
	tests := []struct {
		name    string
		mat     *Mat
		wantMat *Mat
	}{
		{
			name:    "testMarginMat",
			mat:     testMarginMat,
			wantMat: testMarginMat,
		},
		{
			name:    "testEmptyMat",
			mat:     testEmptyMat,
			wantMat: testEmptyMat,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMat := tt.mat.Clone()

			if &gotMat == &tt.mat {
				t.Errorf("gotMat pointer equal tt.mat")
			}

			if !reflect.DeepEqual(gotMat, tt.wantMat) {
				t.Errorf("Mat.Clone() = %v, want %v", gotMat, tt.wantMat)
			}
		})
	}
}
