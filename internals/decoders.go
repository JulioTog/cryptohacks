package internals

import (
	"encoding/hex"
	"fmt"
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
