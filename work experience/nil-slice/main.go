package main

import "fmt"

func main() {
	strSlice := []string{}

	fmt.Println("strSlice == nil -->", strSlice == nil)

	strSlice = nil

	fmt.Println("strSlice is nullable -->", strSlice == nil)
}
