package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func sqr(number float64) (float64, error) {
	if number < 0 {
		return -1, errors.New(fmt.Sprintf("Err : %v is Negative number", number))
	}
	return math.Sqrt(number), nil
}

func main() {
	// fmt.Println(sqr(100))
	// fmt.Print(sqr(-100))
	_, err := sqr(-10)
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}
