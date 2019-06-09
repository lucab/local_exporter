package server

import (
	"context"
	"errors"
	"io"
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
			ErrorResponse(w, errNilServer, "", 500)
			return
		}

		queryParams, ok := req.URL.Query()["selector"]
		if !ok || len(queryParams) == 0 || len(queryParams[0]) == 0 {
			err := errors.New("missing selector")
			ErrorResponse(w, err, "", 500)
			return
		}
		selector := queryParams[0]

		backend, ok := le.Selectors[selector]
		if !ok || backend == nil {
			err := errors.New("backend not found")
			ErrorResponse(w, err, selector, 404)
			return
		}

		timeout := 3 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		metricsSource, err := backend.OpenMetrics(ctx)
		if err != nil {
			ErrorResponse(w, err, selector, 500)
			return
		}
		defer metricsSource.Close()

		io.Copy(w, metricsSource)

	}

	return http.HandlerFunc(handler)
}
