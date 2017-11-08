package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	z_n_1 := 0.0
	delta := 1.0e-6
	count := 0
	for math.Abs(z-z_n_1) >= delta {
		z_n_1 = z
		z = z - (z*z-x)/(2*z)
		count++
	}
	return z, count
}

func main() {
	for x := 1.0; x < 30; x++ {
		z, count := Sqrt(x)
		fmt.Printf("Sqrt(%f): loop count: %d\n", x, count)
		fmt.Printf(" correct: %f\n calc: %f\n", math.Sqrt(x), z)
	}
}
