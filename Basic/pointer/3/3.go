package main

import (
	"fmt"
)

func main() {
	age := 17
	fmt.Printf("my age in 2017 ==> %v \n", age)
	changeage(&age)
	fmt.Printf("my age in 2018 ==> %v", age)
}
func changeage(age *int) {
	*age = 18
}
