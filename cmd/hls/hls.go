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
	cmd, err := exec.Command("/bin/sh", "hls.sh").Output()
	if err != nil {
		log.Fatalf("Error")
	}
	output := string(cmd)
	fmt.Println(output)

	return nil
}
