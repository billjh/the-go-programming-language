// go run main.go > svg.html

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
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	minz, maxz := minmaxz()
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if anyInf(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			_, _, z := xyz(i, j)
			fmt.Printf("<polygon style='stroke: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(minz, maxz, z), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func color(minz, maxz, z float64) string {
	step := (maxz - minz) / 256
	level := int((z-minz)/step) % 256
	return fmt.Sprintf("#%02X00%02X", level, 255-level)
}

func minmaxz() (minz, maxz float64) {
	minz, maxz = math.MaxFloat64, -math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			_, _, z := xyz(i, j)
			if z < minz {
				minz = z
			}
			if z > maxz {
				maxz = z
			}
		}
	}
	return
}

func xyz(i, j int) (x, y, z float64) {
	// Find point (x,y) at corner of cell (i,j).
	x = xyrange * (float64(i)/cells - 0.5)
	y = xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z = f(x, y)
	return
}

func corner(i, j int) (float64, float64) {
	x, y, z := xyz(i, j)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func anyInf(nums ...float64) bool {
	for _, n := range nums {
		if math.IsInf(n, 0) {
			return true
		}
	}
	return false
}
