package utils

import "math"

func CuadraticResidue(a, mod int) (result []int) {
	result = make([]int, 0)
	for x := 1; x < mod; x++ {
		if res := int(math.Pow(float64(x), 2)) % mod; res == a {
			result = append(result, x)
		}
	}
	return
}

func CuadraticResidueBatch(list []int, mod int) []int {
	res := make([]int, 0)
	for _, v := range list {
		val := CuadraticResidue(v, mod)
		for _, v := range val {
			res = append(res, v)
			res = append(res, -v)
		}
	}
	return res
}
