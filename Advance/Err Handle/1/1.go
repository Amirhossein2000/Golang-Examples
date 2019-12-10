package main

import (
	"fmt"
	"log"
	"os"
)

func ErrorCheck(err error) {
	if err != nil {
		log.Println(err, " this is log.pl")
		fmt.Println(err, " this is fmt.pl")
		fmt.Print("this is fatal :")
		log.Fatal(err)
		fmt.Print("\n")
	}
}

func panicex(err error) {
	panic(err)
}

func main() {
	_, err := os.Open("1.txt")
	panicex(err) //only one of these works --  they dont work to gether
	// ErrorCheck(err)
}
