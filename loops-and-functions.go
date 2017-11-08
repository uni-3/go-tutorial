package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	for x := 1.0; x < 10; x++ {
		fmt.Printf("Sqrt(%f)\n", x)
		fmt.Printf(" correct: %f\n calc: %f\n", math.Sqrt(x), Sqrt(x))
	}
}
