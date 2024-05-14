package hls

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

type HLS struct {
}

func NewHLS() *HLS {
	return &HLS{}
}

func (h HLS) Create(input string) error {
	fmt.Println("Creating HLS files...")
	cmd := exec.Command("/bin/sh", "cmd/hls/hls.sh", input)

	stdout, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error with pipe")
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Error with start")
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Error with end")
	} else {
		fmt.Println("completed!")
	}

	return nil
}
