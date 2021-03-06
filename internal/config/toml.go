package config

import (
	"errors"

	"github.com/BurntSushi/toml"
)

// tomlConfig is the top-level TOML configuration fragment
type tomlConfig struct {
	Service *serviceSection `toml:"service"`
	Bridge  *bridgeSection  `toml:"bridge"`
}

// serviceSection holds the optional `service` fragment
type serviceSection struct {
	Address *string `toml:"address"`
	Port    *uint64 `toml:"port"`
	TLS     *bool   `toml:"tls"`
}

// bridgeSection holds the optional `bridge` fragment
type bridgeSection struct {
	Selectors             map[string]selectorSection `toml:"selectors"`
	DbusSessionBusAddress *string                    `toml:"dbus_session_bus_address"`
	DbusSystemBusAddress  *string                    `toml:"dbus_system_bus_address"`
}

// selectorSection holds the optional `selector` fragment
type selectorSection struct {
	Kind string `toml:"kind"`
	Path string `toml:"path"`

	// DBus specific
	Bus         string `toml:"bus"`
	Destination string `toml:"destination"`
	Method      string `toml:"method"`
}

// parseConfig tries to parse and merge TOML config and default settings
func parseConfig(fpath string, defaults Settings) (Settings, error) {
	cfg := tomlConfig{}
	if _, err := toml.DecodeFile(fpath, &cfg); err != nil {
		return Settings{}, err
	}

	runSettings := defaults

	if err := mergeToml(&runSettings, cfg); err != nil {
		return Settings{}, err
	}

	return runSettings, nil
}

// mergeToml applies a TOML configuration fragment on top of existing settings
func mergeToml(settings *Settings, cfg tomlConfig) error {
	if settings == nil {
		return errors.New("nil settings")
	}

	if cfg.Service != nil {
		if err := mergeService(settings, *cfg.Service); err != nil {
			return err
		}
	}
	if cfg.Bridge != nil {
		if err := mergeBridge(settings, *cfg.Bridge); err != nil {
			return err
		}
	}

	return nil
}

func mergeService(settings *Settings, cfg serviceSection) error {
	if settings == nil {
		return errors.New("nil settings")
	}

	if cfg.Address != nil {
		settings.ServiceAddress = *cfg.Address
	}
	if cfg.Port != nil {
		settings.ServicePort = *cfg.Port
	}
	if cfg.TLS != nil {
		settings.ServiceTLS = *cfg.TLS
	}

	return nil
}

func mergeBridge(settings *Settings, cfg bridgeSection) error {
	if settings == nil {
		return errors.New("nil settings")
	}

	if cfg.DbusSessionBusAddress != nil {
		settings.DbusSessionBusAddress = *cfg.DbusSessionBusAddress
	}
	if cfg.DbusSystemBusAddress != nil {
		settings.DbusSystemBusAddress = *cfg.DbusSystemBusAddress
	}

	for selector, entry := range cfg.Selectors {
		backend, err := ParseBackend(entry)
		if err != nil {
			return err
		}

		settings.Selectors[selector] = backend
	}

	return nil
}
