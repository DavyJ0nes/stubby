package router_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davyj0nes/stubby"
	"github.com/davyj0nes/stubby/router"
)

type customHeader struct {
	Key   string
	Value string
}

type expected struct {
	body         string
	customHeader customHeader
	status       int
}

func TestNewRouter(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		routes []stubby.Route
		want   expected
	}{
		{
			name: "no routes gives 404",
			path: "/",
			want: expected{
				body:   "404 page not found\n",
				status: http.StatusNotFound,
			},
		},
		{
			name: "supplied route give expected response",
			path: "/hey",
			routes: []stubby.Route{
				{
					Path:     "/hey",
					Response: "Yo, Yo, Yo",
					Status:   http.StatusOK,
				},
			},
			want: expected{
				body:   "Yo, Yo, Yo",
				status: http.StatusOK,
			},
		},
		{
			name: "supplied route without status still give expected response",
			path: "/salut",
			routes: []stubby.Route{
				{
					Path:     "/salut",
					Response: "Yo, Yo, Yo",
				},
			},
			want: expected{
				body:   "Yo, Yo, Yo",
				status: http.StatusOK,
			},
		},
		{
			name: "supplied route with query params matches the right handler",
			path: "/things?with_some_param=foo",
			routes: []stubby.Route{
				{
					Path:     "/things",
					Queries:  []string{"with_some_param", "foo"},
					Response: "stuff",
				},
			},
			want: expected{
				body:   "stuff",
				status: http.StatusOK,
			},
		},
		{
			name: "supplied route with headers matches the right handler",
			path: "/head",
			routes: []stubby.Route{
				{
					Path:     "/head",
					Headers:  map[string]string{"Custom": "custom"},
					Response: "at the head",
				},
			},
			want: expected{
				body:         "at the head",
				status:       http.StatusOK,
				customHeader: customHeader{Key: "Custom", Value: "custom"},
			},
		},
		{
			name: "supplied route with parameterised path value matches correct route",
			path: "/wildcard/test",
			routes: []stubby.Route{
				{
					Path:     "/wildcard/{wildcard}",
					Response: "wildcard response",
				},
			},
			want: expected{
				body:   "wildcard response",
				status: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := router.NewRouter(tt.routes)
			ts := httptest.NewServer(r)
			defer ts.Close()

			res, err := http.Get(ts.URL + tt.path)
			if err != nil {
				t.Fatal(err)
			}

			if res.StatusCode != tt.want.status {
				t.Errorf("expected: (%d), got: (%d)", tt.want.status, res.StatusCode)
			}

			headerVal := getResponseHeader(t, res, tt.want.customHeader.Key)
			if headerVal != tt.want.customHeader.Value {
				t.Errorf("expected: (%s), got: (%s)", tt.want.customHeader.Value, headerVal)
			}

			body := getResponseBody(t, res)
			if body != tt.want.body {
				t.Errorf("expected: (%s), got: (%s)", tt.want.body, body)
			}
		})
	}
}

func getResponseBody(t *testing.T, r *http.Response) string {
	t.Helper()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatal(err)
	}

	return string(body)
}

func getResponseHeader(t *testing.T, r *http.Response, wantHeader string) string {
	t.Helper()

	return r.Header.Get(wantHeader)
}
