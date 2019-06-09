package backend

import (
	"context"
	"errors"
	"io"
)

// DbusBackendKind is the configuration kind for the dbus backend
const DbusBackendKind = "dbus"

// DbusBackend is the dbus backend
type DbusBackend struct {
	Path string
}

// OpenMetrics returns the endpoint for this backend
func (db DbusBackend) OpenMetrics(ctx context.Context) (io.ReadCloser, error) {
	return nil, errors.New("unimplemented")
}
