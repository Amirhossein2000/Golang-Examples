package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello and welcome to go web Sever !!!!</h1>")
}
func main() {
	var h hello
	err := http.ListenAndServe("localhost:4000", h)
	checkerror(err)
}
func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}
