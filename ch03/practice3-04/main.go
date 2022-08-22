package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	defaultWidth, defaultHeight = 600, 320
	defaultCells                = 100
	defaultXYRagne              = 30.0
)

type Points struct {
	x float64
	y float64
	z float64
}

var sin30, cos30 = math.Sin(math.Pi / 6), math.Cos(math.Pi / 6)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	var width, height, cells int
	var xyrange float64
	var err error
	queries := r.URL.Query()

	if q := queries.Get("width"); q != "" {
		width, err = strconv.Atoi(q)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		width = defaultWidth
	}

	if q := queries.Get("height"); q != "" {
		height, err = strconv.Atoi(q)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		height = defaultHeight
	}

	if q := queries.Get("cells"); q != "" {
		cells, err = strconv.Atoi(q)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		cells = defaultCells
	}

	if q := queries.Get("xyrange"); q != "" {
		xyrange, err = strconv.ParseFloat(q, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		xyrange = defaultXYRagne
	}

	fmt.Fprint(w, genSvg(width, height, xyrange, cells))
}

func genSvg(width, height int, xyrange float64, cells int) string {
	var points [][]Points
	for i := 0; i < cells; i++ {
		points = append(
			points,
			[]Points{},
		)
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/float64(cells) - 0.5)
			y := xyrange * (float64(j)/float64(cells) - 0.5)
			points[i] = append(
				points[i],
				Points{
					x: x,
					y: y,
					z: f(x, y),
				},
			)
		}
	}

	xyscale := float64(width) / 2.0 / xyrange
	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells-1; i++ {
		for j := 0; j < cells-1; j++ {
			ax, ay := corner(
				points[i+1][j].x,
				points[i+1][j].y,
				points[i+1][j].z,
				width,
				height,
				xyscale,
			)
			bx, by := corner(points[i][j].x,
				points[i][j].y,
				points[i][j].z,
				width,
				height,
				xyscale,
			)
			cx, cy := corner(points[i][j+1].x,
				points[i][j+1].y,
				points[i][j+1].z,
				width,
				height,
				xyscale,
			)
			dx, dy := corner(points[i+1][j+1].x,
				points[i+1][j+1].y,
				points[i+1][j+1].z,
				width,
				height,
				xyscale,
			)
			svg += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg += "</svg>"
	return svg
}

func corner(x, y, z float64, width, height int, xyscale float64) (float64, float64) {
	zscale := float64(height) * 0.4
	sx := float64(width)/2.0 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
