package main

import "fmt"

func main() {
	x := 4
	y := 5
	// Add is available because of the package main
	ans := Add(x, y)
	// use fmt to make the console log universal
	fmt.Println(ans)

	// array
	var Names [1]string
	Names[0] = "Bill"
	fmt.Println(Names[0])

	// dynamic list
	var Numbers = []int{1, 2, 3}
	Numbers = append(Numbers, 4)
	fmt.Println(Numbers)

	// map
	ids := map[string]int{"joe": 23, "bob": 22}
	fmt.Println(ids)

}
