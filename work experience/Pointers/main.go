package main

import "fmt"

type dontKnow struct {
	data string
}

func (dk dontKnow) setData() {
	dk.data = ""
}

func (dk *dontKnow) setPointerData() {
	dk.data = ""
}

func main() {
	var Varrr int = 10

	var b *int = &Varrr

	fmt.Println(Varrr, &Varrr)

	fmt.Println(b, *b, &b)

	dk := dontKnow{"this is data"}
	dk.setData()
	fmt.Println(dk.data)
	(&dk).setPointerData()
	fmt.Println(dk.data)
}
