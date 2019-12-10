package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["key"] = "amirhossein"
	m["key2"] = "Baniasadi"
	m["key3"] = "1379"
	m["key4"] = "2000"
	fmt.Print(m["key"], "\n")
	fmt.Print(m, "\n")
	for k, v := range m {
		fmt.Print(k, " ==>> ", v, "\n")
	}
	keys := make([]string, len(m))
	i := 0
	for ke := range m {
		keys[i] = ke
		i++
	}
	fmt.Print(keys, "\n")
	i = 0
	for x := range keys {
		fmt.Print(m[keys[x]], "\n")
	}
	for _, f := range m {
		fmt.Print(f, "\n")
	}
}
