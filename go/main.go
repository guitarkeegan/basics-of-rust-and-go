package main

import (
	"fmt"
	"net/http"
	"text/template"

	"femm.com/server/api"
	"femm.com/server/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from go server"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/new", api.Post)

	err := http.ListenAndServe(":3000", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}
