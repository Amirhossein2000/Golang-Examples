package main

import (
	"fmt"
	"net/url"
)

type myUrl url.URL

func main() {
	old, _ := url.Parse("http://10.104.17.26:5080/RPC2")

	fmt.Printf(old.Hostname())
}
