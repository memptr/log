package main

import (
	"github.com/memptr/log/internal/server"
	"log"
)

func main() {
	s := server.NewHTTPServer(":6969")
	log.Fatal(s.ListenAndServe())
}
