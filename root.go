package main

import (
	"flag"
	"log"
)

type Root struct {
	cmd *flag.FlagSet
}

var root = Root{

	cmd: flag.NewFlagSet("", flag.ExitOnError),
}

func init() {

	flag.Bool("ls", false, "`list` all core configurations")
	flag.Bool("ps", false, "`list` all core configurations")

}

func parseRootFlags(name string, value string) {

	switch name {
	case "ls":
		listNtpServer()

	default:
		log.Println("root command not known: try -ls")

	}

}
