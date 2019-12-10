package main

import (
	"fmt"
	"strings"
)

func main() {

	var netBillClassPure []string
	str := "2155591729|92|502b.7326.d04a|520176440|0|0|344680552|[P:55591729|S:ADSL|C:dPLaaR8pnGJJRhVi2MCr3KcofuuCt4MO234tmdoupgQ_0pbaVQhyW-BSEAWeMfpXHWn|23e24w]"

	for _, v := range strings.Split(str, "[") {
		fmt.Println(v)
		for _, s := range strings.Split(v[:len(v)-1], "|") {
			netBillClassPure = append(netBillClassPure, s)
		}
	}

	for k, v := range netBillClassPure {
		fmt.Println(k, v)
	}
}
