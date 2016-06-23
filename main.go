package main

import (
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"

	"github.com/nfnt/resize"
)

func main() {
	inPath := flag.String("in", "-", "input file, url, or - for stdin")
	outPath := flag.String("out", "-", "output file, or - for stdout")
	method := flag.String("method", "nearest", "image resize method - 'nearest', 'bicubic', 'bilinear', 'mitchell', 'lanczos2', or 'lanzoc3'")
	width := flag.Int("width", 0, "maximum width")
  height := flag.Int("height", 0, "maximum height")

	flag.Parse()

	var input = io.Reader(os.Stdin)
	if *inPath != "-" {
		parsedURL, err := url.Parse(*inPath)
		if err != nil {
			log.Fatal(err)
		} else {
			switch parsedURL.Scheme {
			case "http", "https":
				response, err := http.Get(parsedURL.String())
				if err != nil {
					log.Fatal(err)
				}
				input = response.Body
			case "file":
				inpath := filepath.FromSlash(filepath.Clean(parsedURL.Path))
				if runtime.GOOS == "windows" {
					// A file:/ URL on windows will usually look like file:/C:/foo/bar/baz.png
					// The Path part of that URL ends up as /C:/foo/bar/baz.png
					// This block strips off the leading slash
					r, _ := regexp.Compile("\\\\[a-zA-Z]:\\\\")
					if r.MatchString(inpath) {
						inpath = inpath[1:len(inpath)]
					}
				}
				f, err := os.Open(inpath)
				if err != nil {
					log.Fatal(err)
				}
				input = f
				defer f.Close()
			case "":
				f, err := os.Open(*inPath)
				if err != nil {
					log.Fatal(err)
				}
				input = f
				defer f.Close()
			default:
				log.Fatalf("Unsupported URL scheme %s", parsedURL.Scheme)
			}
		}

	}

	img, format, err := image.Decode(input)
	if err != nil {
		log.Fatal(err)
	}

  methods := make(map[string]resize.InterpolationFunction)
  methods["nearest"] = resize.NearestNeighbor
  methods["bicubic"] = resize.Bicubic
  methods["bilinear"] = resize.Bilinear
  methods["mitchell"] = resize.MitchellNetravali
  methods["lanczos2"] = resize.Lanczos2
  methods["lanczos3"] = resize.Lanczos3

	log.Printf("Resizing a %s to maximum %v x %v", format, *width, *height)

	resized := resize.Thumbnail(uint(*width), uint(*height), img, methods[*method])

	var output = os.Stdout
	if *outPath != "-" {
		f, err := os.Create(*outPath)
		if err != nil {
			log.Fatal(err)
		}
		output = f
		defer f.Close()
	}

	switch format {
	case "jpeg":
		jpeg.Encode(output, resized, nil) // TODO set jpeg quality
	case "png":
		png.Encode(output, resized)
	case "gif":
		gif.Encode(output, resized, nil) // TODO set gif options
	}
}
