package utils

import "fmt"

func XorCalc(b1, b2 []byte) (result []byte) {
	if len(b1) != len(b2) {
		fmt.Println("Not same lenght")
		result = []byte(nil)
		return
	} else {
		for i, v := range b1 {
			key := v ^ b2[i]
			result = append(result, key)
		}
	}
	return
}
