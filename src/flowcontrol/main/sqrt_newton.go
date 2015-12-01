package main

import (
	"fmt"
	"math"
)

func mySqrt(x float64) float64 {

	epsilon := 1e-20
	var z0, z1 float64 = 1.0, 1.0

	for {
		z0 = z1
		z1 = (x / z0 + z0) / 2.0

		if math.Abs(z0 - z1) < epsilon {
			break
		}
	}

	return z1
}

func main() {
	fmt.Println(mySqrt(10))
}
