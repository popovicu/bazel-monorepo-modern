// gallery writes an HTML page showing the PNG files given as arguments.
package main

import (
	"flag"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/popovicu/bazel-monorepo-modern/tools/gallery"
)

func main() {
	out := flag.String("out", "", "path of the HTML file to write")
	scale := flag.Int("scale", 2, "pixels per CSS pixel in the images")
	flag.Parse()

	var images []gallery.Image
	for _, path := range flag.Args() {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		cfg, err := png.DecodeConfig(f)
		f.Close()
		if err != nil {
			log.Fatalf("%s: %v", path, err)
		}
		images = append(images, gallery.Image{
			Name:   filepath.Base(path),
			Width:  cfg.Width / *scale,
			Height: cfg.Height / *scale,
		})
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}
	if err := gallery.Render(f, images); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
