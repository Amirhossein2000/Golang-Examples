package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a Number --> ")
	var input int
	fmt.Scan(&input)
	fmt.Printf("%v! = %v", input, <-fact(create(input)))
}

func create(n int) chan int {
	c := make(chan int)
	go func() {
		for i := 1; i < n+1; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
func fact(c chan int) chan int {
	facchan := make(chan int)
	fac := 1
	go func() {
		for l := range c {
			fac *= l
		}
		facchan <- fac
		close(facchan)
	}()
	return facchan
}
