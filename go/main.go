package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go server"))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	err := http.ListenAndServe(":3000", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}
