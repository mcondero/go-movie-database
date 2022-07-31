package main

import (
	"fmt"
	"log"
	"net/http"

	"gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "11273113353424306320161402164463", Title: "Pulp Fiction", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}})
	movies = append(movies, Movie{ID: "2", Isbn: "14532562642216133024705075440256", Title: "From Dusk Till Dawn", Director: &Director{Firstname: "Robert", Lastname: "Rodriguez"}})
	movies = append(movies, Movie{ID: "3", Isbn: "23273350013315342503532172521613", Title: "Bram Stoker's Dracula", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{ID: "4", Isbn: "13213773162001332267237751666266", Title: "Big Trouble in Little China", Director: &Director{Firstname: "John", Lastname: "Carpenter"}})
	movies = append(movies, Movie{ID: "5", Isbn: "40525572443777533431762277521414", Title: "Alien", Director: &Director{Firstname: "Riddley", Lastname: "Scott"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
