package main

import (
	"fmt"
)

func main() {
	// for n:=range sqf(create(1,2,3,4,5)){
	// 	fmt.Println(n)
	// }
	// for n:=range sqf(sqf(create(1,2,3,4,5))){
	// 	fmt.Println(n)
	// }
	// for n:=range sqf(sqf(sqf(create(1,2,3,4,5)))){
	// 	fmt.Println(n)
	// }
	c := sqf(create(2, 3, 4, 5))
	for i := range create(2, 3, 4, 5) {
		fmt.Printf("%v ^2 = %v\n", i, <-c)
	}
	fmt.Printf("##################\n")
	c = sqf(sqf(create(2, 3, 4, 5)))
	for i := range sqf(create(2, 3, 4, 5)) {
		fmt.Printf("%v ^2 = %v\n", i, <-c)
	}
	fmt.Printf("##################\n")
	c = sqf(sqf(sqf(create(2, 3, 4, 5))))
	for i := range sqf(sqf(create(2, 3, 4, 5))) {
		fmt.Printf("%v ^2 = %v\n", i, <-c)
	}
	fmt.Print("i am king of programming")
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
