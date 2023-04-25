package routes

import (
	"github.com/gorilla/mux"
	"github.com/maximys963/pizza-order/controllers"
)

func RegisterFilmsRoutes(router *mux.Router) {
	router.HandleFunc("/films", controllers.GetFilmsHandler).Methods("GET")
	router.HandleFunc("/film", controllers.AddFilmHandler).Methods("POST")
	router.HandleFunc("/film/{filmId}", controllers.GetFilmHandler).Methods("GET")
	router.HandleFunc("/film/{filmId}", controllers.UpdateFilmHandler).Methods("PUT")
	router.HandleFunc("/film/{filmId}", controllers.DeleteFilmHandler).Methods("DELETE")
	router.HandleFunc("/films/from-file", controllers.ImportFilmHandler).Methods("POST")
}
