package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Resp value --> %v \nResp type --> %T", resp, resp)
	defer resp.Body.Close()
	byte, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byte))
}
