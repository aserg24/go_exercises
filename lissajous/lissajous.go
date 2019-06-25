// Package lissajous generates GIF animations of random Lissajous figures with given parameters. Exercises 1.5, 1.12
package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.Black, color.RGBA{G: 255, R: 0, B: 0, A: 100}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func Lissajous(out io.Writer, parseCycles int, parseRes float64, parseSize, parseNframes, parseDelay int) {
	cycles, res, size, nframes, delay := getParameters(parseCycles, parseRes, parseSize, parseNframes, parseDelay)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func getParameters(parseCycles int, parseRes float64, parseSize int, parseNframes int, parseDelay int,
) (int, float64, int, int, int) {
	var (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if parseCycles != 0 {
		cycles = parseCycles
	}
	if parseRes != 0.0 {
		res = parseRes
	}
	if parseSize != 0 {
		size = parseSize
	}
	if parseNframes != 0 {
		nframes = parseNframes
	}
	if parseDelay != 0 {
		delay = parseDelay
	}
	return cycles, res, size, nframes, delay
}
