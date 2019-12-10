package main

import (
	"fmt"
	"os"
)

func init() {
	os.Create("example.txt")
}

func main() {
	for i := 0; i < 10; i++ {
		f, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0600)

		if err != nil {
			panic(err)
		}
		if _, err = f.WriteString(fmt.Sprintf("%v\n", i)); err != nil {
			panic(err)
		}

		f, err = os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}

		f.Close()
	}
}
