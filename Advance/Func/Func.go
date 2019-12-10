package main

import (
	"fmt"
)

func plusone(a int) {
	a = a + 1
}

func plusonepointer(a *int) {
	*a = *a + 1
}

func main() {
	a := 10
	plusone(a)
	fmt.Println(a)
	plusonepointer(&a)
	fmt.Println(a)
}
