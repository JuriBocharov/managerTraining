package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := ViewData{
			Title:   "Привет Мир!",
			Message: "Мы все живем на одном огромном куске астероида",
		}
		tmpl, _ := template.ParseFiles("templates/index.gohtml")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
