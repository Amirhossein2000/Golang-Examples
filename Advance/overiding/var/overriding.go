package main

import (
	"fmt"
)

type Father struct {
	name string //same name in Son struct
	age  int    //same name in Son struct
}

type Son struct {
	name string //same name in Father struct
	age  int    //same name in Father struct
	Father
}

func main() {
	p1 := Son{
		Father: Father{
			name: "Mohammad",
			age:  50,
		},
		name: "Amirhossein",
		age:  18,
	}
	fmt.Printf("The name of person: %v\nThe age of person: %v\nthe person's Fathername: %v\nage of his Father: %v",
		p1.name, p1.age, p1.Father.name, p1.Father.age)
}
