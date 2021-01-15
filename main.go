package main

import (
	h "course/http"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/", h.Index)
	r.HandleFunc("/order", h.Order)
	r.HandleFunc("/add_guest", h.AddGuest)
	r.HandleFunc("/contract", h.Contract)
	r.HandleFunc("/guest", h.Guest)

	r.PathPrefix("/src/css").Handler(http.StripPrefix("/src/css",
		http.FileServer(http.Dir("templates/src/css"))))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static",
		http.FileServer(http.Dir("templates/static"))))
	r.PathPrefix("/src/scripts").Handler(http.StripPrefix("/src/scripts",
		http.FileServer(http.Dir("templates/src/scripts"))))
	r.PathPrefix("/templates").Handler(http.StripPrefix("/templates",
		http.FileServer(http.Dir("templates"))))

	http.Handle("/", r)
	http.ListenAndServe(":8181", nil)
}
