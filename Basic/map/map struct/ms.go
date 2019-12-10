package main

import "fmt"

type arryinmap struct {
	name    string
	numbers []int
}

func main() {
	m := make(map[int]arryinmap)
	Numbers := []int{2, 3, 5, 76, 23}
	m[10] = arryinmap{"Amir", Numbers}
	fmt.Println(m[10].name)
	fmt.Println(m[10].numbers[2])

	m2 := make(map[int]arryinmap)
	var p arryinmap
	a := []int{1, 2, 3}
	p = arryinmap{"Amirhossein", a}
	m2[10] = p
	fmt.Println(m2[10].name)
	fmt.Println(m2[10].numbers[2])

	m3 := make(map[int]arryinmap)
	m3[10] = arryinmap{"AmirHossein", []int{2, 3, 5, 76, 23, 23, 44, 78}}
	fmt.Println(m3[10].name)
	fmt.Println(m3[10].numbers[4])
}
