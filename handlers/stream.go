package handlers

import (
	"context"
	"log"
	"net/http"
	"rewind/client/googledrive"
	"rewind/cmd/hls"
)

type Stream struct {
}

func NewStream() *Stream {
	return &Stream{}
}

func (s *Stream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	credPath := "client/googledrive/credentials.json"
	hls := hls.NewHLS()

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
}

// func createHLS() error {
// 	cmd, err := exec.Command("/bin/sh", "hls.sh").Output()
// 	if err != nil {
// 		log.Fatalf("Error")
// 	}
// 	output := string(cmd)
// 	fmt.Println(output)

// 	return nil
// }
