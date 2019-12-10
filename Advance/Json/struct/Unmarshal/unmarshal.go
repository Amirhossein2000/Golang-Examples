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
	var p1 myjson
	a := `{"Name":"Amirhossein","Lname":"Baniasadi","Age":18,"Ep":10000}`
	bytes := []byte(a)
	json.Unmarshal(bytes, &p1)
	fmt.Println(p1)
	fmt.Println(a)
	fmt.Println(bytes)
}
