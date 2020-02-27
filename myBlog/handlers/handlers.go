package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type PageData struct {
	Title       string
	Description string
	Status      bool
}

func Router() *mux.Router {

	rout := mux.NewRouter()

	rout.HandleFunc("/", home()).Methods("GET")
	rout.HandleFunc("/about", about()).Methods("GET")
	rout.HandleFunc("/blog", blog()).Methods("GET")
	rout.HandleFunc("/blog/{id:[0-9]+}", detail()).Methods("GET")
	rout.HandleFunc("/contact", contact()).Methods("GET")
	rout.Handle("send_contact", sendform()).Methods("POST")

	// This will serve files under http://localhost:8080/assets/css//<filename>
	rout.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	return rout
}
