package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex sync.Mutex
var counter int
var wg sync.WaitGroup

func incore(s string) {
	for x := 0; x < 20; x++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		counter++
		fmt.Println(s, x, "Conter:", counter)
		mutex.Unlock()
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
