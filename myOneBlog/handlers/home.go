package handlers

import (
	"html/template"
	"net/http"
)

func home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		header, _ := template.ParseFiles("assets/web/header.gohtml")
		header.Execute(w, nil)

		footer, _ := template.ParseFiles("assets/web/footer.gohtml")
		footer.Execute(w, nil)
	}
}
