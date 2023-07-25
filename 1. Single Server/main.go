package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(rw, "404 not Found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(rw, "method not allowed", http.StatusBadGateway)
	}
	fmt.Fprintf(rw, "hello!")
}

func formhandler(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(rw, " POST Request Successfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(rw, "Name = %s\n", name)
	fmt.Fprintf(rw, "Address = %s\n", address)

}
func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)
	fmt.Printf("Starting Server at port ")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
