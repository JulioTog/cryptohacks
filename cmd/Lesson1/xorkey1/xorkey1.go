package main

import (
	"JulioTog/CryptohacksCode/utils"
	"encoding/hex"
	"fmt"
	"math"
)

var hexMessage = "0e0b213f26041e480b26217f27342e175d0e070a3c5b103e2526217f27342e175d0e077e263451150104"
var alwaysInit = []byte("crypto{")

func main() {
	b1, _ := hex.DecodeString(hexMessage)

	normalizedByte := make([]byte, len(b1))
	copy(normalizedByte, alwaysInit)
	result := utils.XorCalc(b1, normalizedByte)
	fmt.Println(utils.AsciiDecoder(result[:7]))

	key := result[:7]
	key = append(key, 'y')
	key2 := make([]byte, len(b1))
	for i, _ := range b1 {
		key2[i] = key[int(math.Mod(float64(i), float64(len(key))))]
	}

	message := utils.XorCalc(key2, b1)
	fmt.Println(utils.AsciiDecoder(message))
}
