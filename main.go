package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func main() {
	// create a new router
	r := mux.NewRouter()

	// tell the router what to do with "/assests/" path
	fs := http.FileServer(http.Dir("assests"))
	r.PathPrefix("/assests").Handler(http.StripPrefix("/assests", fs))

	// create routers for pages
	homeRouter := r.Methods(http.MethodGet).Subrouter()
	homeRouter.HandleFunc("/", ServeHome)

	// create a new server
	s := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	fmt.Println("Running server on localhost:", s.Addr)
	s.ListenAndServe()
}
