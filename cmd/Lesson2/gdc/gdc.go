package main

import (
	"JulioTog/CryptohacksCode/utils"
	"fmt"
)

var a = 3
var b = 1

func main() {
	fmt.Println(utils.EuclideanClassic(b, a))
	fmt.Println(utils.EuclideanDivision(b, a))
}
