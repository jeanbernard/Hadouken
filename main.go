package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"rewind/client/googledrive"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

/*
ffmpeg -i input.mp4 -c:v libx265 -preset medium -x265-params crf=23 -c:a copy output.mp4

ffmpeg -i input.mp4 \
  -c:v libx265 -crf 32 -preset medium -c:a aac -b:a 128k \
  -vf "scale=640:-2" -hls_time 6 -hls_list_size 0 -hls_segment_filename "360p_%03d.ts" 360p.m3u8 \
  -vf "scale=1280:-2" -hls_time 6 -hls_list_size 0 -hls_segment_filename "720p_%03d.ts" 720p.m3u8 \
  -vf "scale=1920:-2" -hls_time 6 -hls_list_size 0 -hls_segment_filename "1080p_%03d.ts" 1080p.m3u8
*/

func main() {
	//ctx := context.Background()
	var crf, video, preset string

	flag.StringVar(&video, "video", "", "H.264 video file")
	flag.StringVar(&preset, "preset", "medium", "low;medium;high - preset quality for H.265 video")
	flag.StringVar(&crf, "crf", "32", "CRF - Constant Rate Factor")

	flag.Parse()

	if video == "" {
		fmt.Println("no video file provided")
		flag.Usage()
		os.Exit(1)
	}

	videoName := strings.Split(video, ".")[0]
	output := fmt.Sprintf("%v_output.mp4", videoName)

	s := fmt.Sprintf("ffmpeg -i /videos/H.264/%v -c:v libx265 -preset %v -x265-params crf=%v -c:a copy videos/H.265/%v", video, preset, crf, output)

	args := strings.Split(s, " ")
	fmt.Println(args[2], args[6], args[8], args[11])
	// cmd := exec.Command(args[0], args[1:]...)

	// stdout, err := cmd.StderrPipe()
	// if err != nil {
	// 	log.Fatalf("Error with pipe")
	// }

	// err = cmd.Start()
	// if err != nil {
	// 	log.Fatalf("Error with start")
	// }

	// scanner := bufio.NewScanner(stdout)
	// scanner.Split(bufio.ScanLines)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// err = cmd.Wait()
	// if err != nil {
	// 	log.Fatalf("Error with end")
	// } else {
	// 	fmt.Println("completed!")
	// }

	fmt.Println("Uploading to Google...")

	// upload to Google Drive
	// if err := upload(ctx); err != nil {
	// 	log.Fatalf(err.Error())
	// }
}

func upload(ctx context.Context) error {
	file, err := os.ReadFile("client/googledrive/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := googledrive.GetConfig(file)
	if err != nil {
		log.Fatalf("Unable to get config: %v", err)
	}

	client, err := googledrive.GetClient(ctx, config)
	if err != nil {
		log.Fatalf("Unable to get client: %v", err)
	}

	srv, err := drive.NewService(ctx, option.WithHTTPClient(client), option.WithScopes("drive.DriveScope"))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}

	// Open the video file
	video, err := os.Open("videos/H.265/output.mp4")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer video.Close()

	f := &drive.File{Name: "SF6Yay"}

	resp, err := srv.Files.Create(f).Media(video).ProgressUpdater(func(now, size int64) {
		fmt.Printf("%d, %d\r", now, size)
	}).Do()

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("new file id: %s\n", resp.Id)

	return nil
}
