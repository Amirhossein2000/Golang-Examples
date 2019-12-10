package main

import (
	"encoding/json"
	"fmt"
)

type myjson struct {
	Name  string
	Lname string
	Age   int
	Ep    int
}

func main() {
	j := myjson{"Amirhossein", "Baniasadi", 18, 15000}
	jj, _ := json.Marshal(j)
	fmt.Println(jj)
	fmt.Printf("%T\n", jj)
	fmt.Print(string(jj))
}
