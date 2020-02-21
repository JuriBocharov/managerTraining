package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func detail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		response := fmt.Sprintf("Product %s", id)
		_, err := fmt.Fprint(w, response)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
