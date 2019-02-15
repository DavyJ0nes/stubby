package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/davyj0nes/stubby/config"
	"github.com/davyj0nes/stubby/server"
)

func main() {
	configFile := flag.String("config", "config.yaml", "config file to use")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(cfg.Routes)
	addr := ":8080"

	log.Println("starting server on ", addr)
	log.Fatal(http.ListenAndServe(addr, srv))
}
