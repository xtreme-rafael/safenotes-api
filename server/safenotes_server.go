package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	DefaultPort = "5000"
)

func NewSafeNotesServer(router mux.Router) ServerRunner {
	return &SafeNotesServer{Router: router}
}

type SafeNotesServer struct {
	Router mux.Router
}

func (self *SafeNotesServer) RunServer() error {
	self.Router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	log.Println("Listening on port", DefaultPort)
	return HttpListenAndServe(fmt.Sprintf(":%s", DefaultPort), &self.Router)
}
