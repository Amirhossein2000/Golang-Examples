package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			c <- x
		}
	}()
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}
