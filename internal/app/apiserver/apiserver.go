package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"log"
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
	s.router.HandleFunc("/hello", s.handleHello()).Methods("GET")
	s.router.HandleFunc("/order", s.handleOrder()).Methods("POST")
}

func (s *APIServer) handleOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := make(map[string]string)
		resp["message"] = "You successfully create pizza order"

		respJson, err := json.Marshal(resp)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.WriteHeader(200)
		w.Write(respJson)
		return
	}
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
