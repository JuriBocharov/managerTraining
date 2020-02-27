package handlers

import (
	"html/template"
	"net/http"
)

func blog() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			Title:       "Блог",
			Description: "Рассказываю о том, что считаю интересным",
		}

		header, _ := template.ParseFiles("tmpl/header.gohtml")
		header.Execute(w, data)

		body, _ := template.ParseFiles("tmpl/blog.gohtml")
		body.Execute(w, data)

		footer, _ := template.ParseFiles("tmpl/footer.gohtml")
		footer.Execute(w, nil)
	}
}
