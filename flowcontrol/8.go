package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z_prev := 0.0
	for z != z_prev {
		z_prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	input := float64(math.Pow(2.3, 32.0))
	fmt.Println(Sqrt(input))
	fmt.Println(math.Sqrt(input))
}
