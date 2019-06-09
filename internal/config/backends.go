package config

import (
	"fmt"

	"github.com/lucab/local_exporter/internal/backend"
)

// ParseBackend parses a backend configuration
func ParseBackend(cfg selectorSection) (backend.MetricsSource, error) {
	switch cfg.Kind {
	case backend.FileBackendKind:
		return parseFileBackend(cfg)
	case backend.UnixBackendKind:
		return parseUnixBackend(cfg)
	case backend.DbusBackendKind:
		return parseDbusBackend(cfg)
	default:
		return nil, fmt.Errorf("unknown backend: %q", cfg.Kind)
	}
}

func parseFileBackend(cfg selectorSection) (*backend.FileBackend, error) {
	fb := backend.FileBackend{
		Path: cfg.Path,
	}
	return &fb, nil
}

func parseUnixBackend(cfg selectorSection) (*backend.UnixBackend, error) {
	ub := backend.UnixBackend{
		Path: cfg.Path,
	}
	return &ub, nil
}

func parseDbusBackend(cfg selectorSection) (*backend.DbusBackend, error) {
	db := backend.DbusBackend{
		Path: cfg.Path,
	}
	return &db, nil
}
