package main

import (
	"fmt"
)

type Dog struct {
	name string
	age  int
}

func main() {
	peter := Dog{"Peter", 4}
	fmt.Println(peter)
	fmt.Println(peter.name)
	fmt.Println(peter.age)
	fmt.Printf("%+v \n", peter)
}
