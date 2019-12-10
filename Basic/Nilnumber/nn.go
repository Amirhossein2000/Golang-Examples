package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)
	fmt.Printf("value ==> %v \ttype ==> %T \n", a, a)
	fmt.Printf("value ==> %v \ttype ==> %T \n", b, b)
	fmt.Printf("value ==> %v \ttype ==> %T \n", c, c)
	fmt.Printf("value ==> %v \ttype ==> %T \n", d, d)
}
