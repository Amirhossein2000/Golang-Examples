package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 3)
	c <- "HI"
	go func() {
		for A := range c {
			fmt.Println(A)
		}

		fmt.Println("channel closed")
	}()

	time.Sleep(4 * time.Second)
	close(c)
	fmt.Println("closing")
	time.Sleep(4 * time.Second)
}
