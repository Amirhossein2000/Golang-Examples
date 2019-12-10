package main

import (
	"fmt"
)

func main() {
	a := func(k, v int, sum *int) {
		*sum = k + v
	}
	k := 10
	v := 5
	sum := 0
	a(k, v, &sum)
	fmt.Printf("%v + %v = %v\n", k, v, sum)

	func() {
		fmt.Printf("%v + %v = %v", k, v, sum)
	}()
}
