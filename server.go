//go:build !js && !wasm

package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))
	log.Printf("Serving %s on http://localhost%s", *dir, *listen)
	log.Fatal(http.ListenAndServe(*listen, fs))

}
