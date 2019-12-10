package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var mapjson map[string]map[string]string
	var mapinterface map[string]interface{}

	res, err := http.Get("https://api.cryptonator.com/api/ticker/btc-usd")
	checkerr(err)
	byemarshal, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(byemarshal, &mapjson)
	_ = json.Unmarshal(byemarshal, &mapinterface)

	fmt.Println(mapjson)
	fmt.Println(mapinterface)

}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
