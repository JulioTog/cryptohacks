package internals

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

func HexDecoder(hexString string) ([]byte, error) {
	ascii, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Printf("Error decoding hex")
		return nil, err
	}
	return ascii, err
}

func AsciiDecoder(stringToDecode []byte) string {
	var decodedString string
	for _, character := range stringToDecode {
		decodedString += string(character)
	}
	return decodedString
}

func LongNumDecoder(stringToDecode big.Int) string {
	b := stringToDecode.Bytes()
	return AsciiDecoder(b)
}

func BruteForceDecode(stringToDecode string) (string, string) {
	b, _ := hex.DecodeString(stringToDecode)
	for i := 0; i < 255; i++ {
		result := []byte(nil)
		for _, v := range b {

			result = append(result, v^byte(i))
		}
		if DecodedString := AsciiDecoder(result); strings.Contains(DecodedString, "crypto") {
			return DecodedString, fmt.Sprint(i)
		}
	}
	return "Key not found", ""
}
