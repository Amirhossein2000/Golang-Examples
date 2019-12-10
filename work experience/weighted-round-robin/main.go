package main

import "fmt"

func main() {
	channelA := make(chan int, 10)
	for i := 0; i < 10; i++ {
		channelA <- i
	}
	channelB := make(chan int, 20)
	for i := 0; i < 20; i++ {
		channelB <- i
	}
	channelC := make(chan int, 30)
	for i := 0; i < 30; i++ {
		channelC <- i
	}
	channelD := make(chan int, 40)
	for i := 0; i < 40; i++ {
		channelD <- i
	}
	channelList := []chan int{channelA, channelB, channelC, channelD}

	wList := []int{1, 2, 3, 4}

	for {
		for k, v := range wList {
			fmt.Println("channel number =", k)
			for n := 0; n < v; n++ {
				fmt.Println(<-channelList[k])
			}
		}
		fmt.Println("end of loop")
	}
}
