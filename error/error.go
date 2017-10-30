package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ",float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x<0 {
		return x, ErrNegativeSqrt(x)
	}
	z := float64(1)

	for i:=0; i<10; i++ {
		z1 := (z + x/z) / 2
		if (math.Abs(z1-z) < 1e-9) {
			return z1, nil
		} else {
			z = z1
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

