package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//import "encoding/binary"

var word = "label"
var operand byte = 13
func main(){
	b := []byte(word)
	buf := new(bytes.Buffer)
	err := binary.Write(buf,binary.LittleEndian,b)

	if err != nil {
		fmt.Println("binary failed")
	}

	for _,v := range buf.Bytes(){
		res := v ^ operand
		fmt.Printf("%b \n", res)
	}
}
