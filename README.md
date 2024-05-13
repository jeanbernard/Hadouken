# Hadouken!!!

A quick, hackathon style app to learn [ffmpeg](https://ffmpeg.org/) for video encoding/decoding and streaming my Street Fighter 6 clips.

## The Idea

Convert my local SF6 clips from H.264 to H.265, upload them to my Google Drive and then stream the clips on my local webserver using HLS. Long term, it'd be nice to stream the clips on my phone.

### Why H.265?

Encoding from H.264 to H.265 will save space on Google Drive; A free account is only 15 GB.

One H.264 30 second clip goes from ~76 MB down to ~25 MB without any noticeable loss in quality.

## Progress so far

**NOTE:** This repo is a WIP! But here's thy progress so far:

- `cmd/encoder` CLI tool to convert H.264 to H.265. Also includes option for uploading to Google Drive

- `cmd/hls` Executes bash script to create HLS segments

- `server/main.go` The web server with two routes:
    - `/` root where the HLS player is
    - `/download` will download the video from Drive and create the HLS segments. Re-directs to root once done

- `client/googledrive` Quick n' dirty client for authenticating to Drive and for uploads/downloads
