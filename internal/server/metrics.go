package server

import (
	"errors"
	"net/http"
)

const (
	// MetricsEndpoint is the endpoint for exposing own internal metrics.
	MetricsEndpoint = "/metrics"
)

// MetricsHandler is the handler for the `/metrics` endpoint.
func (le *LocalExporter) MetricsHandler() http.Handler {
	handler := func(w http.ResponseWriter, req *http.Request) {
		if le == nil {
			ErrorResponse(w, errNilServer, "", 500)
			return
		}

		err := errors.New("metrics: unimplemented")
		ErrorResponse(w, err, "", 500)
		return

	}

	return http.HandlerFunc(handler)
}
