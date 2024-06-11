package main

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/marcsek/ascii_video/internal/image_renderer"
	"github.com/marcsek/ascii_video/internal/video_decoder"
)

func main() {
	vd, err := video_decoder.NewVideoDecoder("/home/marek/Pictures/fire.mp4")

	if err != nil {
		fmt.Println(err)
	}

	images := vd.DecodeVideo()
	ir := image_renderer.NewImageRenderer(50)

	dTarget := (1.0 / float64(vd.Fps)) * 1_000_000.0
	delta := 0.0
	frameIdx := 0
	startTime := time.Now()

	for frameIdx < len(images) {
		if delta < dTarget {
			delta += float64(time.Since(startTime).Microseconds())
			startTime = time.Now()
			time.Sleep(1 * time.Millisecond)
			continue
		}

		tm.MoveCursor(1, 1)
		delta = 0.0
		startTime = time.Now()
		ir.RenderImage(images[frameIdx])
		frameIdx += 1
		tm.Flush()
	}
}
