package main

import (
	"JulioTog/CryptohacksCode/utils"
	"fmt"
	"math/big"
)

var module = big.NewInt(29)
var list = []big.Int{*big.NewInt(14), *big.NewInt(6), *big.NewInt(11)}

func main() {
	fmt.Println(utils.CuadraticResidueBatch(list, *module))
}
