package router

import (
	"mongoApi/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.AddMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.SetMovieWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteMovies).Methods("DELETE")

	return router
}
