package main

import (
	"flag"
)

var verbose bool

func init() {
	mode := flag.NewFlagSet("mode", flag.ExitOnError)
	mode.BoolVar(&verbose, "v", false, "verbose mode")
}
