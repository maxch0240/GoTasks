package main

import (
	"fmt"
	"math"
)

func main() {
	var R *float64
	var length float64 = 50

	R = &length
	S := 3.14 * math.Pow(*R/(2*3.14), 2)

	fmt.Println(round(S))
}

func round(num float64) float64 {
	return math.Round(num*100) / 100
}
