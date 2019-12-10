package main

import (
	"fmt"
	"os"
)

func main() {

	macArray := make([]byte, 6)
	file, _ := os.Create("inc-macs.txt")
	for i := 1000000; i > 0; i-- {
		if macArray[0] == 238 {
			for j := 1; j <= 5; j++ {
				if macArray[j] < 238 {
					macArray[j]++
					for n := j - 1; n >= 0; n-- {
						macArray[n] = 0
					}
					goto inclbl
				}
			}
		}

		//		macArray[0]++

	inclbl:
		file.WriteString(fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x\n",
			macArray[5],
			macArray[4],
			macArray[3],
			macArray[2],
			macArray[1],
			macArray[0],
		))
		macArray[0]++

	}
}
