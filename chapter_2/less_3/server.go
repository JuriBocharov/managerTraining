package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8080"

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Это страница About</h1>")
	})

	log.Printf("Установлен порт по умолчанию %s", port)
	log.Printf("Откройте http://localhost:%s в своем браузере", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "view/header.html")
	http.ServeFile(w, r, "view/body.tmpl")
	http.ServeFile(w, r, "view/footer.tmpl")
}
