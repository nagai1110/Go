// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		xmin := parseFloat64(query.Get("xmin"), -2)
		ymin := parseFloat64(query.Get("ymin"), -2)
		xmax := parseFloat64(query.Get("xmax"), 2)
		ymax := parseFloat64(query.Get("ymax"), 2)
		width := parseInt(query.Get("width"), 1024)
		height := parseInt(query.Get("height"), 1024)

		writeMandelbrotImage(w, xmin, ymin, xmax, ymax, width, height)
	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parseFloat64(s string, defaultValue float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}

	return f
}

func parseInt(s string, defaultValue int) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return defaultValue
	}

	return i
}

func writeMandelbrotImage(writer io.Writer, xmin float64, ymin float64, xmax float64, ymax float64, width int, height int) {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(writer, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-
