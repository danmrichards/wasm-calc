package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	port, dir string
)

func init() {
	flag.StringVar(&port, "port", "8080", "The port to serve on")
	flag.StringVar(&dir, "dir", ".", "The directory to serve files from")
}

func main() {
	flag.Parse()

	fmt.Println("serving on port", port)

	log.Fatalln(http.ListenAndServe(
		net.JoinHostPort("", port),
		http.FileServer(http.Dir(dir)),
	))
}
