package main

import (
	"fmt"
)

func main() {
	var herher [3]string
	var number = [5]int{0, 1, 2, 3, 4}
	herher[0] = "hi"
	herher[1] = "whats up ???!!???"
	herher[2] = "GOODbye"
	fmt.Print(herher, "\n", number)
}
