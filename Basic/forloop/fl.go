package main

import (
	"fmt"
)

func main() {
	amir := []string{"amir", "hossein", "baniasadi", "is", "the worst", "hacker"}
	for x := range amir {
		fmt.Println(x, "-->", amir[x])
	}
	fmt.Println("======== its new ========")
	for x := 0; x < len(amir); x++ {
		fmt.Println(x, "-->", amir[x])
	}
	fmt.Println("======== its new ========")
	for sum := 1; sum < 1024; {
		sum = sum + sum
		fmt.Println(sum)
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		if i > 29 {
			goto end
		}
		if i > 70 {
			break
		}
	}
end:
	fmt.Println("This is end of this program")
}
