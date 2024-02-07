package utils

import (
	"fmt"
	"math/big"
)

func CuadraticResidue(a, mod *big.Int) (result []big.Int) {
	result = make([]big.Int, 0)
	for x := big.NewInt(1); x.Cmp(mod) == -1; x.Add(big.NewInt(1), x) {
		pow := x
		if pow.Exp(x, big.NewInt(2), mod); pow.Cmp(a) == 0 {
			result = append(result, *x)
		}
	}
	return
}

func CuadraticResidueBatch(list []big.Int, mod big.Int) []big.Int {
	res := make([]big.Int, 0)
	for _, v := range list {
		val := CuadraticResidue(&v, &mod)
		for _, v := range val {
			res = append(res, v)
			//res = append(res, -v)
			fmt.Println("test")
		}
	}
	return res
}

func LegendreSymbol(val, module big.Int) bool {

	var module1, division big.Int
	division.Sub(&module, big.NewInt(1))
	division.Div(&division, big.NewInt(2))

	module1.Exp(&val, &division, &module)

	return module1.Cmp(big.NewInt(1)) == 0
}
