package handlers

import (
	"hadouken/client/googledrive"
	"log"
	"net/http"
)

type Stream struct {
	GoogleDrive *googledrive.GoogleDrive
}

func NewStream(srv *googledrive.GoogleDrive) *Stream {
	return &Stream{srv}
}

func (s *Stream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//hls := hls.NewHLS()

	if err := s.GoogleDrive.Download(); err != nil {
		log.Fatal(err)
		return
	}

	// if err := hls.Create(); err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	http.ServeFile(w, r, "index.html")
}
