package data

import "fmt"

var Names [2]string

func init() {
	Names[0] = "Bill"
	Names[1] = "Ted"
	fmt.Println(Names)
}
