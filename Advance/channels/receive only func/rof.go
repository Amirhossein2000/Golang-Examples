package main

import (
	"fmt"
)

func main() {
	for n := range sum(crate()) {
		fmt.Println(n)
	}
}

func crate() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10000; i++ {
			out <- i
		}
		fmt.Println("created")
		close(out)
	}()
	return out
}

func sum(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		fmt.Println("sumed")
		close(out)
	}()
	return out
}
