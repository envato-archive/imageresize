# Envato Sites Image Resizer (`imageresize`)

Uses https://github.com/nfnt/resize to scale images (JPG, PNG, or GIF) down to
fit within a bounding box with a given width and height. If the input image is
smaller than the bounding box, it is output unmodified.

```sh
go get https://github.com/envato-sites/imageresize
```

## Usage

```
$ imageresize -h
Usage of imageresize:
  -height int
        maximum height
  -in string
        input file, url, or - for stdin (default "-")
  -out string
        output file, or - for stdout (default "-")
  -width int
        maximum width
```

Input and output default to `stdin` and `stdout`, allowing you to pipe image
data to/from this utility, or paths may be specified using the `-in` and `-out`
flags:

```sh
imageresize -in cats.jpg -height 100 -width 100 > cats-thumb.jpg
imageresize -in logo.png -height 300 -width 200 -out logo-resized.png
cat dogs.gif | imageresize -width 50 -height 100 > dogs-small.gif
imageresize -in http://placehold.it/400x400 -width 50 -height 100 > place.png
```
