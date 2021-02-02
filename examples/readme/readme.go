package main

import (
	"fmt"
	"log"
	"os"

	"github.com/romnn/flags4urfavecli/values"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "sample CLI app",
		Usage: "This demonstrates the usage of additional flags",
		Flags: []cli.Flag{
			&cli.GenericFlag{
				Name: "format",
				Value: &values.EnumValue{
					Enum:    []string{"json", "xml", "csv"},
					Default: "xml",
				},
				EnvVars: []string{"FILEFORMAT"},
				Usage:   "input file format",
			},
		},
		Action: func(ctx *cli.Context) error {
			fmt.Printf("Format is: %s\n", ctx.String("format"))
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
