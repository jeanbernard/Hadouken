#!/bin/bash

ffmpeg -hide_banner -i cmd/hls/$1 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 -map 0:v:0 -map 0:a:0 \
-c:v libx265 -profile:v main -crf 32 -sc_threshold 0 -g 48 -keyint_min 48 -c:a aac -ar 48000 \
-filter:v:0 scale=w=640:h=360:force_original_aspect_ratio=decrease -maxrate:v:0 856k -bufsize:v:0 1200k -preset fast -b:a:0 96k \
-filter:v:1 scale=w=640:h=480:force_original_aspect_ratio=decrease -maxrate:v:1 1600k -bufsize:v:1 2100k -preset fast -b:a:1 128k \
-filter:v:2 scale=w=1280:h=720:force_original_aspect_ratio=decrease -maxrate:v:2 4000k -bufsize:v:2 4200k -preset fast -b:a:2 128k \
-filter:v:3 scale=w=1920:h=1080:force_original_aspect_ratio=decrease -maxrate:v:3 6000k -bufsize:v:3 7500k -preset fast -b:a:3 192k \
-var_stream_map "v:0,a:0 v:1,a:1 v:2,a:2 v:3,a:3" \
-f hls -hls_time 6 -hls_playlist_type vod -hls_flags independent_segments \
-master_pl_name master.m3u8 \
-hls_segment_filename output/%v_%03d.ts output/%v.m3u8
