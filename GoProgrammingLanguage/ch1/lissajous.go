// Lissajous generate GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var redie = color.RGBA{0x10, 0xff, 0xBB, 0xff}

var palette = []color.Color{color.White, color.Black, redie}

const (
	whiteIndex = 0 // first color
	blackIndex = 1 // next color in palette
	redish     = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revs
		res     = 0.001 // angular resolution
		size    = 100   // image canvas coverfs [-size..+size]
		nframes = 64    // animation frames
		delay   = 8     // delay between frames in 10m units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscialltor
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			var index uint8
			if rand.Float64() > 0.5 {
				index = blackIndex
			} else {
				index = redish
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // note: ignoring encoding errors.
}
