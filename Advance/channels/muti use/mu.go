package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for x := 0; x < 10000; x++ {
			c <- x
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println("from chanel az a", n)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println("from chanel az b", n)
		}
		done <- true
	}()
	<-done
	<-done
}
