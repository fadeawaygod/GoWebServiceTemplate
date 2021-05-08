package main

import (
	"log"
	"net/http"

	requestHandler "./handler"
	"github.com/rs/cors"
)

func main() {
	http.HandleFunc("/", requestHandler.SayHelloWorld)
	mux := http.NewServeMux()
	mux.HandleFunc("/", requestHandler.SayHelloWorld)
	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":80", handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
