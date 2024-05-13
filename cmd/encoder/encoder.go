package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"hadouken/client/googledrive"
	"log"
	"os"
	"os/exec"
	"strings"
)

type encoder struct {
	preset string
	output string
	CRF    string
}

func main() {
	e := &encoder{}
	ctx := context.Background()
	credPath := "client/googledrive/credentials.json"

	drive, err := googledrive.NewDriveService(ctx, credPath)
	if err != nil {
		return
	}

	flag.StringVar(&e.preset, "preset", "medium", "low;medium;high - preset quality for H.265 video")
	flag.StringVar(&e.CRF, "crf", "32", "CRF - Constant Rate Factor")
	flag.StringVar(&e.output, "output", "output.mp4", "output for uploading to Drive")

	// encode video
	flag.CommandLine.Func("encode", "input video to encode", func(input string) error {
		if err := e.encode(input); err != nil {
			return err
		}
		return nil
	})

	// upload to Google Drive
	flag.CommandLine.Func("upload", "upload to Google Drive", func(filename string) error {
		if err := drive.Upload(ctx, filename); err != nil {
			log.Fatalf(err.Error())
		}
		return nil
	})

	flag.Parse()
}

func (e encoder) encode(input string) error {
	input = strings.Trim(input, " ")

	if input == "" {
		fmt.Println("no video file provided")
		flag.Usage()
		os.Exit(1)
	}

	s := fmt.Sprintf("ffmpeg -i videos/H.264/%v -c:v libx265 -preset %v -x265-params crf=%v -c:a copy videos/H.265/%v",
		input, e.preset, e.CRF, e.output)

	args := strings.Split(s, " ")
	cmd := exec.Command(args[0], args[1:]...)

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
