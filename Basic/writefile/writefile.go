package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("./file.txt")
	checkerr(err)
	defer file.Close()
	len, err := io.WriteString(file, "Hello its a file !")
	checkerr(err)
	fmt.Printf("file with %v charactors Created \n", len)
	bytes := []byte("It was a binary and changed to txt !")
	fmt.Println(bytes)
	ioutil.WriteFile("./binary.txt", bytes, 0)
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
