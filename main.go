package main

import (
	"flag"
	"log"

	"github.com/raisa320/API/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":9000", "The server listenAddr")
	flag.Parse()

	server := api.NewServer(*listenAddr)

	log.Fatal(server.Start())

}
