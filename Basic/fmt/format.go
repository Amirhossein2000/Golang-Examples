package main

import (
	"fmt"
)

func main() {
	for x := 1; x < 33; x++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", x, x, x, x)
	}
}
