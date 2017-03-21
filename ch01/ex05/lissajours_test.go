// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"bytes"
	"image/color"
	"image/gif"
	"testing"
)

func TestLissajours(t *testing.T) {
	var bgPalette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}
	var wbPalette = []color.Color{color.White, color.Black}

	// カラーだけ検査
	var tests = []struct {
		palette []color.Color
		want    []color.Color
	}{
		{bgPalette, bgPalette},
		{wbPalette, wbPalette},
	}

	for _, test := range tests {
		out = new(bytes.Buffer)
		createLissajours(test.palette)

		reader := bytes.NewReader(out.(*bytes.Buffer).Bytes())
		gif, err := gif.DecodeAll(reader)
		if err != nil {
			t.Errorf("gif.DecodeAll(%v): %v\n", reader, err)
			continue
		}

		// 最初のフレームのみ
		firstPalette := gif.Image[0]
		for i := 0; i < len(test.want); i++ {
			if firstPalette.Palette[i] != firstPalette.Palette.Convert(test.want[i]) {
				t.Errorf("Palette[%d] = %v, want %v", i, firstPalette.Palette[i], firstPalette.Palette.Convert(test.want[i]))
			}
		}
	}
}
