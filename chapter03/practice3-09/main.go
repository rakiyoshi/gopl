package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	w.Header().Set("Content-Type", "image/png")

	const (
		defaultXMax, defaultYMax = +2, +2
		defaultXMin, defaultYMin = -2, -2
		width, height            = 1024, 1024
		defaultScale             = 1.0
	)

	var scale float64
	var err error
	var xmin, ymin, xmax, ymax float64

	if q := queries.Get("scale"); q != "" {
		scale, err = strconv.ParseFloat(q, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		scale = defaultScale
	}

	if q := queries.Get("x"); q != "" {
		xmax, err = strconv.ParseFloat(q, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		xmax /= scale
		xmin = -xmax
	} else {
		xmax = defaultXMax / scale
		xmin = -xmax
	}

	if q := queries.Get("y"); q != "" {
		ymax, err = strconv.ParseFloat(q, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		ymax /= scale
		ymin = -ymax
	} else {
		ymax = defaultYMax / scale
		ymin = -ymax
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	// nolint:errcheck
	png.Encode(w, img)
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
