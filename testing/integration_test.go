// +build integration

package testing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegration_Basic(t *testing.T) {
	resp, err := http.Get(getHost() + "/basic")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	got := getBody(t, resp)
	want := map[string]string{
		"message": "foo",
	}

	assert.Equal(t, want, got)
}

func TestIntegration_Auth(t *testing.T) {
	resp, err := http.Get(getHost() + "/auth")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	got := getBody(t, resp)
	want := map[string]string{
		"message": "unauthorized",
	}

	assert.Equal(t, want, got)
}

func TestIntegration_QueryParams(t *testing.T) {
	resp, err := http.Get(getHost() + "/query-params?query=blah")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	got := getBody(t, resp)
	want := map[string]string{
		"bar": "42",
	}

	assert.Equal(t, want, got)
}

func getHost() string {
	host := os.Getenv("STUBBY_ADDR")
	if host == "" {
		host = "http://localhost:8080"
	}

	return host
}

func getBody(t *testing.T, resp *http.Response) map[string]string {
	t.Helper()

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	response := map[string]string{}
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	return response
}
