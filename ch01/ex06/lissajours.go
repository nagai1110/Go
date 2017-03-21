// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var out io.Writer = os.Stdout

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	bgColorIndex    = 0
	line1ColorIndex = 1
	line2ColorIndex = 2
	line3ColorIndex = 3
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	createLissajours(palette)
}

func createLissajours(palette []color.Color) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			sinX := math.Sin(t)
			sinY := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(sinX*size+0.5), size+int(sinY*size+0.5),
				line1ColorIndex)

			cosX := math.Cos(t)
			cosY := math.Cos(t*freq + phase)
			img.SetColorIndex(size+int(cosX*size+0.5), size+int(cosY*size+0.5),
				line2ColorIndex)

			tanX := math.Tan(t)
			tanY := math.Tan(t*freq + phase)
			img.SetColorIndex(size+int(tanX*size+0.5), size+int(tanY*size+0.5),
				line3ColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
