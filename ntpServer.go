package main

import (
	"flag"
	"fmt"
	"log"
)

type ValueSet map[string]any

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

	Enabled    bool
	MacAddr    string
	VlanEnable bool
	VlanValue  string

	IpMode            string
	UnicastMode       bool
	MulticastMode     bool
	BroadcastMode     bool
	PrecisionValue    rune
	PollIntervalValue string
	StratumValue      string
	ReferenceId       string
	IpAddrMode        string
	IpAddr            string
	Commands          *flag.FlagSet

	UtcSmearing         bool
	UtcLeap61InProgress bool
	UtcLeap59InProgress bool
	UtcLeap61           bool
	UtcLeap59           bool
	UtcOffsetVal        bool
	UtcOffsetValue      int64

	RequestsValue        string
	ResponsesValue       string
	RequestsDroppedValue string
	BroadcastsValue      string

	VersionValue  string
	ClearCounters bool
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

	Enabled:       false,
	MacAddr:       "NA",
	VlanEnable:    false,
	VlanValue:     "NA",
	IpMode:        "NA",
	UnicastMode:   false,
	MulticastMode: false,
	BroadcastMode: false,

	UtcSmearing:         false,
	UtcLeap61InProgress: false,
	UtcLeap59InProgress: false,
	UtcLeap61:           false,
	UtcLeap59:           false,
	UtcOffsetVal:        false,
	UtcOffsetValue:      0,

	RequestsValue:        "NA",
	ResponsesValue:       "NA",
	RequestsDroppedValue: "NA",
	BroadcastsValue:      "NA",

	ClearCounters: false,

	PrecisionValue:    'N',
	PollIntervalValue: "NA",
	StratumValue:      "NA",

	ReferenceId: "NA",
	IpAddrMode:  "NA",
	IpAddr:      "NA",

	Commands: flag.NewFlagSet("ntp", flag.ExitOnError),
}

func init() {
	ntpServer.Commands.Bool("ls", false, "`list` ntp server configuration")
	ntpServer.Commands.BoolVar(&verbose, "v", false, "`VERBOSE` ntp server configuration")
	ntpServer.Commands.BoolVar(&ntpServer.Enabled, "enable", false, "enable ntp `server`")

	ntpServer.Commands.StringVar(&ntpServer.IpAddr, "ip", "0.0.0.0", "ntp server ip `address`")
	ntpServer.Commands.StringVar(&ntpServer.MacAddr, "mac", "0:0:0:0", "ntp server mac `address`")

}

func parseNtpFlags(name string, value string) {

	switch name {
	case "ls":
		listNtpServer()

	case "enable":
		enableNtpServer(ntpServer.Enabled)

	case "ip":
		if value == "-ls" {
			log.Println("show the ip addr: ", ntpServer.IpAddr)
		} else {
			log.Println("you want to set the ip to: ", ntpServer.IpAddr)
		}

	case "mac":
		//log.Println("you want to set the ma to: ", ntpServer.MacAddr)

	case "v":
		if verbose {
			fmt.Println("VERBOSE: NTP")
		}

	default:
		log.Println("ntp command not known: try -ls")

	}

}

func listNtpServer() {
	thiscore := Core{}
	readCoreConfig(types.NtpServerCoreType, &thiscore)
	readCoreValues(thiscore)
	fmt.Println("ntpServer.Enabled:              ", ntpServer.Enabled)
	fmt.Println("ntpServer.MacAddr:              ", ntpServer.MacAddr)
	fmt.Println("ntpServer.VlanEnabled:          ", ntpServer.VlanEnable)
	fmt.Println("ntpServer.VlanValue:            ", ntpServer.VlanValue)
	fmt.Println("ntpServer.IpMode:               ", ntpServer.IpMode)
	fmt.Println("ntpServer.UnicastMode:          ", ntpServer.UnicastMode)
	fmt.Println("ntpServer.MulticastMode:        ", ntpServer.MulticastMode)
	fmt.Println("ntpServer.BroadcastMode:        ", ntpServer.BroadcastMode)
	fmt.Println("ntpServer.PrecisionValue:       ", ntpServer.PrecisionValue)
	fmt.Println("ntpServer.PollIntervalValue:    ", ntpServer.PollIntervalValue)
	fmt.Println("ntpServer.StratumValue:         ", ntpServer.StratumValue)
	fmt.Println("ntpServer.ReferenceId:          ", ntpServer.ReferenceId)
	fmt.Println("ntpServer.IpAddr:               ", ntpServer.IpAddr)
	fmt.Println("ntpServer.UtcSmearing:          ", ntpServer.UtcSmearing)
	fmt.Println("ntpServer.UtcLeap61InProgress:  ", ntpServer.UtcLeap61InProgress)
	fmt.Println("ntpServer.UtcLeap59InProgress:  ", ntpServer.UtcLeap59InProgress)
	fmt.Println("ntpServer.UtcLeap61:            ", ntpServer.UtcLeap61)
	fmt.Println("ntpServer.UtcLeap59:            ", ntpServer.UtcLeap59)
	fmt.Println("ntpServer.UtcOffsetVal:         ", ntpServer.UtcOffsetVal)
	fmt.Println("ntpServer.UtcOffsetValue:       ", ntpServer.UtcOffsetValue)
	fmt.Println("ntpServer.RequestsValue:        ", ntpServer.RequestsValue)
	fmt.Println("ntpServer.ResponsesValue:       ", ntpServer.ResponsesValue)
	fmt.Println("ntpServer.RequestsDroppedValue: ", ntpServer.RequestsDroppedValue)
	fmt.Println("ntpServer.BroadcastsValue:      ", ntpServer.BroadcastsValue)
	fmt.Println("ntpServer.VersionValue:         ", ntpServer.VersionValue)

}

func enableNtpServer(state bool) {
	var data int64 = 0x00000000
	if state {
		data = 0x00000001
	}
	write_reg(0xB0020000, &data)
}
