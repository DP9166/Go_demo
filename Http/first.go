package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", r))
}