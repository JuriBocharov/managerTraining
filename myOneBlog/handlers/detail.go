package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type ViewData struct {
	postId  string
	Message string
}

func detail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		response := fmt.Sprintf("Product %s", id)

		data := ViewData{
			postId:  response,
			Message: "Некоторое сообщение",
		}

		header, _ := template.ParseFiles("assets/web/header.gohtml")
		header.Execute(w, nil)

		body, _ := template.ParseFiles("assets/web/body.gohtml")
		body.Execute(w, data)

		footer, _ := template.ParseFiles("assets/web/footer.gohtml")
		footer.Execute(w, nil)

	}
}
