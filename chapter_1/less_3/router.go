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
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "hello.html")
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
	_, err := fmt.Fprint(w, "<h1>Привет! Главна стрница.</h1>")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
