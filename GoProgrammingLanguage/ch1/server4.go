// handler echos the HTTP request
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var redie = color.RGBA{0x10, 0xff, 0xBB, 0xff}

var palette = []color.Color{color.White, color.Black, redie}

const (
	whiteIndex = 0 // first color
	blackIndex = 1 // next color in palette
	redish     = 2
)

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

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
