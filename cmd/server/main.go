package main

import (
	"groupie-tracker/internal/handler"
	"log"
	"net/http"
)

func main() {
	// Create a custom ServeMux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handler.ErrorHandler(w, 404, "Page not found", "Sorry, the page you requested does not exist.")
			return
		}
		handler.HomeHandler(w, r)

	})
	mux.HandleFunc("/artist/", handler.ArtistHandler) // trailing slash to match /artist/<id>
	mux.Handle("/web/assets/", http.StripPrefix("/web/assets/", http.FileServer(http.Dir("web/assets"))))

	// Wrap mux to handle 404 dynamically
	wrappedMux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, pattern := mux.Handler(r); pattern != "" {
			// Route exists → serve normally
			mux.ServeHTTP(w, r)
		} else {
			// Route not found → use your custom 404 handler
			handler.ErrorHandler(w, 404, "Page not found", "Sorry, the page you requested does not exist.")
		}
	})

	// Start the server
	log.Println("SERVER IS RUNNING ON http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
