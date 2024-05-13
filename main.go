package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rewind/client/googledrive"
)

func main() {
	ctx := context.Background()
	credPath := "client/googledrive/credentials.json"

	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello!")
		drive, err := googledrive.NewDriveService(ctx, credPath)
		if err != nil {
			log.Fatal(err)
			return
		}
		if err := drive.Download(); err != nil {
			log.Fatal(err)
			return
		}

	})
	fmt.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
