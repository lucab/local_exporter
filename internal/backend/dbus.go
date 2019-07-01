package backend

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"

	"github.com/godbus/dbus"
)

// DbusBackendKind is the configuration kind for the dbus backend
const DbusBackendKind = "dbus"

// DbusBackend is the dbus backend
type DbusBackend struct {
	Bus         string
	Destination string
	Method      string
	Path        string
}

// OpenMetrics returns the endpoint for this backend
func (db DbusBackend) OpenMetrics(ctx context.Context) (io.ReadCloser, error) {
	conn, err := db.connect(ctx)
	if err != nil {
		return nil, err
	}
	obj := conn.Object(db.Destination, dbus.ObjectPath(db.Path))

	var result []byte
	call := obj.CallWithContext(ctx, db.Method, 0)
	if call.Err != nil {
		return nil, call.Err
	}
	if err := call.Store(&result); err != nil {
		return nil, err
	}

	out := ioutil.NopCloser(bytes.NewReader(result))
	return out, nil
}

// Connect to the bus
func (db DbusBackend) connect(ctx context.Context) (*dbus.Conn, error) {
	switch db.Bus {
	case "system":
		return dbus.SystemBus()
	case "session":
		return dbus.SessionBus()
	default:
		return nil, errors.New("invalid bus type")
	}

}
