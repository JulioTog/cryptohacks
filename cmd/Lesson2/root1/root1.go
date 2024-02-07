package main

import (
	"JulioTog/CryptohacksCode/utils"
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, _ := os.Open("/Users/cascotito/Documents/cryptohacks/resources/output.txt")
	defer data.Close()
	//check(err)
	p := big.NewInt(0)
	p2 := big.NewInt(0)
	scanner := bufio.NewScanner(data)
	ints := make([]big.Int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "p = ") {
			prevP := strings.Replace(line, "p = ", "", 1)
			p.SetString(prevP, 10)
			p2.SetString(prevP, 10)
		} else if strings.Contains(line, "ints = ") {
			line = strings.Replace(line, "[", "", 1)
			line = strings.Replace(line, "]", "", 1)
			intsOnString := strings.Split(strings.Trim(strings.Replace(line, "ints = ", "", 1), "  "), ",")

			for _, v := range intsOnString {
				aConcat, _ := new(big.Int).SetString(v, 10)
				ints = append(ints, *aConcat)
			}
			//fmt.Println(ints)

		}
	}
	//fmt.Println(string(*ints))
	cuadratics := make([]big.Int, 0)
	cnt := 0
	for _, v := range ints {
		if utils.LegendreSymbol(v, *p) {
			cuadratics = append(cuadratics, v)
			cnt += 1
		}
	}
	fmt.Printf("cantidad de sqr encontrados: %d \n el valor es %s \n", cnt, cuadratics[0].String())

	//Sabemos que p = 3 mod 4 = 3

	modulo := *p2
	modulo.Add(&modulo, big.NewInt(1))
	modulo.Div(&modulo, big.NewInt(4))

	println(cuadratics[0].Exp(&cuadratics[0], &modulo, p).String())
}
