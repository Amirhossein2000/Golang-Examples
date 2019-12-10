package macutils

import "fmt"

func GenerateMacList(i int) []string {
	macArray := make([]byte, 6)
	macList := []string{}

	for ; i > 0; i-- {
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

		macArray[0]++

	inclbl:
		macList = append(macList,
			fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
				macArray[5],
				macArray[4],
				macArray[3],
				macArray[2],
				macArray[1],
				macArray[0],
			))
	}

	return macList
}
