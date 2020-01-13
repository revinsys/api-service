package apiserver

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"{GIT_PATH}/internal/app/store"
)

// APIServer ...
type APIServer struct {
	config               *ServerConfig
	logger               *logrus.Logger
	store                *store.Store
	router               *mux.Router
}

// NewAPIServer ...
func NewAPIServer() *APIServer {
	return &APIServer{
		config: NewServerConfig(),
		logger: logrus.New(),
		store:  store.NewStore(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.configureStore(); err != nil {
		return err
	}

	s.configureRoutes()
	s.logger.Info("Server started on %s", s.config.BindAddr)

	handler := cors.AllowAll().Handler(s.router)

	return http.ListenAndServe(s.config.BindAddr, handler)
}

func (s *APIServer) configureStore() error {
	st := store.NewStore()
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRoutes() {

}

// error ...
func (s *APIServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

// respond ...
func (s *APIServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

