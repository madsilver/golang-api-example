package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/madsilver/golang-api-example/pkg/routes"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	routes.MockBook()

	// Route Handlers
	r.Handle("/api/books", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(routes.GetBooks))).Methods("GET")
	r.Handle("/api/books/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(routes.GetBook))).Methods("GET")
	r.Handle("/api/books", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(routes.CreateBook))).Methods("POST")
	r.Handle("/api/books/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(routes.UpdateBook))).Methods("PUT")
	r.Handle("/api/books/{id}", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(routes.DeleteBook))).Methods("DELETE")

	port := ":8000"

	log.Fatal(http.ListenAndServe(port, handlers.CORS()(r)))
}
