package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// z -= (z * z - x) / (2 * z)
	const E = 0.000000000001
	var z = 1.0
	for math.Abs(z*z-x) >= E {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
