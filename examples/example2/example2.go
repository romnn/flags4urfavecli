package main

import (
	"os"

	"github.com/romnnn/flags4urfavecli/values"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "sample CLI app",
		Usage: "This demonstrates the usage of additional flags",
		Flags: []cli.Flag{
			// Generic enum list flag
			&cli.GenericFlag{
				Name: "ducks",
				Value: &values.EnumListValue{
					Enum:    []string{"tick", "trick", "track"},
					Default: []string{"trick"},
				},
				EnvVars: []string{"DUCKS"},
				Usage:   "An enum list only accepts comma separated enum values",
			},
		},
		Action: func(ctx *cli.Context) error {
			log.Infof("Ducks are: %s\n", values.EnumListValue{}.Parse(ctx.String("ducks")))
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
