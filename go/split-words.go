package main

import (
	"fmt"
	"strings"
)

func splitWords(s []string, downstream chan []string) {
	for _, val := range s {
		downstream <- strings.Fields(val)
	}
	close(downstream)
}

func getWords(upstream chan []string) {
	for val := range upstream {
		for _, v := range val {
			fmt.Println(v)
		}
	}
}
