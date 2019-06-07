package server

import (
	"errors"

	"github.com/lucab/local_exporter/internal/config"
)

var (
	// errNilServer is returned on nil server
	errNilServer = errors.New("nil server")
)

// LocalExporter is the main service
type LocalExporter struct {
	config.Settings
}
