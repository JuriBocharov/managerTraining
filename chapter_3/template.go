package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := "Go Template"
		tmpl, _ := template.New("data").Parse("<h1>{{ .}}</h1>")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
