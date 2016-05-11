package main

import (
	"flag"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	inPath := flag.String("in", "-", "input file, or - for stdin")
	outPath := flag.String("out", "-", "output file, or - for stdout")
	width := flag.Int("width", 0, "maximum width")
	height := flag.Int("height", 0, "maximum height")

	flag.Parse()

	var input = os.Stdin
	if *inPath != "-" {
		f, err := os.Open(*inPath)
		if err != nil {
			log.Fatal(err)
		}
		input = f
	}

	img, format, err := image.Decode(input)
	if err != nil {
		log.Fatal(err)
	}

	input.Close()

	log.Printf("Resizing a %s to maximum %v x %v", format, *width, *height)

	resized := resize.Thumbnail(uint(*width), uint(*height), img, resize.Bicubic)

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
