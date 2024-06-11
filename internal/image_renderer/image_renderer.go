package image_renderer

import (
	"fmt"
	tm "github.com/buger/goterm"
	"image"
	"image/color"
)

const ascii = "`^\\,:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

type ImageRenderer struct {
	ppi int
}

func NewImageRenderer(ppi int) *ImageRenderer {
	return &ImageRenderer{ppi}
}

func (ir *ImageRenderer) RenderImage(img image.Image) {
	var wFix float64 = 2.0

	sy, sx := float64(img.Bounds().Dy()), float64(img.Bounds().Dx())
	ratio := sx / sy

	aspY, aspX := ratio*float64(int(sy)/ir.ppi), float64(int(sx)/ir.ppi)

	output := ""
	for y := range int(sy / aspY) {
		for x := range int(wFix * sx / aspX) {

			r, g, b, _ := rgba64ToRGBA(img.At(int(aspX*float64(x)/wFix), int(aspY*float64(y))))
			avg := (float64(int(r)+int(g)+int(b)) / 3.0) / 255.0
			char := ascii[int(max(float64(len(ascii))*avg-1.0, 0))]

			ansiCode := rgbToAnsiEscape(r, g, b, true)
			output += fmt.Sprint(ansiCode, string(char), ansiReset())
		}
		output += "\n"
	}
	tm.Println(output)
}

func rgbToAnsiEscape(r, g, b uint8, foreground bool) string {
	if foreground {
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	} else {
		return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
	}
}

func ansiReset() string {
	return "\x1b[0m"
}

func rgba64ToRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	r, g, b, a := c.RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}
