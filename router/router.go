package router

import (
	"log"
	"net/http"

	"github.com/davyj0nes/stubby"
	"github.com/gorilla/mux"
)

// NewRouter creates a router with the desired routes attached
func NewRouter(routes []stubby.Route) *mux.Router {
	r := mux.NewRouter()

	for _, route := range routes {
		h := Handler{
			Response: route.Response,
			Status:   checkStatus(route.Status),
			Headers:  route.Headers,
		}

		r.NewRoute().
			Path(route.Path).
			Queries(route.Queries...).
			Handler(&h)
	}

	return r
}

// Handler describes an HTTP handler with set response and status code
type Handler struct {
	Response string
	Status   int
	Headers  map[string]string
}

// ServeHTTP is used to adhere to the http.Handler interface
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("received (%s) request to %s", req.Method, req.URL.String())

	for k, v := range h.Headers {
		w.Header().Add(k, v)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.Status)

	_, err := w.Write([]byte(h.Response))
	if err != nil {
		log.Printf("err writing response: %s", err)
	}

}

func checkStatus(status int) int {
	if status == 0 {
		return http.StatusOK
	}

	return status
}
