package handlers

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	rout := mux.NewRouter()
	rout.HandleFunc("/", home()).Methods("GET")
	rout.HandleFunc("/post/{id:[0-9]+}", detail()).Methods("GET")

	return rout
}
