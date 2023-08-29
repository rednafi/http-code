package main

import (
	"flag"
	"github.com/rednafi/httpurr/src"
	"os"
	"text/tabwriter"
)

// Ldflags filled by goreleaser
var version string

func main() {
	w := tabwriter.NewWriter(flag.CommandLine.Output(), 0, 4, 4, ' ', 0)
	src.Cli(w, version, os.Exit)
}
