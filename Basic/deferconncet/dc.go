package main

import (
	"fmt"
)

var isconnect bool = false

func main() {
	fmt.Println("Connection open = ", isconnect)
	fmt.Println("Pool Request =====>>")
	dosomething()
	fmt.Println("pool Request Sent!")
	fmt.Println("============== answer of your pool reuest ==============")
	dosomething()
	fmt.Println("Pool Accepted ! your code is nice ")
}
func dosomething() {
	connect()
	fmt.Println("defer the disconnect")
	defer disconnect()
	fmt.Println("Doing some thing")
}
func connect() {
	isconnect = true
	fmt.Println("Connection open = ", isconnect)
}
func disconnect() {
	isconnect = false
	fmt.Println("Connection open = ", isconnect)
}
