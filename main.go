package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting on http")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the club"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specific snip"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snip"))
}
