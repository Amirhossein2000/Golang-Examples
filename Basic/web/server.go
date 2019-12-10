package main

import (
	"fmt"
	"log"
	"net/http"
)

func hiserver(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	fmt.Fprintf(w, "<h1>HI , welcome ,Try to have some fun!!</h1> \n Url = %s \n", req.URL)
}
func main() {
	fmt.Println("connect please")
	http.HandleFunc("/hello", hiserver)
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
