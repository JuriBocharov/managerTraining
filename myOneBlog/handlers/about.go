package handlers

import (
	"html/template"
	"net/http"
)

func about() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			Title:       "Контакты",
			Description: "Информация как со мной связаться",
		}

		header, _ := template.ParseFiles("web/header.gohtml")
		header.Execute(w, data)

		body, _ := template.ParseFiles("web/about.gohtml")
		body.Execute(w, nil)

		footer, _ := template.ParseFiles("web/footer.gohtml")
		footer.Execute(w, nil)
	}
}
