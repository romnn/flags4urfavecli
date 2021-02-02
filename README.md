## flags4urfavecli

[![Build Status](https://github.com/romnn/flags4urfavecli/workflows/test/badge.svg)](https://github.com/romnn/flags4urfavecli/actions)
[![GitHub](https://img.shields.io/github/license/romnn/flags4urfavecli)](https://github.com/romnn/flags4urfavecli)
[![GoDoc](https://godoc.org/github.com/romnn/flags4urfavecli?status.svg)](https://godoc.org/github.com/romnn/flags4urfavecli)
[![Test Coverage](https://codecov.io/gh/romnn/flags4urfavecli/branch/master/graph/badge.svg)](https://codecov.io/gh/romnn/flags4urfavecli)

Extends the [github.com/urfave/cli/v2](https://github.com/urfave/cli/v2) CLI package for golang with some useful flags and values you might want to use with your CLI.

**Note:** This package is intended to be used with `v2` only!


#### Usage

```golang
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
```

For more examples, see `examples/`. You can try the example above with:
```bash
go run github.com/romnn/flags4urfavecli/examples/readme --format json
```

#### Extensions

As of now, the following generic value wrappers are implemented:

- `EnumValue`
- `TimestampValue`

Also, there are some commonly used pre-configured generic flags:

- `LogLevelFlag`

If you wrote your own, please feel free to share and submit a pull request!

#### Development

######  Prerequisites

Before you get started, make sure you have installed the following tools::

    $ python3 -m pip install -U cookiecutter>=1.4.0
    $ python3 -m pip install pre-commit bump2version invoke ruamel.yaml halo
    $ go get -u golang.org/x/tools/cmd/goimports
    $ go get -u golang.org/x/lint/golint
    $ go get -u github.com/fzipp/gocyclo
    $ go get -u github.com/mitchellh/gox  # if you want to test building on different architectures

**Remember**: To be able to excecute the tools downloaded with `go get`, 
make sure to include `$GOPATH/bin` in your `$PATH`.
If `echo $GOPATH` does not give you a path make sure to run
(`export GOPATH="$HOME/go"` to set it). In order for your changes to persist, 
do not forget to add these to your shells `.bashrc`.

With the tools in place, it is strongly advised to install the git commit hooks to make sure checks are passing in CI:
```bash
invoke install-hooks
```

You can check if all checks pass at any time:
```bash
invoke pre-commit
```

Note for Maintainers: After merging changes, tag your commits with a new version and push to GitHub to create a release:
```bash
bump2version (major | minor | patch)
git push --follow-tags
```

#### Note

This project is still in the alpha stage and should not be considered production ready.
