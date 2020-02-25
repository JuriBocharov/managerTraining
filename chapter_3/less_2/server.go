package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title string
	Users []string
}

func main() {

	data := ViewData{
		Title: "Users List",
		Users: []string{"Tom", "Bob", "Sam"},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
