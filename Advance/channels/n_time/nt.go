package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	n := 3
	counter := 0

	for i := 0; i < n; i++ {
		go func() {
			for x := 0; x < 10; x++ {
				c <- x
			}
			counter++
			fmt.Println("done : ", counter)
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for n := range c {
		fmt.Println("in chanel:", n)
	}
}
