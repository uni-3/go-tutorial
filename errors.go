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
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	z_n_1 := 0.0
	delta := 1.0e-6
	count := 0
	for math.Abs(z-z_n_1) >= delta {
		z_n_1 = z
		z = z - (z*z-x)/(2*z)
		count++
	}
	return z, nil
}

func main() {
	sp, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sqrt(2)", sp)

	sn, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sqrt(-2)", sn)
}
