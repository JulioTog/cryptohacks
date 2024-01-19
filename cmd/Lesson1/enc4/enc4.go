package main

import (
	"JulioTog/CryptohacksCode/utils"
	"fmt"
	"math/big"
)

var defaultMessage, err = new(big.Int).SetString("11515195063862318899931685488813747395775516287289682636499965282714637259206269", 10)

func main() {
	fmt.Printf(utils.LongNumDecoder(*defaultMessage))
}
