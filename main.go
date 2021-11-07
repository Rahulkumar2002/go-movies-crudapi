package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json :"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json : "director"`
}

type Director struct {
	Firstname string `json : "firstname"`
	Lastname  string `json : "lastname"`
}

var movie []Movie

func main() {
	fmt.Println("Your server is connected on port 8000")
	r := mux.NewRouter()
	r.HandlerFunc("/movies", getMovies).Method("GET")
	r.HandlerFunc("/movies/{id}", getMovie).Method("GET")
	r.HandlerFunc("/movies", createMovie).Method("POST")
	r.HandlerFunc("/movies/{id}", updateMovie).Method("PUT")
	r.HandlerFunc("/movies/{id}", deleteMovie).Method("DELETE")

	fmt.Fprintln("Starting your port at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
