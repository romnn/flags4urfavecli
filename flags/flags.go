package flags

import (
	"github.com/romnnn/flags4urfavecli/values"
	"github.com/urfave/cli/v2"
)

var (
	// LogLevelFlag ...
	LogLevelFlag = cli.GenericFlag{
		Name: "log",
		Value: &values.EnumValue{
			Enum:    []string{"info", "debug", "warn", "fatal", "trace", "error", "panic"},
			Default: "info",
		},
		Aliases: []string{"log-level"},
		EnvVars: []string{"LOG", "LOG_LEVEL"},
		Usage:   "Log level",
	}
)
