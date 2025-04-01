package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var writeFlag bool = false

func cli() {

	if len(os.Args) >= 2 {
		modules()
	} else {
		fmt.Println("Please enter an module (ntp ip 10.10.10.1)")
	}

}

func properties() {
	switch os.Args[3] {
		case ""
	}
}

func modules() {
	switch os.Args[2] {
	case "ntp":
		fmt.Println("ntp")
	case "other":
		fmt.Println("other")
	case "ptp":
		fmt.Println("ptp")
	default:
		log.Println("no module selected?")
	}
}



func parseNtpFlags(f *flag.Flag) {

	switch f.Name {

	case "ip":
		ntpServerIpOperation()

	case "mac":
		ntpServerMacOperation()

	}
}

func ntpServerIpOperation() {

	if writeFlag {
		log.Println("write ntp ip reg", ntpServer.IpAddr)
	} else {
		log.Println("read ntp ip reg", ntpServer.IpAddr)
	}

}

func ntpServerMacOperation() {
	if writeFlag {
		log.Println("write ntp mac reg")
	} else {
		log.Println("read ntp mac reg")
	}
}

//func parseRootCmdNoFlags() {
//	fmt.Println("root command without flags....")
//	flag.Parse()
//	flag.PrintDefaults()
//
//}

/*
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
			parseNtpFlags(f.Name, f.Value.String(), f.)
		})

	case "help":
		fmt.Println("Here is how to use the Novus Time Server Configuration Tool: ")

	default:
		flag.Parse()

		fmt.Println("root command with flags")
	}
}

*/
