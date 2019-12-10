package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 30, 1, 1, 1, 1, time.UTC)
	fmt.Printf("time = %v \n", t)
	fmt.Print(t.Day(), t.Year(), t.Month(), "\n")
	fmt.Println(t.AddDate(1, 1, 1).Year())
	fmt.Println(time.Now())
	fmt.Print(time.Now().Day())
}
