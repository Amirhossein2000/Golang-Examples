package main

import (
	"os"
	//"fmt"
	"encoding/json"
)

type myjson struct {
	Name  string
	Lname string
	Age   int
	Ep    int
}

func main() {
	j := myjson{"Amirhossein", "Baniasadi", 18, 15000}
	json.NewEncoder(os.Stdout).Encode(j)
}
