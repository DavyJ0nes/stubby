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
	"github.com/gorilla/mux"
)

func main() {
	configFile := flag.String("config", "config.yaml", "config file to use")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
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

	log.Println("stubby is ready to serve...")
	log.Println("routes configured for stubby are:")
	if err := r.Walk(outputRouteInfo); err != nil {
		log.Fatal(err)
	}

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

func outputRouteInfo(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	path, err := route.GetPathTemplate()
	if err != nil {
		return err
	}
	queries, err := route.GetQueriesTemplates()
	if err != nil {
		return err
	}
	log.Println("path: " + path)
	if len(queries) != 0 {
		log.Println("queries:")
		for _, query := range queries {
			log.Println("  - " + query)
		}
	}
	return nil
}
