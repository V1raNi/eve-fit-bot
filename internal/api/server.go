package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	router  *mux.Router
	httpSrv *http.Server
}

func NewServer(port string) *Server {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "status ok")
	})

	s := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	return &Server{
		router:  r,
		httpSrv: s,
	}
}

func (s *Server) setProjectRoutes() {

}

func (s *Server) Run() error {
	s.setProjectRoutes()

	return s.httpSrv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}
