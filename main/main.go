package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("../assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/candidates", GetCandidateHandler).Methods("GET")
	return r
}

func main()  {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World!")
}

func handler2(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "List of candidates")
}
