package main

import "fmt"

func printnumbers(number []int, p func(int)) {
	for _, n := range number {
		p(n)
	}
}

func main() {
	printnumbers([]int{1, 2, 3, 4, 5}, func(x int) {
		fmt.Println(x)
	})
}
