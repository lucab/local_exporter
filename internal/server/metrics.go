package server

import (
	"context"
	"net/http"
	"time"
)

const (
	// MetricsEndpoint is the endpoint for exposing metrics.
	MetricsEndpoint = "/metrics"
)

// MetricsHandler is the handler for the `/metrics` endpoint.
func (le *LocalExporter) MetricsHandler() http.Handler {
	handler := func(w http.ResponseWriter, req *http.Request) {
		if le == nil {
			http.Error(w, errNilServer.Error(), 500)
			return
		}

		timeout := 3 * time.Second
		_, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
	}

	return http.HandlerFunc(handler)
}
