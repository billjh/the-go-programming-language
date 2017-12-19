// go run main.go
// http://localhost:8000/?color=green&width=900&height=600

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	linecolor     = "grey"              // default stroke color
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		strokeColor := linecolor
		svgWidth := width
		svgHeight := height
		if c, ok := r.Form["color"]; ok {
			strokeColor = c[0]
		}
		if w, ok := r.Form["width"]; ok {
			if ww, err := strconv.Atoi(w[0]); err == nil {
				svgWidth = ww
			}
		}
		if h, ok := r.Form["height"]; ok {
			if hh, err := strconv.Atoi(h[0]); err == nil {
				svgHeight = hh
			}
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		svgSurface(w, strokeColor, svgWidth, svgHeight)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svgSurface(out io.Writer, strokeColor string, width, height int) {
	fmt.Fprintf(out,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: %s; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", strokeColor, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			if anyInf(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}
			fmt.Fprintf(out,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j, width, height int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
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
