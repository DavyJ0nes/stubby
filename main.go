package main

import (
	"log"
	"net/http"

	"github.com/davyj0nes/stubby/config"
	"github.com/davyj0nes/stubby/server"
)

func main() {
	cfg, err := config.LoadConfig("config")
	if err != nil {
		panic(err)
	}

	addr := ":" + cfg.Port
	server := server.NewServer(cfg.Routes)

	log.Println("starting server on ", addr)
	log.Fatal(http.ListenAndServe(addr, server))

}
