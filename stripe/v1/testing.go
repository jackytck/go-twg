package stripe

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestClient is a fake client for testing purpose.
func TestClient(t *testing.T) (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	c := &Client{
		baseURL: server.URL,
	}
	return c, mux, func() {
		server.Close()
	}
}
