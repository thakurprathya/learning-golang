package controllers

import (
	"encoding/json"
	"mongoApi/middlewares"
	"mongoApi/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/x-www-form-urlencode")
	movies := middlewares.GetMovies()
	json.NewEncoder(w).Encode(movies)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") // what type of values and methods are allowed

	var movie models.WatchMovie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	result := middlewares.AddMovie(movie)
	movie.ID = result.(primitive.ObjectID)
	json.NewEncoder(w).Encode(movie)
}

func SetMovieWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	middlewares.UpdateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	middlewares.DeleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applcation/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	middlewares.DeleteMovies()
	json.NewEncoder(w).Encode("All the movies deleted")
}
