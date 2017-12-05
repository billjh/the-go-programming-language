// $ go build main.go temp.go
// $ ./main -t temperature 0 100
//
// 0°F = -17.77777777777778°C, 0°C = 32°F
// 100°F = 37.77777777777778°C, 100°C = 212°F

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var convType = flag.String("t", "", "type of conversion (eg. -t temperature)")

func main() {
	flag.Parse()
	switch *convType {
	case "temperature":
		for _, arg := range flag.Args() {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "tempconv: %v\n", err)
				os.Exit(1)
			}
			fmt.Print(tempconv(t))
		}
	// case "length"
	// case "weight"
	default:
		flag.Usage()
	}
}

func tempconv(t float64) string {
	f := Fahrenheit(t)
	c := Celsius(t)
	return fmt.Sprintf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
}
