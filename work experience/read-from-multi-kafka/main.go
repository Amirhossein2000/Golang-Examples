package main

import (
	"fmt"
	"time"
)

type kafkaServer struct {
	ip     string
	name   string
	weight uint
}

func (ks kafkaServer) connect() {
	fmt.Println("connected to", ks.ip, ks.name)
	for i := uint(0); i < ks.weight; i++ {
		go ks.writename()
	}
}

func (ks kafkaServer) writename() {
	fmt.Println(ks.name)
}

func main() {
	serverList := []kafkaServer{
		{"1.1.1.1", "A", 1},
		{"1.1.1.2", "B", 2},
		{"1.1.1.3", "C", 3},
		{"1.1.1.4", "D", 4},
		{"1.1.1.5", "E", 5},
	}

	for _, ks := range serverList {
		go ks.connect()
	}

	time.Sleep(time.Second * 3)
}
