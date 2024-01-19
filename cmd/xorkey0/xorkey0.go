package main

import (
	"JulioTog/CryptohacksCode/internals"
	"fmt"
)

var hexString = "73626960647f6b206821204f21254f7d694f7624662065622127234f726927756d"

func main() {
	decode, key := internals.BruteForceDecode(hexString)
	fmt.Printf("el mensaje es: \"%s\" y la key es \"%s\"", decode, key)
}
