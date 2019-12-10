package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["amirhossein"] = "baniasadi"
	m["khashayar"] = "rajayi"
	m["Rasul"] = "Taremi"
	fmt.Println(m)
	delete(m, "amirhossein")
	fmt.Println(m)
}
