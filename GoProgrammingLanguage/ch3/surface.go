// Surface computes an SVG rendering of a 3-d surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixes per z unit
	angle         = math.Pi / 8         // angle of x, y axes (=30 deg) i dont know how to make that lol
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin 30 cos 30

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	var baddieCount int
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, z := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)

			if math.IsInf(ax, 0) || math.IsInf(ay, 0) ||
				math.IsInf(bx, 0) || math.IsInf(by, 0) ||
				math.IsInf(cx, 0) || math.IsInf(cy, 0) ||
				math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				baddieCount++
			}
			fmt.Printf("<polygon style='fill: %s' points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				c(z),
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
	// fmt.Printf("Official Baddie Count %d", baddieCount)
}

func clamp(y float64) int {
	return int(math.Min(256.0, math.Max(y, 0.0)))
}

func c(y float64) string {
	intensity := clamp(y) / 2

	return fmt.Sprintf("#%x00%x",
		128-intensity, 128+intensity)
}

func corner(i, j int) (float64, float64, float64) {
	// find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface hegith
	z := f(x, y)

	// Project (x,y,z) isometrically onto a 2-d SVG canvas (sx,sy) lol
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z * zscale
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from 0 , 0
	return math.Sin(r) / r
}
