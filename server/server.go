package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davyj0nes/stubby/config"
	"github.com/gorilla/mux"
)

func NewServer(routes []config.Route) *mux.Router {
	mux := mux.NewRouter()

	for _, route := range routes {
		mux.Handle(route.Path, &Handler{route.Response})
	}

	return mux
}

type Handler struct {
	Response string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("received (%s) request to %s", req.Method, req.URL.String())

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, h.Response)
}
