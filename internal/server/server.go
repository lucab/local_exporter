package server

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"

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

// ErrorResponse handles errors
func ErrorResponse(w http.ResponseWriter, err error, selector string, code int) {
	logrus.WithFields(logrus.Fields{
		"selector": selector,
		"code":     code,
		"value":    err.Error(),
	}).Error("endpoint error")

	w.WriteHeader(code)

	return
}
