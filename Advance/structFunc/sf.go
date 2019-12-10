package main

import (
	"fmt"
)

type Properties struct {
	name        string
	lname       string
	age         int
	phonenumber string
}

func (person Properties) fullname() string {
	return person.name + " " + person.lname
}
func main() {
	person := Properties{"Amirhossein", "Baniasadi", 18, "09120634185"}
	fmt.Printf("the name: %v\nthe lname: %v\nthe age: %v\nphonenumber: %v\n",
		person.name, person.lname, person.age, person.phonenumber)
	fmt.Printf("full name : %v", person.fullname())
}
