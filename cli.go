package main

import (
	"flag"
	"fmt"
	"os"
)

func cli() {

	if len(os.Args) == 1 {
		parseRootCmdNoFlags()

	} else {
		parseCommands()
	}

}

func parseRootCmdNoFlags() {
	fmt.Println("root command without flags....")
	flag.Parse()
	flag.PrintDefaults()

}

func parseCommands() {
	switch os.Args[1] {

	case "connect":

		if connect() == 0 {
			fmt.Println("connected")
		} else {
			fmt.Println("not connected")
		}

	case "core":
		coreConfig.cmd.Parse(os.Args[2:])

		coreConfig.cmd.Visit(func(f *flag.Flag) {
			parseCoreConfigProperties(f.Name, f.Value.String())
		})

	case "ntp":
		ntpServer.Commands.Parse(os.Args[2:])
		ntpServer.Commands.Visit(func(f *flag.Flag) {
			parseNtpFlags(f.Name, f.Value.String())
		})

	case "help":
		fmt.Println("Here is how to use the Novus Time Server Configuration Tool: ")

	default:
		flag.Parse()

		fmt.Println("root command with flags")
	}
}
