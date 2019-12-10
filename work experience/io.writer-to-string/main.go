package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	var card1 = map[int]string{1: "S", 2: "H", 3: "C", 4: "D"}
	t := template.Must(template.New("t1").
		Parse("Dot:{{.}}\n"))

	B := bytes.NewBufferString("")
	t.Execute(B, card1)
	fmt.Println(B.String())
}
