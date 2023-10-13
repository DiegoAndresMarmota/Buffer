package main

import (
	"flag"
)

func main() {
	port := flag.String("port", ":8080", "Port to listen")
	flag.Parse()

	network := LogPlataform{}}
	server := NewServer(network, *port)
	server.Run()
}