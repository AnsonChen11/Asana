package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go server with Gorilla Mux!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods("GET")
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
