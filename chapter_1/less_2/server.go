package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "8080"
	log.Printf("Установлен порт по умолчанию %s", port)
	log.Printf("Слушаем порт %s", port)
	log.Printf("Откройте http://localhost:%s в своем браузере", port)

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Привет будущие разработчики!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
