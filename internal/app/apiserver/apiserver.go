package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maximys963/pizza-order/models"
	"github.com/maximys963/pizza-order/pkg/config"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("Starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/film", s.addFilm()).Methods("POST")
	s.router.HandleFunc("/films/from-file", s.importFilm()).Methods("POST")
	s.router.HandleFunc("/film/{filmId}", s.deleteFilm()).Methods("DELETE")
	s.router.HandleFunc("/films", s.getFilms()).Methods("GET")
}

func (s *APIServer) addFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var db = config.GetDatabase()

		w.Header().Set("Content-Type", "application/json")
		var film models.Film
		_ = json.NewDecoder(r.Body).Decode(&film)

		result := db.Create(film)

		if result.Error != nil {
			logrus.Error(result.Error)
			return
		}

		marshaledJson, err := json.Marshal(film)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshaledJson)
	}
}

func (s *APIServer) importFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "film imported")
	}
}

func (s *APIServer) deleteFilm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var film models.Film

		vars := mux.Vars(r)
		filmId := vars["filmId"]

		var db = config.GetDatabase()
		result := db.Where("ID=?", filmId).Delete(film)

		if result.Error != nil {
			s.logger.Error(result.Error)
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		marshaledJson, err := json.Marshal(film)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Print(marshaledJson)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshaledJson)
	}
}

// TODO: move it to controller

// TODO: create repository

func (s *APIServer) getFilms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var films []models.Film

		var db = config.GetDatabase()
		result := db.Find(&films)

		if result.Error != nil {
			s.logger.Error(result.Error)
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		marshaledJson, err := json.Marshal(films)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Print(marshaledJson)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(marshaledJson)

	}
}
