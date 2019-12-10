package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	r, err := ioutil.ReadFile("./hello.txt")
	checkerr(err)
	fmt.Printf("the fucking result :-->>\n%v", string(r))
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
