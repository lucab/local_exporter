package config

import (
	"github.com/BurntSushi/toml"
)

// tomlConfig is the top-level TOML configuration fragment
type tomlConfig struct {
	Service *serviceSection `toml:"service"`
	Metrics *metricsSection `toml:"metrics"`
}

// serviceSection holds the optional `service` fragment
type serviceSection struct {
	Address *string `toml:"address"`
	Port    *uint64 `toml:"port"`
}

// metricsSection holds the optional `metrics` fragment
type metricsSection struct {
	Selectors map[string]selectorSection `toml:"selectors"`
}

// selectorSection holds the optional `selector` fragment
type selectorSection struct {
	Kind string `toml:"kind"`
	Path string `toml:"path"`
}

// parseConfig tries to parse and merge TOML config and default settings
func parseConfig(fpath string, defaults Settings) (Settings, error) {
	cfg := tomlConfig{}
	if _, err := toml.DecodeFile(fpath, &cfg); err != nil {
		return Settings{}, err
	}
	runSettings := defaults
	mergeToml(&runSettings, cfg)

	return runSettings, nil
}

// mergeToml applies a TOML configuration fragment on top of existing settings
func mergeToml(settings *Settings, cfg tomlConfig) {
	if settings == nil {
		return
	}

	if cfg.Service != nil {
		mergeService(settings, *cfg.Service)
	}
	if cfg.Metrics != nil {
		mergeMetrics(settings, *cfg.Metrics)
	}
}

func mergeService(settings *Settings, cfg serviceSection) {
	if settings == nil {
		return
	}

	if cfg.Address != nil {
		settings.ServiceAddress = *cfg.Address
	}
	if cfg.Port != nil {
		settings.ServicePort = *cfg.Port
	}
}

func mergeMetrics(settings *Settings, cfg metricsSection) {
	if settings == nil {
		return
	}
	for selector, entry := range cfg.Selectors {
		settings.Selectors[selector] = entry
	}
}
