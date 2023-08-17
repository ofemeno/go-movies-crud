package main

import (
	"github.com/gorilla/mux"
)

// struct type for Movie
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"idbn"`
	title    string    `json:"title"`
	Director *Director `json:"director"`
}

// struct types for Director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"firstname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// handle toutes
	r.HandleFunc("/movies", getMovies)
	r.HandleFunc("/movies/{id}", getMovie)
	r.HandleFunc("/movies", createMovie)
	r.HandleFunc("/movies/{id}", updateMovie)
	r.HandleFunc("/movies/{id}", deleteMovie)
}
