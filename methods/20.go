package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	z := 1.0
	z_prev := 0.0
	for math.Abs(z_prev-z) > 1e-12 {
		z_prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	x, err := Sqrt(2)
	if err == nil {
		fmt.Println(x)
	}

	x, err = Sqrt(-2)
	if err == nil {
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}
}
