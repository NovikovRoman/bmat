package bmat

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, err := png.Decode(f)
	return image, err
}

func TestToImage(t *testing.T) {
	im, err := getImageFromFilePath("testdata/10x2.png")
	require.Nil(t, err)

	m := FromImage(im, color.RGBA{A: 255})
	buf := new(bytes.Buffer)

	err = png.Encode(buf, m.ToImage())
	require.Nil(t, err)

	h := sha1.New()
	var n int
	n, err = h.Write(buf.Bytes())
	require.Nil(t, err)
	require.Equal(t, n, 74)
	require.Equal(t, fmt.Sprintf("%x", h.Sum(nil)), "28c4f997569697e3af47af2c3dc6c27d6d9bd76c")
}
