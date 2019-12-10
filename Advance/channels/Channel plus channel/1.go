package main

import (
	"fmt"
)

func main() {
	fmt.Println(<-sum(crate("b")) + <-sum(crate("a")))
}

func crate(name string) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 30; i++ {
			out <- i
			fmt.Println(name, i)
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
