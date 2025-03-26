package main

import (
	"flag"
	"log"
	"os"
)

var en_data int64 = 0x00000001
var disen_data int64 = 0x00000000

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
	ntp.Bool("ls", false, "show the ntp config")
	ntp.Bool("enable", false, "enabled ntp server")
	ntp.String("mac", "00:00:00:00:00:00", "ntp server mac address")

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

		ntp.Visit(func(f *flag.Flag) {
			parseNtpParameters(f.Name, f.Value.String())
		})

		//read_reg(0xB0020000, &en_data)

		//write_reg(0xB0020000, &disen_data)

	default:
		log.Println("Please enter a valid command")

	}

}

func parseNtpParameters(name string, value string) {
	if value == "true" {
		write_reg(0xB0020000, &en_data)

	} else if value == "false" {
		write_reg(0xB0020000, &disen_data)
	}

	if name == "ls" {
		log.Println(value)
		log.Println(read_reg(0xB0020000, &disen_data))

		log.Println(disen_data)
	}

	// enable ntp

}

func getStatus() {
	log.Println("this is the status: 100%")
}
