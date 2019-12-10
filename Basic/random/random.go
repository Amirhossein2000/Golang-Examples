package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for x := 0; x <= 5; x++ {
		fmt.Print(rand.Intn(100), " | ")
		fmt.Print(string(rand.Intn(400)+200), "\n")
	}
	r := rand.Intn(7) + 1
	var massage string
	switch r {
	case 1:
		massage = "its sunday"
	case 2:
		massage = "its Mondey"
	default:
		massage = "its a weekday!!!!!!!!"
	}
	fmt.Print("day ", r, " is ", massage)
}
