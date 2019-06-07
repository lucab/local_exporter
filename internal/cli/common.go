package cli

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/lucab/local_exporter/internal/config"
)

var (
	// localExporterCmd is the top-level cobra command for `localExporter`
	localExporterCmd = &cobra.Command{
		Use:               "local_exporter",
		PersistentPreRunE: commonSetup,
	}

	verbosity   int
	configPath  string
	runSettings *config.Settings
)

// Init initializes the CLI environment for localExporter
func Init() (*cobra.Command, error) {
	localExporterCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "/etc/local_exporter/config.toml", "path to configuration file")
	localExporterCmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "increase verbosity level")

	localExporterCmd.AddCommand(cmdServe)

	return localExporterCmd, nil
}

// commonSetup perform actions commons to all CLI subcommands
func commonSetup(cmd *cobra.Command, cmdArgs []string) error {
	if configPath == "" {
		return errors.New("empty path to configuration file")
	}
	logrus.SetLevel(verbosityLevel(verbosity))

	cfg, err := config.Parse(configPath)
	if err != nil {
		return err
	}
	runSettings = &cfg

	return nil
}

// verbosityLevel parses `-v` count into logrus log-level
func verbosityLevel(verbCount int) logrus.Level {
	switch verbCount {
	case 0:
		return logrus.WarnLevel
	case 1:
		return logrus.InfoLevel
	case 2:
		return logrus.DebugLevel
	default:
		return logrus.TraceLevel
	}
}
