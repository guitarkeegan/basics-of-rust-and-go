package main

import "fmt"

func main() {
	x := 4
	y := 5
	// Add is available because of the package main
	ans := Add(x, y)

	// use fmt to make the console log universal
	fmt.Println(ans)
}
