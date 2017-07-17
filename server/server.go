package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

// https://gist.github.com/peterhellberg/38117e546c217960747aacf689af3dc2
type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}

func New(options ...func(*Server)) *Server {
	s := &Server{mux: http.NewServeMux()}

	for _, f := range options {
		f(s)
	}

	if s.logger == nil {
		s.logger = log.New(os.Stdout, "", 0)
	}

	s.mux.HandleFunc("/", Index)
	s.mux.HandleFunc("/ping", Ping)

	return s
}

func Logger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	s.mux.ServeHTTP(w, r)
	s.logger.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
}
