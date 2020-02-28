package handlers

import (
	"html/template"
	"net/http"
)

var result sendResult

func sendform() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			Title:       "Контакты",
			Description: "Информация как со мной связаться",
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		if name != "" && email != "" && message != "" {
			result = send(email, name, message)
		}

		header, _ := template.ParseFiles("tmpl/header.gohtml")
		header.Execute(w, data)

		body, _ := template.ParseFiles("tmpl/contact.gohtml")
		body.Execute(w, result)

		footer, _ := template.ParseFiles("tmpl/footer.gohtml")
		footer.Execute(w, nil)
	}
}
