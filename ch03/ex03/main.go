package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type Polygon struct {
	ax, ay, bx, by, cx, cy, dx, dy float64
	zAvg                           float64
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	var polygons []Polygon
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			polygons = append(
				polygons,
				Polygon{
					ax:   ax,
					ay:   ay,
					bx:   bx,
					by:   by,
					cx:   cx,
					cy:   cy,
					dx:   dx,
					dy:   dy,
					zAvg: (az + bz + cz + dz) / 4.0,
				},
			)
		}
	}
	var zs []float64
	for _, polygon := range polygons {
		zs = append(zs, polygon.zAvg)
	}
	zMax, ok := max(zs)
	if !ok {
		fmt.Printf("err")
		return
	}
	zMin, ok := min(zs)
	if !ok {
		fmt.Print("err")
		return
	}
	for _, polygon := range polygons {
		var color string
		if polygon.zAvg >= 0 {
			color = fmt.Sprintf("#%02X0000", uint8((0xff * polygon.zAvg / zMax)))
		} else {
			color = fmt.Sprintf("#0000%02X", uint8(-(0xff * polygon.zAvg / zMin)))
		}
		fmt.Printf("<polygon stroke='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			color, polygon.ax, polygon.ay, polygon.bx, polygon.by, polygon.cx, polygon.cy, polygon.dx, polygon.dy)
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r

}
func max(floats []float64) (float64, bool) {
	if len(floats) == 0 {
		return 0, false
	}

	var result float64
	for _, x := range floats {
		if x > result {
			result = x
		}
	}
	return result, true
}

func min(floats []float64) (float64, bool) {
	if len(floats) == 0 {
		return 0, false
	}

	var result float64
	for _, x := range floats {
		if x < result {
			result = x
		}
	}
	return result, true
}
