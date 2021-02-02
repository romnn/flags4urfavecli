package main

import (
	"os"

	"github.com/romnn/flags4urfavecli/flags"
	"github.com/romnn/flags4urfavecli/values"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	myTimestampFormat := "2006-01-02"
	app := &cli.App{
		Name:  "sample CLI app",
		Usage: "This demonstrates the usage of additional flags",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "boolean",
				Aliases: []string{"plain-bool"},
				Usage:   "This one is a normal cli flag",
			},
			// Pre configured enum field that parses log levels
			&flags.LogLevelFlag,
			// Generic enum flag
			&cli.GenericFlag{
				Name: "duck",
				Value: &values.EnumValue{
					Enum:    []string{"tick", "trick", "track"},
					Default: "trick",
				},
				EnvVars: []string{"DUCK"},
				Usage:   "An enum flag only accepts selected values",
			},
			// Generic timestamp flag
			&cli.GenericFlag{
				Name: "timestamp",
				Value: &values.TimestampValue{
					Format: &myTimestampFormat,
				},
				EnvVars: []string{"TIMESTAMP"},
				Usage:   "Parses a timestamp with a specific format",
			},
		},
		Action: func(ctx *cli.Context) error {
			if level, err := log.ParseLevel(ctx.String("log")); err == nil {
				log.SetLevel(level)
			}
			log.Infof("Boolean is: %t\n", ctx.Bool("boolean"))
			log.Infof("Log is: %s\n", ctx.String("log"))
			log.Infof("Duck is: %s\n", ctx.String("duck"))
			log.Infof("Timestamp is: %s\n", ctx.String("timestamp"))
			log.Warn(`Try to hide the messages before using "--log warn"`)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
