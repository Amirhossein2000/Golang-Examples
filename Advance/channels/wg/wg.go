package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for x := 0; x < 10; x++ {
			c <- x
		}
		wg.Done()
	}()
	go func() {
		for x := 0; x < 10; x++ {
			c <- x
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
