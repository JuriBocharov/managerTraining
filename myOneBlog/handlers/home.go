package handlers

import (
	"html/template"
	"net/http"
)

func home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			Title:       "Главная",
			Description: "Мой первый блог на Golang",
		}

		header, _ := template.ParseFiles("web/header.gohtml")
		header.Execute(w, data)

		footer, _ := template.ParseFiles("web/footer.gohtml")
		footer.Execute(w, nil)
	}
}
