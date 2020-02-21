package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func productsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Product %s", id)
	_, err := fmt.Fprint(w, response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/news/{id:[0-9]+}", productsHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
