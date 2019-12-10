package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Since(time.Now()))
	fmt.Println(time.Since(time.Now().Add(time.Hour)))

	fmt.Println(time.Now().Format("20060102"))
}
