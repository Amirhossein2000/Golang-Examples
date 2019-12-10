package main

import (
	"fmt"
)

func ex() func() {
	return func() {
		fmt.Println("a func has returned by another func")
	}
}
func ex2() func() string {
	return func() string {
		return "its a func with string return value"
	}
}
func main() {
	exfunc := func() {
		fmt.Println("its a expression func !!! ")
	}
	exfunc()
	e1 := ex()
	e1()
	e22 := ex2()
	fmt.Println(e22())
}
