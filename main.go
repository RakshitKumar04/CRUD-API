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

// Get All Movies
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Get Single Movie
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	// Loop through movies and find with id
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Movie{})
}

// Create a new Movie
func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "applciation/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // Mock ID
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

// Update Movie
func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// Delete Movie
func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	// Loop through movies and find with id
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

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

	fmt.Println("Server starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
