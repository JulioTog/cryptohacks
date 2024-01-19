package main

import (
	"JulioTog/CryptohacksCode/utils"
	"encoding/hex"
	"fmt"
	//"bytes"
)

var key1 = "a6c8b6733c9b22de7bc0253266a3867df55acde8635e19c73313"
var key21 = "37dcb292030faa90d07eec17e3b1c6d8daf94c35d4c9191a5e1e"
var key23 = "c1545756687e7573db23aa1c3452a098b71a7fbf0fddddde5fc1"
var fullkey = "04ee9855208a2cd59091d04767ae47963170d1660df7f56f5faf"

func main() {
	key1, err := hex.DecodeString(key1)
	b2, err := hex.DecodeString(key21)
	if err != nil {
		println("error decoding")
	}

	key2 := utils.XorCalc(key1, b2)
	fmt.Printf("Key 2 %b", key2)

	b3, _ := hex.DecodeString(key23)

	key3 := utils.XorCalc(key2, b3)
	fmt.Printf("key 3 %b", key3)

	b4, _ := hex.DecodeString(fullkey)

	flag := utils.XorCalc(b4, key1)
	flag = utils.XorCalc(flag, key2)
	flag = utils.XorCalc(flag, key3)

	println(utils.AsciiDecoder(flag))
}
