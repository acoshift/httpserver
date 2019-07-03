package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/acoshift/webstatic"
)

var (
	addr = flag.String("addr", ":8080", "listen address")
)

func main() {
	flag.Parse()

	h := webstatic.NewDir(".")
	h = logger(h)

	log.Println("Start http server on", *addr)
	err := http.ListenAndServe(*addr, h)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
