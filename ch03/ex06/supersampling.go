// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	img := mandelbrotImage(2048, 2048)
	img = superSampling(img, 1024, 1024)

	png.Encode(os.Stdout, img)
}

func mandelbrotImage(width int, height int) image.Image {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)

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

	return img
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - contrast*n
			b := contrast * n
			return color.RGBA{r, 0, b, 255}
		}
	}
	return color.Black
}

func superSampling(src image.Image, width int, height int) image.Image {
	scaleX := float64(width) / float64(src.Bounds().Size().X)
	scaleY := float64(height) / float64(src.Bounds().Size().Y)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			dst.Set(x, y, subPixelColor(src, x, y, scaleX, scaleY))
		}
	}

	return dst
}

func subPixelColor(src image.Image, x int, y int, scaleX float64, scaleY float64) color.Color {
	startX := float64(x) / scaleX
	startY := float64(y) / scaleY
	endX := float64(x+1) / scaleX
	endY := float64(y+1) / scaleY

	imageWidth := src.Bounds().Size().X
	imageHeight := src.Bounds().Size().Y

	var subPixelCount uint32
	var r, g, b, a uint32
	for x := int(startX); x < int(endX); x++ {
		for y := int(startY); y < int(endY); y++ {
			if x >= imageWidth || y >= imageHeight {
				continue
			}

			sr, sg, sb, sa := src.At(x, y).RGBA()
			// TODO:位置による重みなし
			r += sr
			g += sg
			b += sb
			a += sa

			subPixelCount++
		}
	}

	r /= subPixelCount
	g /= subPixelCount
	b /= subPixelCount
	a /= subPixelCount

	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
}

//!-
