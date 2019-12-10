package main

import (
	"fmt"
)

func main() {
	t := 19.78
	p := &t
	fmt.Print(t, "\n", *p)
	*p = *p + 10
	fmt.Print("\n", t, *p)
}
