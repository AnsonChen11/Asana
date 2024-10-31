package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go server with Gorilla Mux is running in docker through CICD in dev environment!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloHandler).Methods("GET")
	fmt.Println("Server started at :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
