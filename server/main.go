package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
)

//go:embed *.png index.html
var site embed.FS

func main() {
	addr := flag.String("addr", ":3333", "address to listen on")
	flag.Parse()

	log.Printf("Serving on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, http.FileServer(http.FS(site))))
}
