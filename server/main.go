package main

import (
	"fmt"
	"hadouken/handlers"
	"log"
	"net/http"
)

func main() {
	stream := handlers.NewStream()

	http.Handle("/", addHeaders(http.FileServer(http.Dir("."))))
	http.Handle("/download", stream)

	log.Printf("Starting server on %v\n", 8080)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8080), nil))
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
