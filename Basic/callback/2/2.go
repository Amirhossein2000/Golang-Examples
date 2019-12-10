package main

import "fmt"

func biggerthanone(numbers []int, callback func(int) bool) []int {
	x := []int{}
	for _, i := range numbers {
		if callback(i) {
			x = append(x, i)
		}
	}
	return x
}

func main() {
	fmt.Println(biggerthanone([]int{1, 2, 3, 4, 5, 1, 0, 1, 1}, func(n int) bool {
		return n > 1
	}))
}
