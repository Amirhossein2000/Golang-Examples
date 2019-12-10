package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("adasd.txt")
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err)
	}
	myerr := errors.New("This an Error!")
	Name := map[string]bool{
		"Amirhossein": true,
		"khashayar":   true}
	k, ok := Name["Amirhossein"]
	if ok {
		fmt.Println("the name has a value?", k)
	} else {
		fmt.Println("the name has no value!!!")
	}
	fmt.Println(k, ok)
	k, ok = Name["kkk"]
	if ok {
		fmt.Println("the name has a value?", k)
	} else {
		fmt.Println("the name has no value!!!")
	}
	fmt.Println(k, ok)
	fmt.Println(myerr)
}
