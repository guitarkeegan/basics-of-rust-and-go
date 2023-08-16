package main

import "fmt"

func WhatTimeIsIt(time *int) string {
	if *time == 1 {
		return "yay!!"
	} else {
		panic("the time can't be anything but 1!!!")
	}
}

func main() {
	// defer will cause whatever follows to be called last in the
	// function
	// notice that the panic is printed after the WhatTimeIsIt(&clocky)
	badTime := 34
	defer WhatTimeIsIt(&badTime)

	// panic will exit a program
	clocky := 1
	res := WhatTimeIsIt(&clocky)
	fmt.Println(res)

}
