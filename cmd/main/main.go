package main

import (
	"flag"
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/marcsek/ascii_video/internal/image_renderer"
	"github.com/marcsek/ascii_video/internal/video_decoder"
)

func main() {
	path := flag.String("s", "", "Source of video.")
	loop := flag.Bool("l", false, "Loop video.")

	flag.Parse()

	if *path == "" {
		fmt.Println("'-s' flag is required (video source)")
		return
	}

	play_video(*path, *loop)
}

func play_video(filepath string, loop bool) {
	vd, err := video_decoder.NewVideoDecoder(filepath)

	if err != nil {
		fmt.Println(err)
	}

	frames := vd.DecodeVideo()
	ir := image_renderer.NewImageRenderer(50)

	dTarget := (1.0 / float64(vd.Fps)) * 1_000_000.0
	delta := 0.0
	frameIdx := 0
	startTime := time.Now()

	for frameIdx < len(frames) {
		if delta < dTarget {
			delta += float64(time.Since(startTime).Microseconds())
			startTime = time.Now()
			time.Sleep(1 * time.Millisecond)
			continue
		}

		tm.MoveCursor(1, 1)
		delta = 0.0
		startTime = time.Now()
		ir.RenderImage(frames[frameIdx])
		frameIdx += 1
		tm.Flush()

		if loop && frameIdx == len(frames) {
			frameIdx = 0
		}
	}
}
