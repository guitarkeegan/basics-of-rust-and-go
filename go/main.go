package main

import "fmt"

// * is a pointer to the value, otherwise it would just be a copy
func turnClockBack(age *int) {
	*age -= 1
}

func main() {
	age := 40
	// & will give the address to the value in memory
	turnClockBack(&age)
	fmt.Printf("I am now %d!\n", age)
}
