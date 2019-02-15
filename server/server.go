package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davyj0nes/stubby"
	"github.com/gorilla/mux"
)

func NewServer(routes []stubby.Route) *mux.Router {
	r := mux.NewRouter()

	for _, route := range routes {
		h := Handler{
			Response: route.Response,
			Status:   checkStatus(route.Status),
		}

		r.Handle(route.Path, &h)
	}

	return r
}

type Handler struct {
	Response string
	Status   int
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("received (%s) request to %s", req.Method, req.URL.String())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.Status)

	fmt.Fprintf(w, h.Response)
}

func checkStatus(status int) int {
	if status == 0 {
		return http.StatusOK
	}

	return status
}
