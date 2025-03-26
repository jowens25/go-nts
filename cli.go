package main

import (
	"flag"
	"log"
	"os"
)

func cli() {

	// define top level commands
	status := flag.NewFlagSet("status", flag.ExitOnError)
	// if the flag is empty or false assume read
	// if the flag is not empty or true assume write

	// define module level commands
	// core
	core := flag.NewFlagSet("core", flag.ExitOnError)
	showCore := core.Bool("ls", false, "show the core")
	blockSize := core.Int64("blocksize", 16, "block size")
	typeInstanceReg := core.Int64("type", 0x00000000, "core type")
	BaseAddrLReg := core.Int64("BaseAddrLReg", 0x00000004, "address low")
	BaseAddrHReg := core.Int64("BaseAddrHReg", 0x00000008, "address high")
	IrqMaskReg := core.Int64("IrqMaskReg", 0x0000000C, "interrupt mask")
	// ntp
	ntp := flag.NewFlagSet("ntp", flag.ExitOnError)
	showNtp := ntp.Bool("ls", false, "show the ntp config")
	ntpEnable := ntp.Bool("enable", false, "enabled ntp server")
	ntpMac := ntp.String("mac", "00:00:00:00:00:00", "ntp server mac address")

	switch os.Args[1] {

	case "status":
		status.Parse(os.Args[2:])
		getStatus()
	case "core":
		if *showCore {

			log.Println("blockSize", *blockSize)
			log.Println("typeInstanceReg", *typeInstanceReg)
			log.Println("BaseAddrLReg", *BaseAddrLReg)
			log.Println("BaseAddrHReg", *BaseAddrHReg)
			log.Println("IrqMaskReg", *IrqMaskReg)

		}
		core.Parse(os.Args[2:])
	case "ntp":
		ntp.Parse(os.Args[2:])

		if !*showNtp {
			log.Println("made these changes to ntp core")
			ntp.Visit(func(f *flag.Flag) { log.Println("updated: ", f.Name) })
		}

		log.Println("NTP Enabled: ", *ntpEnable)
		log.Println("NTP Mac Addr: ", *ntpMac)

	default:
		log.Println("Please enter a valid command")

	}

}

func getStatus() {
	log.Println("this is the status: 100%")
}
