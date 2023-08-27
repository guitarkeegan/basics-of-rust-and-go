package main

import (
	"fmt"
	"os"

	fileworker "reading.com/files/filereader"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	data, err := fileworker.ReadTextFile(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("here is the data: %s", data)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v\n", data, data)
		fileworker.WriteToFile(filePath+".ouput.txt", newContent)
	}
}
