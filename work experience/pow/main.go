package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(SubnetCount("127.1.1.0/24"))
}

//SubnetCount takes a range ip and returns the subnet count.
func SubnetCount(ipStrRange string) (int, error) {
	rangeNumber, _ := strconv.Atoi(strings.Split(ipStrRange, "/")[1])
	return int(math.Pow(2, float64(32-rangeNumber))), nil
}
