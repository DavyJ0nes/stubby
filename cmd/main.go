package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/davyj0nes/stubby/config"
	"github.com/davyj0nes/stubby/router"
)

func main() {
	configFile := flag.String("config", "config.yaml", "config file to use")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		panic(err)
	}

	r := router.NewRouter(cfg.Routes)
	addr := fmt.Sprintf(":%d", cfg.Port)

	log.Println("starting server on ", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
