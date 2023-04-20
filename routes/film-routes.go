package routes

import (
	"github.com/gorilla/mux"
	"github.com/maximys963/pizza-order/controllers"
)

func RegisterFilmsRoutes(router *mux.Router) {
	router.HandleFunc("/film", controllers.AddFilmHandler).Methods("POST")
	router.HandleFunc("/films/from-file", controllers.ImportFilmHandler).Methods("POST")
	router.HandleFunc("/film/{filmId}", controllers.DeleteFilmHandler).Methods("DELETE")
	router.HandleFunc("/films", controllers.GetFilmsHandler).Methods("GET")
}
