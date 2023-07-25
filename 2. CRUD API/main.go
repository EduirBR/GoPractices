package main

import (
	"crud/movies"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", movies.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", movies.GetMovie).Methods("Get")
	router.HandleFunc("/movies", movies.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", movies.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", movies.DeleteMovies).Methods("DELETE")

	port := ":8000"
	fmt.Printf("Start Server in port%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
