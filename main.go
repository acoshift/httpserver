package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/moonrhythm/webstatic/v4"
)

var (
	addr = flag.String("addr", ":8080", "listen address")
)

func main() {
	flag.Parse()

	h := logger(webstatic.Dir("."))

	log.Println("Start http server on", *addr)
	err := http.ListenAndServe(*addr, h)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
