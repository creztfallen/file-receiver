package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.
		Path("/upload").
		Methods("POST").
		HandlerFunc(Upload)
	fmt.Println("Starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}
