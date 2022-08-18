package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
)

type Complex struct {
	re *big.Float
	im *big.Float
}

// TODO: fixme
// these methods don't work well
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := new(big.Float).SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			x := new(big.Float).SetFloat64(float64(px)/width*(xmax-xmin) + xmin)
			z := Complex{
				re: x,
				im: y,
			}
			c := simpleFractal(z)
			img.Set(px, py, c)
		}
	}
	// nolint:errcheck
	png.Encode(os.Stdout, img)
}

func simpleFractal(z Complex) color.RGBA {
	const iterations = 100
	const contrast = 15

	v := Complex{
		re: new(big.Float).SetFloat64(0.0),
		im: new(big.Float).SetFloat64(0.0),
	}
	for n := uint8(0); n < iterations; n++ {
		if z.re.Cmp(new(big.Float).SetFloat64(-2)) == -1 ||
			z.re.Cmp(new(big.Float).SetFloat64(2)) == 1 ||
			z.im.Cmp(new(big.Float).SetFloat64(-2)) == -1 ||
			z.im.Cmp(new(big.Float).SetFloat64(2)) == 1 {
			return color.RGBA{
				R: 0x00,
				G: 0x00,
				B: 0x00,
				A: 0xff,
			}
		}
		v = plus(multiply(v, v), z)
		if abs(f(z)).Cmp(new(big.Float).SetFloat64(1)) == -1 {
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
		z = minus(z, divide(f(z), fDiff(z)))
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
			re: new(big.Float).SetFloat64(1),
			im: new(big.Float).SetFloat64(0),
		},
	)
}

func fDiff(z Complex) Complex {
	// 4 * x**3
	result := multiply(
		Complex{
			re: new(big.Float).SetFloat64(4.0),
			im: new(big.Float).SetFloat64(0),
		},
		multiply(multiply(multiply(z, z), z), z),
	)
	return result
}

func plus(a, b Complex) Complex {
	return Complex{
		re: new(big.Float).Add(a.re, b.re),
		im: new(big.Float).Add(a.im, b.im),
	}
}

func minus(a, b Complex) Complex {
	return Complex{
		re: new(big.Float).Sub(a.re, b.re),
		im: new(big.Float).Sub(a.im, b.im),
	}
}

func multiply(a, b Complex) Complex {
	return Complex{
		re: new(big.Float).Sub(
			new(big.Float).Mul(a.re, a.re),
			new(big.Float).Mul(b.im, b.im),
		),
		im: new(big.Float).Add(
			new(big.Float).Mul(a.re, b.im),
			new(big.Float).Mul(a.im, b.re),
		),
	}
}

func divide(a, b Complex) Complex {
	return Complex{
		re: new(big.Float).Quo(
			new(big.Float).Add(
				new(big.Float).Mul(a.re, b.re),
				new(big.Float).Mul(a.im, b.im),
			),
			new(big.Float).Add(
				new(big.Float).Mul(b.re, b.re),
				new(big.Float).Mul(b.im, b.im),
			),
		),
		im: new(big.Float).Quo(
			new(big.Float).Sub(
				new(big.Float).Mul(b.re, a.im),
				new(big.Float).Mul(a.re, b.im),
			),
			new(big.Float).Add(
				new(big.Float).Mul(b.re, b.re),
				new(big.Float).Mul(b.im, b.im),
			),
		),
	}
}

func abs(z Complex) *big.Float {
	return new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(z.re, z.re),
			new(big.Float).Mul(z.im, z.im),
		),
	)
}
