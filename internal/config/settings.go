package config

import (
	"errors"
)

// Settings stores runtime application configuration
type Settings struct {
	ServiceAddress string
	ServicePort    uint64

	Selectors map[string]selectorSection
}

// SelectorsNames returns the set of all configured selectors names.
func (s Settings) SelectorsNames() []string {
	keys := make([]string, 0)
	for key := range s.Selectors {
		keys = append(keys, key)
	}
	return keys
}

// Parse parses a TOML configuration file and default values
// into runtime settings
func Parse(fpath string) (Settings, error) {
	base := defaultSettings()

	settings, err := parseConfig(fpath, base)
	if err != nil {
		return Settings{}, err
	}

	if err := validate(settings); err != nil {
		return Settings{}, err
	}

	return settings, nil
}

// defaultSettings returns default settings for all commands
func defaultSettings() Settings {
	return Settings{
		ServiceAddress: "0.0.0.0",
		ServicePort:    9598,

		Selectors: make(map[string]selectorSection),
	}
}

// validate sanity-checks all settings
func validate(cfg Settings) error {
	if len(cfg.Selectors) == 0 {
		return errors.New("no selectors configured")
	}

	return nil
}
