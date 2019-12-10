package main

//this program has race : many times over riding on a variable
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

func incore(s string) {
	for x := 0; x < 20; x++ {
		i := counter
		i++
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		counter = i
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
