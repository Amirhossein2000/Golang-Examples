package main

import "fmt"
import "sort"

func main() {
	var a = []string{"amirhossein", "is", "king of learning !!!!", "are joone amat!", "2 bar"}
	fmt.Print(a, "\n")
	a = append(a, "3 bar")
	fmt.Print(a, "\n")
	a = append(a[2:])
	fmt.Print(a, "\n")
	a = append(a[0 : len(a)-1])
	fmt.Print(a, "\n")
	var number = make([]int, 5, 5)
	number[0] = 213
	number[1] = 88
	number[2] = 73
	number[3] = 23
	number[4] = 12
	fmt.Print(number)
	sort.Ints(number)
	fmt.Print("\n", number, "\n")
	fmt.Print("Now the capacity is = ", cap(number), "\n")
	number = append(number, 90)
	fmt.Print("Now the capacity is = ", cap(number))
}
