package main

import (
	"fmt"
	"strconv"
)

func main() {

	numStr := "123444"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed Integer:", num+num)

	numPint, err := strconv.ParseInt(numStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed Integer:", numPint*3)

	floatstr := "2.56"
	floatval, err := strconv.ParseFloat(floatstr, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parsed Float: %.2f\n", floatval*2)

	binaryStr := "1010"
	decimal, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed Binary to Decimal:", decimal)

	hexStr := "F0"
	hex, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Parsed Hex to Decimal:", hex)

}
