package main

import (
	"fmt"
)

type Father struct {
	name string
	age  int
}

type Son struct {
	Name string
	Age  int
	Father
}

func main() {
	p1 := Son{
		Father: Father{
			name: "Mohammad",
			age:  50,
		},
		Name: "Amirhossein",
		Age:  18,
	}
	fmt.Printf("The name of person: %v\nThe age of person: %v\nthe person's Fathername: %v\nage of his Father: %v",
		p1.Name, p1.Age, p1.Father.name, p1.Father.age)
}
