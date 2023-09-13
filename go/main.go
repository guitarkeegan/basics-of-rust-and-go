package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Millisecond * 800)
	}
	channel <- "done!"
}

func main() {
	channel := make(chan string)
	go printMessage("What is for dinnaaa????!", channel)
	res := <-channel
	fmt.Println(res)
}
