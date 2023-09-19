package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello from go server"))
	})
	err := http.ListenAndServe(":3000", nil)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}
