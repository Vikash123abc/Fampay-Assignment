package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from go"))
}
func search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world from go"))
}
func main() {

	// Connect to Database
	http.HandleFunc("/hello", hello)
	r := mux.NewRouter()

	//r.HandleFunc("/hello", hello).Methods("GET")
	r.HandleFunc("/search/searchString", search).Methods("GET")
	http.ListenAndServe(":1000", nil)
}
