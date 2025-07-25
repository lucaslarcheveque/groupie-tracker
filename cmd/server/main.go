package main

import (
	"groupie-tracker/internal/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.HomeHandler)

	log.Println("SERVER IS RUNNING ON http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
