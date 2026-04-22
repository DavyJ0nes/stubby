package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/davyj0nes/stubby/internal/config"
	"github.com/davyj0nes/stubby/internal/router"
	"github.com/gorilla/mux"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	configFile := flag.String("config", "config.yaml", "config file to use")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	r := router.NewRouter(cfg.Routes)
	addr := fmt.Sprintf(":%d", cfg.Port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	slog.Info("starting stubby", "addr", addr)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen failed", "error", err)
			os.Exit(1)
		}
	}()

	slog.Info("stubby is ready to serve")
	if err := r.Walk(outputRouteInfo); err != nil {
		slog.Error("failed to walk routes", "error", err)
		os.Exit(1)
	}

	killSignal := <-interrupt
	slog.Info("shutting down", "signal", killSignal.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("graceful shutdown failed", "error", err)
	}
}

func outputRouteInfo(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
	path, err := route.GetPathTemplate()
	if err != nil {
		return err
	}
	queries, err := route.GetQueriesTemplates()
	if err != nil {
		return err
	}
	if len(queries) > 0 {
		slog.Info("route configured", "path", path, "queries", queries)
	} else {
		slog.Info("route configured", "path", path)
	}
	return nil
}
