package main

import (
	"JulioTog/CryptohacksCode/utils"
	"fmt"
)

var module = 29
var list = []int{14, 6, 11}

func main() {
	fmt.Println(utils.CuadraticResidueBatch(list, module))
}
