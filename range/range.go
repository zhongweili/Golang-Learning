package main

import (
	"fmt"
	"math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Sqrt(x float64) float64 {
	z := float64(1)

	for i:=0; i<10; i++ {
		z1 := (z + x/z) / 2
		if (math.Abs(z1-z) < 1e-9) {
    		return z1
		} else {
			z = z1
		}
	}
	return z
}


func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	fmt.Println(Sqrt(2))
}
