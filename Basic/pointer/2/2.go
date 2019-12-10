package main

import (
	"fmt"
)

func makezero(n *int) {
	*n = 0
}
func makezero2(n int) {
	n = 0
}
func main() {
	x := 5
	makezero(&x)
	fmt.Print(x, "\n") //the value is 0
	y := 5
	makezero2(y)
	fmt.Print(y) //steal 5
}
