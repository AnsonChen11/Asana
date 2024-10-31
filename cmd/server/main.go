package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go server with Gorilla Mux is running!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods("GET")
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
