package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "100",
		Title: "Iron Man",
		Director: &Director{
			Firstname: "Anish",
			Lastname:  "Bista",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "101",
		Title: "Spider Man",
		Director: &Director{
			Firstname: "Sahil",
			Lastname:  "Jhadav",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/create", createMovie).Methods("POST")
	r.HandleFunc("/update/{id}", update).Methods("PUT")
	r.HandleFunc("/delete/{id}", delete).Methods("DELETE")

	fmt.Println("Starting the server at port 8000:")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","applicatin/json")
	parameter := mux.Vars(r)
	for _, item := range movies{
		if item.ID == parameter["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}

}
func createMovie(w http.ResponseWriter , r * http.Request) {
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}
func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	parameter := mux.Vars(r)

	for index,item := range movies{
		if item.ID == parameter["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			_= json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
		}
	}

}
func delete(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type","application/json")
	parameter := mux.Vars(r)
	for index, item := range movies{
		if item.ID == parameter["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
