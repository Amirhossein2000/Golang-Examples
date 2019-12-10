package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	done := make(chan bool)

	go func() {
		for x := 0; x < 10; x++ {
			c <- fmt.Sprintf("a:%v", x)
		}
		fmt.Println("a done")
		done <- true
	}()

	go func() {
		for x := 0; x < 10; x++ {
			c <- fmt.Sprintf("b:%v", x)
		}
		fmt.Println("b done")
		done <- true
	}()

	go func() {
		fmt.Println(<-done)
		// <-done
		fmt.Println(<-done)
		// <-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
