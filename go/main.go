package main

import (
	"fmt"

	"reading.com/files/filereader"
)

func main() {
	data, err := filereader.ReadTextFile("data/text.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("here is the data: %s", data)
	}
}
