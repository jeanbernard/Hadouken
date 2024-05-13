package main

import (
	"context"
	"fmt"
	"hadouken/client/googledrive"
	"hadouken/handlers"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	googleDrive, err := googledrive.NewDriveService(ctx, "client/googledrive/credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	streamHandler := handlers.NewStream(googleDrive)

	http.Handle("/", addHeaders(http.FileServer(http.Dir("."))))
	http.Handle("/download", streamHandler)

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
