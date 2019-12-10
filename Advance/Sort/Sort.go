package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []string{"Jan", "JJJ", "Amirhossein", "Jjk", "nick", "Khashayar"}
	var b = []int{5, 6, 1, 2, 3, 0, 4, 11, 5}
	var c = []string{"amir", "amirreza", "amirmohahamd", "amirhossein"}
	sort.Sort(sort.StringSlice(a))
	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	sort.Sort(sort.StringSlice(c))
	sort.Sort(sort.IntSlice(b))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println("---------------------------------------------------------")

	sort.Sort(sort.Reverse(sort.StringSlice(a)))
	sort.Sort(sort.Reverse(sort.StringSlice(c)))
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
