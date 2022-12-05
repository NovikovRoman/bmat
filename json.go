package bmat

import "encoding/json"

type matJson struct {
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Data   []uint8 `json:"data"`
}

func (m *Mat) ToJson() (b []byte) {
	b, _ = json.Marshal(matJson{
		Width:  m.width,
		Height: m.height,
		Data:   m.data,
	})
	return
}

func FromJson(b []byte) (m *Mat, err error) {
	var j matJson
	if err = json.Unmarshal(b, &j); err != nil {
		return
	}

	m = New(j.Width, j.Height)
	m.data = j.Data
	return
}
