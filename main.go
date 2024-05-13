package main

import (
	"fmt"
	"log"
	"net/http"
	"rewind/handlers"
)

func main() {
	stream := handlers.NewStream()

	http.Handle("/", addHeaders(http.FileServer(http.Dir("."))))
	http.Handle("/download", stream)

	fmt.Printf("Starting server on %v\n", 8080)
	log.Printf("Serving %s on HTTP port: %v\n", ".", 8080)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8080), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
