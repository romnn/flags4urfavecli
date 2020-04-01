package main

import (
	"fmt"

	"github.com/romnnn/flags4urfavecli"
)

func run() string {
	return flags4urfavecli.Shout("This is an example")
}

func main() {
	fmt.Println(run())
}
