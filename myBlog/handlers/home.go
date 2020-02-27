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

		/*post_template := template.Must(
			template.ParseFiles(
				path.Join("templates", "./tmpl/header.gohtml"),
				path.Join("templates", "./tmpl/footer.gohtml")))
		err := post_template.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}*/
		header, _ := template.ParseFiles("./tmpl/header.gohtml")
		header.Execute(w, data)

		body, _ := template.ParseFiles("./tmpl/home.gohtml")
		body.Execute(w, data)

		footer, _ := template.ParseFiles("./tmpl/footer.gohtml")
		footer.Execute(w, nil)
	}
}
