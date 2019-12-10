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

func (k Son) nameplusage() string {
	return fmt.Sprintf("%v %v", k.name, k.age)
}
func (k Father) nameplusage() string {
	return fmt.Sprintf("%v %v", k.name, k.age)
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
	fmt.Printf("the futher age+name: %v\nthe son age+name: %v",
		p1.Father.nameplusage(),
		p1.nameplusage())
}
