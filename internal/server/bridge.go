package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	// BridgeEndpoint is the endpoint for exposing bridged metrics.
	BridgeEndpoint = "/bridge"
)

// BridgeHandler is the handler for the `/bridge` endpoint.
func (le *LocalExporter) BridgeHandler() http.Handler {
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

		source, err := backend.OpenMetrics(ctx)
		if err != nil {
			ErrorResponse(w, err, selector, 500)
			return
		}
		defer source.Close()

		io.Copy(w, source)

	}

	return http.HandlerFunc(handler)
}
