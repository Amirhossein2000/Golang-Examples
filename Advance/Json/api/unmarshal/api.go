package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type j struct {
	Ticker jj
	// Timestamp int
	Success bool
	Err     string
}
type jj struct {
	// Base string
	Target string
	Price  string
	// Volume string
	// Change string
}

func main() {
	var b1 j
	resp, _ := http.Get("https://api.cryptonator.com/api/ticker/btc-usd")
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &b1)
	resp.Body.Close()
	if b1.Success {
		fmt.Printf("the price of bitcoin: %v\nthe curency: %v",
			b1.Ticker.Price, b1.Ticker.Target)
	} else {
		fmt.Println(b1.Err)
	}
}
