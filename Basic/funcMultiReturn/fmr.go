package main

import (
	"fmt"
)

func main() {
	f, l := fullname("Amirhossein", "Baniasadi")
	fmt.Printf("Full name ==>> %v len ==>> %v \n\n", f, l)
	fmt.Println(fullnamenaked("khashayar", "Rajayi"))
}

func fullname(f, l string) (string, int) {
	return (f + " " + l), (len(f + l))
}
func fullnamenaked(f, l string) (full string, lenn int) {
	full = f + " " + l
	lenn = len(full)
	return
}
