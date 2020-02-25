package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("News %s", id)
	_, err := fmt.Fprint(w, response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/news/{id:[0-9]+}", productsHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
