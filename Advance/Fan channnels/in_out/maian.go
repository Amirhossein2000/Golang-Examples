package main

import (
	"fmt"
	"sync"
)

func main() {
	in := create(2, 3, 4, 5)
	// c:=sqf(in)
	// c2:=sqf(in)
	for x := range merge(sqf(in), sqf(in)) {
		fmt.Println(x)
	}

}

func create(nums ...int) chan int {
	c := make(chan int)
	go func() {
		for _, i := range nums {
			c <- i
		}
		close(c)
	}()
	return c
}
func sqf(c chan int) chan int {
	sq := make(chan int)
	go func() {
		for l := range c {
			sq <- l * l
		}
		close(sq)
	}()
	return sq
}
func merge(c ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(c))
	for _, i := range c {
		go func(ch chan int) {
			for j := range ch {
				out <- j
			}
			wg.Done()
		}(i)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
