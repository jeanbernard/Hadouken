package handlers

import (
	"context"
	"hadouken/client/googledrive"
	"hadouken/cmd/hls"
	"log"
	"net/http"
)

type Stream struct {
}

func NewStream() *Stream {
	return &Stream{}
}

func (s *Stream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	hls := hls.NewHLS()
	credPath := "client/googledrive/credentials.json"

	drive, err := googledrive.NewDriveService(ctx, credPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := drive.Download(); err != nil {
		log.Fatal(err)
		return
	}

	if err = hls.Create(); err != nil {
		log.Fatal(err)
		return
	}

	http.ServeFile(w, r, "index.html")
}
