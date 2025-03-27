package main

import (
	"flag"
	"log"
	"os"
)

var coreConfig *flag.FlagSet = flag.NewFlagSet("core", flag.ExitOnError)

func cli() {

	// define top level commands
	// if the flag is empty or false assume read
	// if the flag is not empty or true assume write

	// define module level commands
	// core
	coreConfig.Bool("ls", false, "show the core")
	coreConfig.Int64("blocksize", 16, "block size")
	coreConfig.Int64("type", 0x00000000, "core type")
	coreConfig.Int64("BaseAddrLReg", 0x00000004, "address low")
	coreConfig.Int64("BaseAddrHReg", 0x00000008, "address high")
	coreConfig.Int64("IrqMaskReg", 0x0000000C, "interrupt mask")
	// ntp

	switch os.Args[1] {

	case "status":
		getStatus()
	case "core":
		parseConfig(coreConfig)
	case "ntp":
		parseConfig(ntpServer.flags)

	default:
		log.Println("Please enter a valid command")

	}

}

func parseConfig(core *flag.FlagSet) {
	core.Parse(os.Args[2:])
	switch core {
	case coreConfig:
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

func getStatus() {
	log.Println("THIS IS THE STATUS!!")
}
