package router

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a router with the desired routes attached
func NewRouter(routes []Route) *mux.Router {
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
	//nolint:gosec // values are JSON-encoded by slog's handler; injection is not a risk
	slog.Info("request received", "method", req.Method, "path", req.URL.Path, "query", req.URL.RawQuery)

	for k, v := range h.Headers {
		w.Header().Add(k, v)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(h.Status)

	if _, err := w.Write([]byte(h.Response)); err != nil {
		slog.Error("failed to write response", "error", err)
	}
}

func checkStatus(status int) int {
	if status == 0 {
		return http.StatusOK
	}

	return status
}
