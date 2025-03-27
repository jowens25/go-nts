package main

import (
	"flag"
	"log"
	"os"
)

func cli() {

	switch os.Args[1] {

	case "core":
		parseConfig(coreConfig.flags)
	case "ntp":
		parseConfig(ntpServer.flags)

	default:
		log.Println("Please enter a valid command")

	}

}

func parseConfig(core *flag.FlagSet) {

	core.Parse(os.Args[2:])

	switch core {

	case coreConfig.flags:
		core.Visit(func(f *flag.Flag) {
			parseCoreConfigProperties(f.Name, f.Value.String())

		})
	case ntpServer.flags:
		core.Visit(func(f *flag.Flag) {
			parseNtpProperties(f.Name, f.Value.String())

		})

	default:
		log.Println("Core not found? ")

	}

}
