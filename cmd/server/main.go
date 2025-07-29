package main

import (
	"groupie-tracker/internal/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.HomeHandler)
	http.Handle("/web/assets/", http.StripPrefix("/web/assets/", http.FileServer(http.Dir("web/assets"))))

	http.HandleFunc("/artist/{id}", handler.ArtistHandler)

	log.Println("SERVER IS RUNNING ON http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
