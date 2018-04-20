package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/madsilver/golang-api-example/pkg/handler"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	handler.MockBook()

	// Route Handlers
	r.HandleFunc("/api/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", handler.GetBook).Methods("GET")
	r.HandleFunc("/api/books", handler.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", handler.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", handler.DeleteBook).Methods("DELETE")

	port := 8000
	if portStr := os.Getenv("PORT"); portStr != "" {
		port, _ = strconv.Atoi(portStr)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
