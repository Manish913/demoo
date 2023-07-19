package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:"name"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

var movie []Movie

func main() {
	m := mux.NewRouter()
	movie = append(movie, Movie{"1", "Bahubali", &Director{Fname: "manish", Lname: "kumar"}})
	movie = append(movie, Movie{"2", "shershah", &Director{Fname: "manish shah", Lname: "sighaniya"}})
	movie = append(movie, Movie{"3", "LoVE", &Director{Fname: "DEV", Lname: "kumar"}})
	m.HandleFunc("/movie", getMovies).Methods("GET")
	m.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	m.HandleFunc("/movie/{id}", saveMovie).Methods("POST")
	m.HandleFunc("/movie/{id}", update).Methods("PUT")
	m.HandleFunc("/movie/{id}", remove).Methods("DELETE")
	log.Println("Server listing on 8020")
	log.Println(http.ListenAndServe(":8020", m))
}

func update(writer http.ResponseWriter, request *http.Request) {

}

func remove(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "Application/json")
	param := mux.Vars(request)
	for index, h := range movie {
		if h.ID == param["Id"] {
			movie = append(movie[:index], movie[index+1:]...)
			break
		}

	}
	json.NewEncoder(writer).Encode(movie)
}

func saveMovie(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(writer)
	//writer.Write([]byte)

}
func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("content-type", "Application/json")
	param := mux.Vars(request)
	for _, h := range movie {
		if h.ID == param["Id"] {
			json.NewEncoder(writer).Encode(movie)
			return
		}

	}
	//it returns remaining items
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("key-content", "Application/json")
	json.NewEncoder(writer).Encode(movie)
}
