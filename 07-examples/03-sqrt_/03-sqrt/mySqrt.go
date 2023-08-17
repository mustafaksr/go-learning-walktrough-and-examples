package main

import "math"

func mySqrt(x int) int {
	z := 1.0

	for i := 1; i < 100; i++ {
		z -= (z*z - float64(x)) / (2 * z) // This general approach is called Newton's method

	}
	z = math.Floor(z)
	return int(z)
}
