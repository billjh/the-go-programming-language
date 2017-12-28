// go run main.go > m.png

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

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// https://stackoverflow.com/a/16505538
var colors = [16]color.RGBA{
	color.RGBA{66, 30, 15, 255},
	color.RGBA{25, 7, 26, 255},
	color.RGBA{9, 1, 47, 255},
	color.RGBA{4, 4, 73, 255},
	color.RGBA{0, 7, 100, 255},
	color.RGBA{12, 44, 138, 255},
	color.RGBA{24, 82, 177, 255},
	color.RGBA{57, 125, 209, 255},
	color.RGBA{134, 181, 229, 255},
	color.RGBA{211, 236, 248, 255},
	color.RGBA{241, 233, 191, 255},
	color.RGBA{248, 201, 95, 255},
	color.RGBA{255, 170, 0, 255},
	color.RGBA{204, 128, 0, 255},
	color.RGBA{153, 87, 0, 255},
	color.RGBA{106, 52, 3, 255},
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colors[(n+1)%16]
		}
	}
	return color.Black
}
