package main

import "fmt"

func beginWaterfall(downstream chan string) {
	strArr := []string{"hello", "googbye", "welcome", "welcome", "yoyo"}
	for _, val := range strArr {
		downstream <- val
	}
	close(downstream)
}

func removeIdentical(upstream chan string) {
	last := ""
	for val := range upstream {
		if last != val {
			fmt.Println(val)
		}
		last = val
	}
}
