package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int64
var wg sync.WaitGroup

func incore(s string) {
	for x := 1; x < 21; x++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, x, "Conter:", counter)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go incore("foo:")
	go incore("bar:")
	wg.Wait()
	fmt.Println("final counter :", counter)
}
