package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(forReturn())
}

func forReturn() string {
	for n := 0; n < 10; n++ {
		return strconv.Itoa(n)
	}

	return ""
}
