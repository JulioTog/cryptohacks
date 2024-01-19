package main

import (
	"JulioTog/CryptohacksCode/utils"
	"encoding/base64"
	"fmt"
)

var hexString = "72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf"

func main() {
	result, err := utils.HexDecoder(hexString)
	if err != nil {
		println(err)
		return
	}
	encodedString := base64.RawStdEncoding.EncodeToString([]byte(result))
	fmt.Println(encodedString)

}
