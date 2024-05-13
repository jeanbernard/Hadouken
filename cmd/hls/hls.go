package hls

import (
	"fmt"
	"log"
	"os/exec"
)

type HLS struct {
}

func NewHLS() *HLS {
	return &HLS{}
}

func (h HLS) Create() error {
	fmt.Println("Creating HLS files...")
	cmd, err := exec.Command("/bin/sh", "cmd/hls/hls.sh").Output()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	output := string(cmd)
	fmt.Println(output)

	return nil
}
