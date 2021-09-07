package apiserver

import (
	"io"
	"net/http"

	"github.com/echodiv/simple_blog/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer meta structure
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// Create new instance of API server
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start instance of API server
func (s APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("Starting API server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Config logger with parsing config toml file
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureStore() error {
	var st *store.Store = store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

// Add routers to the server
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

// Hello test handler
func (s *APIServer) handleHello() http.HandlerFunc {
	// Code just for this function
	// run once
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}

}
