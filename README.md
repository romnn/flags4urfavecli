## flags4urfavecli

[![Build Status](https://travis-ci.com/romnnn/flags4urfavecli.svg?branch=master)](https://travis-ci.com/romnnn/flags4urfavecli)
[![GitHub](https://img.shields.io/github/license/romnnn/flags4urfavecli)](https://github.com/romnnn/flags4urfavecli)[![GoDoc](https://godoc.org/github.com/romnnn/flags4urfavecli?status.svg)](https://godoc.org/github.com/romnnn/flags4urfavecli)[![Test Coverage](https://codecov.io/gh/romnnn/flags4urfavecli/branch/master/graph/badge.svg)](https://codecov.io/gh/romnnn/flags4urfavecli)

Your description goes here...



#### Usage as a library

```golang
import "github.com/romnnn/flags4urfavecli"
```

For more examples, see `examples/`.


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
