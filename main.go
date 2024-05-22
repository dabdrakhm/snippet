package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting on http://localhost:4000\n")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to the club"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snipppet with ID %d...", id)
	// w.Write([]byte("display a specific snip"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Set a new cache-control header. If an existing "Cache-Control" header exists
	// it will be overwritten.
	w.Header().Set("Cache-Control", "public, max-age=31536000")
	// In contrast, the Add() method appends a new "Cache-Control" header and can
	// be called multiple times.
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "max-age=31536000")
	// Delete all values for the "Cache-Control" header.
	w.Header().Del("Cache-Control")
	// Retrieve the first value for the "Cache-Control" header.
	w.Header().Get("Cache-Control")
	// Retrieve a slice of all values for the "Cache-Control" header.
	w.Header().Values("Cache-Control")
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("Create a new snip...\n"))
}
