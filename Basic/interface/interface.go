package main

import (
	"fmt"
)

type animal interface {
	speak() string
}
type Dog struct {
}

func (d Dog) speak() string {
	return "hop!"
}

type Cat struct {
}

func (c Cat) speak() string {
	return "Meow!"
}

type Cow struct {
}

func (c Cow) speak() string {
	return "Moo!"
}
func main() {
	huscky := animal(Dog{})
	fmt.Println(huscky)
	huscky.speak()
	animals := []animal{Dog{}, Cat{}, Cow{}}
	for a, b := range animals {
		fmt.Println(a, b, b.speak())
	}
	fmt.Print(animals)
}
