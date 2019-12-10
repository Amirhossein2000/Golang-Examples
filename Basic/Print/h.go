package main

import "fmt"

func main() {
	st1 := "Hello"
	st2 := "Amirhossein"
	st3 := "Welcome to GO"
	st4 := "goodbye"
	fmt.Print(st1, st2, st3, st4)
	fmt.Printf(st1, st2, st3, st4)
	fmt.Println(st1, st2, st3, st4)
	stringS, err := fmt.Println(st1, st2, st3, st4)
	fmt.Print(stringS, err)
	fmt.Print(fmt.Sprintf(st1, st2, st3, st4))
}
