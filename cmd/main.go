package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	srv := http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	log.Println("starting stubby on ", addr)
	go func() {
		log.Fatal(srv.ListenAndServe())

	}()

	log.Print("stubby is ready to serve...")

	killSignal := <-interrupt
	switch killSignal {
	case os.Interrupt:
		log.Println("got SIGINT...")
		log.Println("stubby is shutting down...")
	case syscall.SIGTERM:
		log.Println("got SIGTERM...")
		log.Println("stubby is shutting down...")
	}

	log.Fatal(srv.Shutdown(context.Background()))
}
