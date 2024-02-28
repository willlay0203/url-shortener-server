package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	Port string
}

func (s *Server) Start() error {
	router := httprouter.New()

	// Defining our routes
	router.GET("/healthcheck", healthCheck)

	fmt.Printf("Server listening on %v", s.Port)
	return http.ListenAndServe(s.Port, router)
}

func healthCheck(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "Server is OK")
}
