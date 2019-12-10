package main

import "fmt"

func fac(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return n * fac(n-1)
}

func main() {
	fmt.Println(fac(5))
}
