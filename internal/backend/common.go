package backend

import (
	"context"
	"io"
)

// MetricsSource is the general interface implemented by all backends
type MetricsSource interface {
	OpenMetrics(context.Context) (io.ReadCloser, error)
}
