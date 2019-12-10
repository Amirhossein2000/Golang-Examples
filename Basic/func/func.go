package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(plus(1000, -7000))
	fmt.Println(addallvalues(89, 100, 67, 50))
	for x := 0; x < 10; x++ {
		for i := 0; i < 4; i++ {
			n1 := rand.Intn(899) + 100
			n2 := rand.Intn(899) + 100
			n3 := rand.Intn(899) + 100
			n4 := rand.Intn(899) + 100
			fmt.Printf("%v + %v + %v + %v = %v \n", n1, n2, n3, n4, addallvalues(n1, n2, n3, n4))
		}
	}
}

func plus(v, v2 int) int {
	return v + v2
}

func addallvalues(v ...int) int {
	sum := 0
	for i := range v {
		sum += v[i]
	}
	return sum
}
