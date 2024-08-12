package main

import(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

// main
func main(){
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	movies = append(movies, Movie{ID: "1", Isbn: "448743", Title: "Kuch Kuch Hota Hai", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "847564", Title: "Janeman Janejaha", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fwt.Println("Server starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}