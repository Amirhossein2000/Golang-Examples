package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			c <- x
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
