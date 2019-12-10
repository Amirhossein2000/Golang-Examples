package main

import (
	"fmt"
)

type Dog struct {
	age    int
	weight int
	name   string
}

func (g Dog) printage() {
	fmt.Println(g.age)
}

// func (g Dog) printname() {
// 	g.name=fmt.Sprintf("%v  %v  %v",g.name,g.name,g.name)
// 	fmt.Println(g.name)
// }
func (g *Dog) printnametreetimes() {
	g.name = fmt.Sprintf("%v  %v  %v", g.name, g.name, g.name)
	fmt.Println(g.name)
}
func main() {
	k := Dog{10, 50, "Jerard"}
	fmt.Println(k)
	k.printage()
	k.age = 9
	k.printage()
	k.printnametreetimes()
	k.printnametreetimes()
}
