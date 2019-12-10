package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
)

var (
	str        string
	number     int64
	fb         big.Float
	f1, f2, f3 big.Float
)

func main() {
	f1.SetFloat64(67.73)
	f2.SetFloat64(78.88)
	f3.SetFloat64(56.34)
	fb.Add(&f1, &f2).Add(&fb, &f3)
	fmt.Printf("float= %0.10g \n", &fb)

	//% 2
	fmt.Print("Enter your number ==>> ")
	fmt.Scanln(&str)
	number, _ = strconv.ParseInt(strings.TrimSpace(str), 10, 64)
	if number%2 == 0 {
		fmt.Printf("%v is zoj \n", number)
		fmt.Printf("%v ^ 2 = %v \n", number, math.Pow(float64(number), 2.00))
		//fmt.Printf("%v + %v = %v \n",)
	} else {
		fmt.Printf("%v is Fard \n", number)
	}

}
