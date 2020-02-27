package handlers

import (
	"html/template"
	"net/http"
)

func contact() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			Title:       "Контакты",
			Description: "Вы всегда можете связаться со мной через эту форму или воспользуйтесь кнопками соцсетей и связаться со мной.",
		}

		header, _ := template.ParseFiles("tmpl/header.gohtml")
		header.Execute(w, data)

		body, _ := template.ParseFiles("tmpl/contact.gohtml")
		body.Execute(w, data)

		footer, _ := template.ParseFiles("tmpl/footer.gohtml")
		footer.Execute(w, nil)
	}
}
