package backend

import (
	"context"
	"io"
	"net"
)

// UnixBackendKind is the configuration kind for the unix-socket backend
const UnixBackendKind = "uds"

// UnixBackend is the unix-socket backend
type UnixBackend struct {
	Path string
}

// OpenMetrics returns the endpoint for this backend
func (ub UnixBackend) OpenMetrics(ctx context.Context) (io.ReadCloser, error) {
	dialer := net.Dialer{}
	conn, err := dialer.Dial("unix", ub.Path)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
