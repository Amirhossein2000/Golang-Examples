package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("First")
	fmt.Println("Last")
	defer fmt.Println("Defer in main 1")
	defer fmt.Println("Defer in main 2")
	defer fmt.Println("Defer in main 3")
	defer fmt.Println("Defer in main 4")
	defer fmt.Println("Defer in main 5")
	dfffff()
	defer dfffff()
	fmt.Println("Undefer in main")
}
func dfffff() {
	defer fmt.Println("Defer in func 1")
	defer fmt.Println("Defer in func 2")
	defer fmt.Println("Defer in func 3")
	defer fmt.Println("Defer in func 4")
	fmt.Println("Undefer in Func")
}
