package main

import "fmt"

func main() {
	for {
		fmt.Println("in first loop")
		for {
			fmt.Println("in second loop")
			break
		}
	}
}
