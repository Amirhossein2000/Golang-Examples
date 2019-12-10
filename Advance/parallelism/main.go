package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //give number
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for x := 0; x < 20; x++ {
		fmt.Println("foo:", x)
		time.Sleep(time.Duration(time.Second))
	}
	wg.Done()
}
func bar() {
	for x := 0; x < 20; x++ {
		fmt.Println("bar:", x)
		time.Sleep(time.Duration(time.Second))
	}
	wg.Done()
}
