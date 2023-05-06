package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raisa320/API/api/handlers"
)

type Server struct {
	listenAddr string
	handler    *mux.Router
}

func NewServer(listenAddr string) *Server {
	handler := mux.NewRouter()
	var handlersItem handlers.ItemHandler
	handler.HandleFunc("/items", handlersItem.GetItems).Methods("GET")
	handler.HandleFunc("/items", handlersItem.SaveItem).Methods("POST")

	return &Server{
		listenAddr: listenAddr,
		handler:    handler,
	}
}

func (s *Server) Start() error {
	fmt.Println("Server running on port", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.handler)
}
