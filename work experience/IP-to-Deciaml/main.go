package main

import (
	"fmt"
	"math/big"
	"net"
)

type testTable struct {
	input  net.IP
	output int64
}

func main() {
	tb := []testTable{
		{net.ParseIP("31.184.177.2"), 532197634},
		{net.ParseIP("171.1.1.1"), 2868969729},
		{net.ParseIP("128.12.1.10"), 2148270346},
		{net.ParseIP("198.5.61.33"), 3322232097},
	}

	for _, ttbb := range tb {
		ipv := IP4toInt(ttbb.input)
		if ttbb.output != ipv {
			fmt.Println("fail", ipv)
		}
		fmt.Println(ttbb, ipv)
	}
}

func IP4toInt(IPv4Address net.IP) int64 {
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(IPv4Address.To4())
	return IPv4Int.Int64()
}
