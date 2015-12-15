package main

import (
	"net/http"
)

func main() {
	runServer()
}

func runServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
