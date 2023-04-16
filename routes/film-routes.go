package routes

import (
	"github.com/gorilla/mux"
	"github.com/maximys963/pizza-order/controllers"
)

func RegisterFilmsRoutes(router *mux.Router) {
	router.HandleFunc("/film", controllers.AddFilm).Methods("POST")
	router.HandleFunc("/films/from-file", controllers.ImportFilm).Methods("POST")
	router.HandleFunc("/film/{filmId}", controllers.DeleteFilm).Methods("DELETE")
	router.HandleFunc("/films", controllers.GetFilms).Methods("GET")
}
