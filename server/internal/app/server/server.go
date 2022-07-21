package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/xamust/qtimTestQuiz/internal/app/counter"
	"net/http"
	"sync"
)

type Server struct {
	mu       *sync.Mutex
	config   *Config
	logger   *logrus.Logger
	mux      *mux.Router
	handlers Handlers
	counter  *counter.Counter
}

func NewServer(config *Config) *Server {
	return &Server{
		mu:     new(sync.Mutex),
		config: config,
		logger: logrus.New(),
		mux:    mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	//config logger...
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Logger ready...")

	//config router (gorilla/mux)...
	s.configureRouter()
	s.logger.Info("Router ready...")

	s.configureCounter()

	//handlers init...
	s.handlers = Handlers{
		logger:  s.logger,
		counter: s.counter,
	}

	s.logger.Info(fmt.Sprintf("Starting service (bind on %v)...", s.config.BindAddr))
	//start web server...
	return http.ListenAndServe(s.config.BindAddr, s.mux)
}

//config logger...
func (s *Server) configureLogger() error {
	//get level for logrus from configs...
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	//set level for logrus...
	s.logger.SetLevel(level)
	return nil
}

//config router...
func (s *Server) configureRouter() {
	//register handle on router...
	s.mux.HandleFunc("/detect", s.handlers.Detect)
}

//config counter...
func (s *Server) configureCounter() {
	s.counter = counter.NewCounter(s.config.Counter, s.mu)
}
