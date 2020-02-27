package main

import (
	"fmt"
	"log"
	"myBlog/handlers"
	"net/http"
)

func main() {
	port := "8080"
	log.Printf("Установлен порт по умолчанию %s", port)
	log.Printf("Откройте http://localhost:%s в своем браузере", port)

	router := handlers.Router()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
