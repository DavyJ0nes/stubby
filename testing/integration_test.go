//go:build integration

package testing

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var httpClient = &http.Client{}

func TestIntegration_Basic(t *testing.T) {
	resp, err := get(t, getHost()+"/basic")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, map[string]string{"message": "foo"}, getBody(t, resp))
}

func TestIntegration_Auth(t *testing.T) {
	resp, err := get(t, getHost()+"/auth")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, map[string]string{"message": "unauthorized"}, getBody(t, resp))
}

func TestIntegration_QueryParams(t *testing.T) {
	resp, err := get(t, getHost()+"/query-params?query=blah")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, map[string]string{"bar": "42"}, getBody(t, resp))
}

func get(t *testing.T, url string) (*http.Response, error) {
	t.Helper()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return httpClient.Do(req)
}

func getHost() string {
	if host := os.Getenv("STUBBY_ADDR"); host != "" {
		return host
	}
	return "http://localhost:8080"
}

func getBody(t *testing.T, resp *http.Response) map[string]string {
	t.Helper()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	result := map[string]string{}
	require.NoError(t, json.Unmarshal(body, &result))
	return result
}
