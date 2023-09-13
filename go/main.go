package main

import (
	"fmt"

	"basicasofgo.com/andrust/stack"
)

// func printMessage(text string, channel chan string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(text)
// 		time.Sleep(time.Millisecond * 800)
// 	}
// 	channel <- "done!"
// }

func main() {
	myStack := stack.NewStack()
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Peek()
	fmt.Println(myStack.Length)
	myStack.Pop()
	myStack.Peek()
	fmt.Println(myStack.Length)
	// channel := make(chan string)
	// go printMessage("What is for dinnaaa????!", channel)
	// res := <-channel
	// fmt.Println(res)
}
