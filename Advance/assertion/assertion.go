package main

import (
	"fmt"
)

func main() {
	var name interface{} = "Airhossein"
	str, ok := name.(string)
	if ok {
		fmt.Printf("this is string --> %v is %T\n", str, str)

	} else {
		fmt.Print("this is not string \n")
	}
	var number interface{} = 11
	fmt.Println(number)
	fmt.Println(number.(int) + 11)
}
