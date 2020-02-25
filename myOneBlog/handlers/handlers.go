package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type PageData struct {
	Title       string
	Description string
}

func Router() *mux.Router {

	rout := mux.NewRouter()

	rout.HandleFunc("/", home()).Methods("GET")
	rout.HandleFunc("/about", about()).Methods("GET")
	rout.HandleFunc("/post/{id:[0-9]+}", detail()).Methods("GET")

	// This will serve files under http://localhost:8080/assets/css//<filename>
	rout.PathPrefix("/assets/css/").Handler(http.StripPrefix("/assets/css/", http.FileServer(http.Dir("assets/css/"))))

	return rout
}
