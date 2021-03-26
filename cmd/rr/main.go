package main

import (
	"github.com/fatih/color"
	"github.com/spiral/roadrunner-binary/v2/internal/pkg/cli"
	"os"
	"path/filepath"
)

// exitFn is a function for application exiting.
var exitFn = os.Exit //nolint:gochecknoglobals

// main CLI application entrypoint.
func main() { exitFn(run()) }

// run this CLI application.
func run() int {
	cmd := cli.NewCommand(filepath.Base(os.Args[0]))

	if err := cmd.Execute(); err != nil {
		_, _ = color.New(color.FgHiRed, color.Bold).Fprintln(os.Stderr, err.Error())

		return 1
	}

	return 0
}
