package utils

import (
	"fmt"
	"time"
)

func EuclideanClassic(a, b int) int {
	start := time.Now()
	defer fmt.Printf("Classic took %s \n", time.Since(start))
	for a != b {

		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return b
}

func EuclideanDivision(a, b int) int {
	start := time.Now()
	defer fmt.Printf("Division took %s \n", time.Since(start))

	divisor := max(a, b)
	dividend := min(a, b)
	for dividend != 0 {
		rest := divisor % dividend
		divisor = dividend
		dividend = rest
		if dividend == 0 {
			break
		}
	}
	return divisor
}

func EuclideanExtended(a, b int) (int, int, int) {

	if a == 0 {
		return b, 0, 1
	} else {
		g, x, y := EuclideanExtended(b%a, a)
		res := int(y - (b/a)*x)
		return g, res, x
	}
}
