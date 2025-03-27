package main

import (
	"flag"
	"fmt"
	"log"
)

type NtpServer struct {
	ControlReg           int64
	StatusReg            int64
	VersionReg           int64
	CountControlReg      int64
	CountReqReg          int64
	CountRespReg         int64
	CountReqDroppedReg   int64
	CountBroadcastReg    int64
	ConfigControlReg     int64
	ConfigModeReg        int64
	ConfigVlanReg        int64
	ConfigMac1Reg        int64
	ConfigMac2Reg        int64
	ConfigIpReg          int64
	ConfigIpv61Reg       int64
	ConfigIpv62Reg       int64
	ConfigIpv63Reg       int64
	ConfigReferenceIdReg int64
	UtcInfoControlReg    int64
	UtcInfoReg           int64

	EnableBit bool
	MacAddr   string
	IpAddr    string
	flags     *flag.FlagSet
}

var ntpServer = NtpServer{

	ControlReg:           0x00000000,
	StatusReg:            0x00000004,
	VersionReg:           0x0000000C,
	CountControlReg:      0x00000010,
	CountReqReg:          0x00000014,
	CountRespReg:         0x00000018,
	CountReqDroppedReg:   0x0000001C,
	CountBroadcastReg:    0x00000020,
	ConfigControlReg:     0x00000080,
	ConfigModeReg:        0x00000084,
	ConfigVlanReg:        0x00000088,
	ConfigMac1Reg:        0x0000008C,
	ConfigMac2Reg:        0x00000090,
	ConfigIpReg:          0x00000094,
	ConfigIpv61Reg:       0x00000098,
	ConfigIpv62Reg:       0x0000009C,
	ConfigIpv63Reg:       0x000000A0,
	ConfigReferenceIdReg: 0x000000A4,
	UtcInfoControlReg:    0x00000100,
	UtcInfoReg:           0x00000104,

	EnableBit: false,
	MacAddr:   "",
	IpAddr:    "",

	flags: flag.NewFlagSet("ntp", flag.ExitOnError),
}

func init() {
	ntpServer.flags.Bool("ls", false, "`list` ntp server configuration")
	ntpServer.flags.BoolVar(&verbose, "v", false, "`VERBOSE` ntp server configuration")

	ntpServer.flags.StringVar(&ntpServer.IpAddr, "ip", "0.0.0.0", "ntp server ip `address`")
	ntpServer.flags.BoolVar(&ntpServer.EnableBit, "enable", false, "enable ntp `server`")
	ntpServer.flags.StringVar(&ntpServer.MacAddr, "mac", "0:0:0:0", "ntp server mac `address`")

}

func parseNtpProperties(name string, value string) {

	switch name {
	case "ls":
		listNtpServer()

	case "enable":
		enableNtpServer(ntpServer.EnableBit)

	case "ip":
		log.Println("you want to set the ip to: ", ntpServer.IpAddr)

	case "mac":
		log.Println("you want to set the ma to: ", ntpServer.MacAddr)

	default:
		log.Println("ntp command not known: try -ls")

	}

}

func listNtpServer() {
	thiscore := Core{}
	read_core_config(types.NtpServerCoreType, &thiscore)
	read_core_parameters(thiscore)
	fmt.Println(thiscore)

	//read_ntp_core_config()
	// for addr in ntp block

	//read parse output
	var this_data int64 = 0x00000000

	read_reg(0xB0020000, &this_data)

	log.Println(this_data)
}

func enableNtpServer(state bool) {
	var data int64 = 0x00000000
	if state {
		data = 0x00000001
	}
	write_reg(0xB0020000, &data)
}
