package bmat

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestFromImage(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		backColor color.Color
		wantM     *Mat
	}{
		{
			name:      "8x2",
			filename:  "testdata/8x2.png",
			backColor: color.RGBA{A: 255},
			wantM: &Mat{
				width:      8,
				widthBytes: 1,
				height:     2,
				data:       []uint8{0b01010101, 0b10101010},
			},
		},
		{
			name:      "inverse 8x2",
			filename:  "testdata/inverse_8x2.png",
			backColor: color.RGBA{A: 255},
			wantM: &Mat{
				width:      8,
				widthBytes: 1,
				height:     2,
				data:       []uint8{0b10101010, 0b01010101},
			},
		},
		{
			name:      "10x2",
			filename:  "testdata/10x2.png",
			backColor: color.RGBA{A: 255},
			wantM: &Mat{
				width:      10,
				widthBytes: 2,
				height:     2,
				data:       []uint8{0b01010101, 0b01000000, 0b10101010, 0b10000000},
			},
		},
		{
			name:      "inverse 10x2",
			filename:  "testdata/inverse_10x2.png",
			backColor: color.RGBA{A: 255},
			wantM: &Mat{
				width:      10,
				widthBytes: 2,
				height:     2,
				data:       []uint8{0b10101010, 0b10000000, 0b01010101, 0b01000000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			im, err := getImageFromFilePath(tt.filename)
			require.Nil(t, err)

			if gotM := FromImage(im, tt.backColor); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("FromImage() = %v, want %v", gotM, tt.wantM)
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

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, err := png.Decode(f)
	return image, err
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMRes := tt.args.m.Area(tt.args.x, tt.args.y, tt.args.width, tt.args.height); !reflect.DeepEqual(gotMRes, tt.wantMRes) {
				t.Errorf("Mat.Area() = %v, want %v", gotMRes, tt.wantMRes)
			}
		})
	}
}
