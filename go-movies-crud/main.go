package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

) 
//https://www.youtube.com/watch?v=jFfo23yIWac
type Movie struct {
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

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID ==params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content.Type","application/json")
	var movie Movie
	_= json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter,r *http.Request){
	//set json content type
	w.Header().Set("Content-Type","application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range
	//delete the movies with the id that youl've sent
	// add a new movie - the movie that we send in the body of postman
	for index, item :=range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index],movies[index + 1:]...)
			var movie Movie
			_ =json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
		
	}
}
func main(){
	r :=mux.NewRouter()
	movies = append(movies,Movie{ID : "1",Isbn : "4030",Title : "Pomplamusse", Director: &Director{Firstname: "luis",Lastname: "Ruiz"}})
	movies = append(movies,Movie{ID : "2",Isbn : "40320",Title : "tesa", Director :&Director {Firstname: "Juan",Lastname: "Romero"}})
	
	 
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}