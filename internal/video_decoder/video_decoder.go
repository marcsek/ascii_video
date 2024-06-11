package video_decoder

import (
	"fmt"
	"image"

	"github.com/AlexEidt/Vidio"
)

type VideoDecoder struct {
	Fps    float64
	Width  int
	Height int
	Frames int
	vidio  *vidio.Video
}

func NewVideoDecoder(filepath string) (*VideoDecoder, error) {
	vid, err := vidio.NewVideo(filepath)

	if err != nil {
		return nil, fmt.Errorf("There was an error while trying to load the video (%s)\n", err)
	}

	fmt.Printf("FPS: %f, Total Frames: %d\n", vid.FPS(), vid.Frames())

	return &VideoDecoder{
		Fps:    vid.FPS(),
		Width:  vid.Width(),
		Height: vid.Height(),
		Frames: vid.Frames(),
		vidio:  vid,
	}, nil
}

func (vd *VideoDecoder) DecodeVideo() []*image.RGBA {
	images := make([]*image.RGBA, 0, vd.Frames)
	for vd.vidio.Read() {
		frame := vd.vidio.FrameBuffer()

		img := image.NewRGBA(image.Rect(0, 0, vd.Width, vd.Height))

		copy(img.Pix, frame)
		images = append(images, img)
	}

	return images
}
