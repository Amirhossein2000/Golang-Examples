package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := addchaninchan(createchan("b"), createchan("a"))
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
}

func createchan(str string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%v %v", str, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
		}
	}()
	return c
}

func addchaninchan(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
