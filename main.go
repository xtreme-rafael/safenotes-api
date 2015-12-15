package main

import (
	"github.com/gorilla/mux"
	"github.com/xtreme-rafael/safenotes-api/server"
)

func main() {
	runServer()
}

func runServer() {
	router := mux.NewRouter()
	server := server.NewSafeNotesServer(*router)

	server.RunServer()
}
