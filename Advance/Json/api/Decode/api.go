package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type j struct {
	Ticker    jj
	Timestamp int
	Success   bool
	Err       string
}
type jj struct {
	Base   string
	Target string
	Price  string
	Volume string
	Change string
}

func main() {
	var b1 j
	resp, _ := http.Get("https://api.cryptonator.com/api/ticker/btc-usd")
	bytes, _ := ioutil.ReadAll(resp.Body)
	rd := strings.NewReader(string(bytes))
	json.NewDecoder(rd).Decode(&b1)
	if b1.Success {
		fmt.Printf("the price of bitcoin: %v\nthe curency: %v\n",
			b1.Ticker.Price, b1.Ticker.Target)
	} else {
		fmt.Println(b1.Err)
	}
}
