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
		width, height          = 1024, 1024
	)

	var pointColor [][]color.RGBA
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		pointColor = append(pointColor, []color.RGBA{})
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			c := simpleFractal(z)
			pointColor[py] = append(
				pointColor[py],
				c,
			)
		}
	}
	pointColor = supersample(pointColor)
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py, pointColor[py][px])
		}
	}
	// nolint:errcheck
	png.Encode(os.Stdout, img)
}

func simpleFractal(x complex128) color.RGBA {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + x
		if cmplx.Abs(f(x)) < 0.1 {
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

func supersample(pointColor [][]color.RGBA) [][]color.RGBA {
	height := len(pointColor)
	width := len(pointColor[0])
	var result [][]color.RGBA
	for y := 0; y < height; y++ {
		result = append(result, []color.RGBA{})
		for x := 0; x < width; x++ {
			y0, x0 := y, x
			var y1, x1 int
			if y+1 < height {
				y1 = y + 1
			} else {
				y1 = y
			}
			if x+1 < width {
				x1 = x + 1
			} else {
				x1 = x
			}
			r0, g0, b0, a0 := pointColor[y0][x0].RGBA()
			r1, g1, b1, a1 := pointColor[y1][x0].RGBA()
			r2, g2, b2, a2 := pointColor[y0][x1].RGBA()
			r3, g3, b3, a3 := pointColor[y1][x1].RGBA()
			result[y] = append(
				result[y],
				color.RGBA{
					R: uint8((uint16(uint8(r0)) + uint16(uint8(r1)) + uint16(uint8(r2)) + uint16(uint8(r3))) / 4),
					G: uint8((uint16(uint8(g0)) + uint16(uint8(g1)) + uint16(uint8(g2)) + uint16(uint8(g3))) / 4),
					B: uint8((uint16(uint8(b0)) + uint16(uint8(b1)) + uint16(uint8(b2)) + uint16(uint8(b3))) / 4),
					A: uint8((uint16(uint8(a0)) + uint16(uint8(a1)) + uint16(uint8(a2)) + uint16(uint8(a3))) / 4),
				},
			)
		}
	}
	return result
}

func f(x complex128) complex128 {
	return cmplx.Pow(x, 4) - 1
}

func fDiff(x complex128) complex128 {
	return 4 * cmplx.Pow(x, 3)
}
