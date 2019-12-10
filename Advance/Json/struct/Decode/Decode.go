package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type myjson struct {
	Name  string
	Lname string
	Age   int
	Ep    int
}

func main() {
	var p1 myjson
	rd := strings.NewReader(`{"Name":"Amirhossein","Lname":"Baniasadi","Age":18,"Ep":2000}`)
	json.NewDecoder(rd).Decode(&p1)
	fmt.Print(p1)
}
