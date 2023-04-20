package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/maximys963/pizza-order/models"
	"net/http"
)

func AddFilmHandler(w http.ResponseWriter, r *http.Request) {
	var film models.Film
	_ = json.NewDecoder(r.Body).Decode(&film)

	models.AddFilm(film)

	marshaledJson, err := json.Marshal(film)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledJson)

}

func ImportFilmHandler(w http.ResponseWriter, r *http.Request) {}

func DeleteFilmHandler(w http.ResponseWriter, r *http.Request) {
	var film models.Film

	vars := mux.Vars(r)
	filmId := vars["filmId"]

	film = models.DeleteFilm(filmId)

	marshaledJson, err := json.Marshal(film)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(marshaledJson)
}

func GetFilmsHandler(w http.ResponseWriter, r *http.Request) {
	var films = models.GetAllFilms()

	marshaledJson, err := json.Marshal(films)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(marshaledJson)
}
