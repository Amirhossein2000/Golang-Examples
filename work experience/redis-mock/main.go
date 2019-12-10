package main

import (
	"fmt"

	"github.com/alicebob/miniredis"
)

func main() {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	err = s.Set("kkk", "this is val")
	if err != nil {
		panic(err)
	}
	str, err := s.Get("kkk")
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}
