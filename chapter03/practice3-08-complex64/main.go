package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024 * 4, 1024 * 4
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex64(complex(x, y))
			img.Set(px, py, simpleFractal(z))
		}
	}
	// nolint:errcheck
	png.Encode(os.Stdout, img)
}

func simpleFractal(x complex64) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + x
		if cmplx.Abs(complex128(f(x))) < 0.1 {
			r, g, b := color.YCbCrToRGB(
				0x00,
				uint8(0xff)-contrast*n,
				contrast*n,
			)
			return color.RGBA{
				R: r,
				G: g,
				B: b,
				A: 0xff,
			}
		}
		x = x - f(x)/fDiff(x)
	}
	r, g, b, a := color.Black.RGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

func f(x complex64) complex64 {
	return complex64(cmplx.Pow(complex128(x), 4) - 1)
}

func fDiff(x complex64) complex64 {
	return complex64(4 * cmplx.Pow(complex128(x), 3))
}
