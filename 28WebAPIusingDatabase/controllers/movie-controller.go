package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/sahib139/go_Basic_API_using_mongo/models"
	repository "github.com/sahib139/go_Basic_API_using_mongo/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Greater(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"msg":"hi"}`)
}

func CreateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.MovieId = primitive.NewObjectID()
	repository.InsertOne(movie)
	json.NewEncoder(w).Encode(movie)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie := repository.GetById(params["id"])
	json.NewEncoder(w).Encode(movie)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := repository.GetAll()
	json.NewEncoder(w).Encode(movies)
}

func UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie model.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	result := repository.UpdateById(params["id"], movie)
	json.NewEncoder(w).Encode(result)
}

func DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := repository.DeleteById(params["id"])
	json.NewEncoder(w).Encode(result)

}

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result := repository.DeleteAll()
	json.NewEncoder(w).Encode(result)
}

func IsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := repository.MarkAsWatched(params["id"])
	json.NewEncoder(w).Encode(result)
}
