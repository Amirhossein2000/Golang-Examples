package main

import (
	"fmt"
)

func main() {
	switch x := 10; x {
	case 10:
		fmt.Println("its Ten")
	default:
		fmt.Println("its Not ten")
	}
	switch x := 5; {
	case x > 0:
		fmt.Println("x > 0")
	case x == 0:
		fmt.Println("x = 0")
	case x < 0:
		fmt.Println("x < 0")
	}
}
