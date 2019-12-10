package main

import "fmt"

type mystructe struct {
	num int
}

func main() {
	myMp := make(map[string]mystructe)

	//this is wrong you must override
	myMp["str"] = mystructe{10}
	myMp["str"].num++

	fmt.Println(myMp)
}
