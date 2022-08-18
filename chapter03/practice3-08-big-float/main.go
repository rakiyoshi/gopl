package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

type Complex struct {
	x *big.Float
	y *big.Float
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024 * 4, 1024 * 4
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := new(big.Float).SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := new(big.Float).SetFloat64(float64(px)/width*(xmax-xmin) + xmin)
			z := Complex{
				x: x,
				y: y,
			}
			c := simpleFractal(z)
			img.Set(px, py, c)
		}
	}
	// nolint:errcheck
	png.Encode(os.Stdout, img)
}

func simpleFractal(x Complex) color.RGBA {
	const iterations = 200
	const contrast = 15

	v := Complex{
		x: new(big.Float).SetFloat64(0.0),
		y: new(big.Float).SetFloat64(0.0),
	}
	for n := uint8(0); n < iterations; n++ {
		fx := f(x)
		if fx.x.IsInf() || fx.y.IsInf() {
			return color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0xff,
			}
		}
		v = plus(multiply(v, v), x)
		if abs(f(x)).Cmp(new(big.Float).SetFloat64(1)) == -1 {
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
		x = minus(x, divide(f(x), fDiff(x)))
	}
	r, g, b, a := color.Black.RGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

func f(z Complex) Complex {
	// x**4 - 1
	return minus(
		multiply(multiply(multiply(multiply(z, z), z), z), z),
		Complex{
			x: new(big.Float).SetFloat64(1),
			y: new(big.Float).SetFloat64(0),
		},
	)
}

func fDiff(z Complex) Complex {
	// 4 * x**3
	result := multiply(
		Complex{
			x: new(big.Float).SetFloat64(4.0),
			y: new(big.Float).SetFloat64(0),
		},
		multiply(multiply(multiply(z, z), z), z),
	)
	return result
}

func plus(a, b Complex) Complex {
	return Complex{
		x: new(big.Float).Add(a.x, b.x),
		y: new(big.Float).Add(a.y, b.y),
	}
}

func minus(a, b Complex) Complex {
	return Complex{
		x: new(big.Float).Sub(a.x, b.x),
		y: new(big.Float).Sub(a.y, b.y),
	}
}

func multiply(a, b Complex) Complex {
	return Complex{
		x: new(big.Float).Sub(
			new(big.Float).Mul(a.x, a.x),
			new(big.Float).Mul(b.y, b.y),
		),
		y: new(big.Float).Add(
			new(big.Float).Mul(a.x, b.y),
			new(big.Float).Mul(a.y, b.x),
		),
	}
}

func divide(a, b Complex) Complex {
	x, _ := b.x.Float64()
	y, _ := b.y.Float64()
	fmt.Fprintf(os.Stderr, "%X, %X\n", x, y)
	return Complex{
		x: new(big.Float).Quo(
			new(big.Float).Add(
				new(big.Float).Mul(a.x, b.x),
				new(big.Float).Mul(a.y, b.y),
			),
			new(big.Float).Add(
				new(big.Float).Mul(b.x, b.x),
				new(big.Float).Mul(b.y, b.y),
			),
		),
		y: new(big.Float).Quo(
			new(big.Float).Sub(
				new(big.Float).Mul(b.x, a.y),
				new(big.Float).Mul(a.x, b.y),
			),
			new(big.Float).Add(
				new(big.Float).Mul(b.x, b.x),
				new(big.Float).Mul(b.y, b.y),
			),
		),
	}
}

func abs(z Complex) *big.Float {
	return new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(z.x, z.x),
			new(big.Float).Mul(z.y, z.y),
		),
	)
}
