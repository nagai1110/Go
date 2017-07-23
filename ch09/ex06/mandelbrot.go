// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"time"
)

func main() {
	createMandelbrot(1)
	createMandelbrot(2)
	createMandelbrot(4)
	createMandelbrot(8)
	createMandelbrot(16)
}

func createMandelbrot(maxProcs int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	runtime.GOMAXPROCS(maxProcs)
	start := time.Now()
	done := make(chan struct{})

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		go func(py int) {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}

			done <- struct{}{}
		}(py)
	}

	for py := 0; py < height; py++ {
		<-done
	}

	// png.Encode(os.Stdout, img)
	end := time.Now()
	fmt.Printf("GOMAXPROCS = %d\t Time = %v\n", maxProcs, end.Sub(start))
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
