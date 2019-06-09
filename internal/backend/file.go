package backend

import (
	"context"
	"io"
	"os"
)

// FileBackendKind is the configuration kind for the file backend
const FileBackendKind = "file"

// FileBackend is the file backend
type FileBackend struct {
	Path string
}

// OpenMetrics returns the endpoint for this backend
func (fb FileBackend) OpenMetrics(ctx context.Context) (io.ReadCloser, error) {
	fp, err := os.Open(fb.Path)
	if err != nil {
		return nil, err
	}

	return fp, nil
}
