package main

import (
	"fmt"
	"log"
)

func read_NtpServerCoreType(core Core) {

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
	}

	var temp_data int64 = 0x00000000
	// enabled
	if read_reg(core.BaseAddrLReg+ntpServer.ControlReg, &temp_data) == 0 {
		if (temp_data & 0x00000001) == 0 {
			log.Println("NtpServerEnabled: False")
		} else {
			log.Println("NtpServerEnabled: True")
		}
	} else {
		log.Println("NtpServerEnabled: False!")
	}

	// mac
	this_string := make([]byte, 0, 32)

	if read_reg(core.BaseAddrLReg+ntpServer.ConfigMac1Reg, &temp_data) == 0 {

		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>0)&0x000000FF)...)
		this_string = append(this_string, ':')

		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>8)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>16)&0x000000FF)...)
		this_string = append(this_string, ':')
		this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>24)&0x000000FF)...)
		this_string = append(this_string, ':')

		if read_reg(core.BaseAddrLReg+ntpServer.ConfigMac2Reg, &temp_data) == 0 {
			this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>0)&0x000000FF)...)
			this_string = append(this_string, ':')

			this_string = append(this_string, fmt.Sprintf("%02x", (temp_data>>8)&0x000000FF)...)

			log.Println("NtpServerMacValue: ", string(this_string))
		} else {
			log.Println("NtpServerMacValue: NA")
		}

	} else {
		log.Println("NtpServerMacValue: NA")
	}

	// vlan
	if read_reg(core.BaseAddrLReg+ntpServer.ConfigVlanReg, &temp_data) == 0 {
		if (temp_data & 0x00010000) == 0 {
			log.Println("NtpServerVlanEnable: False")
		} else {
			log.Println("NtpServerVlanEnable: True")
		}

		temp_data &= 0x0000FFFF
		log.Println("NtpServerVlanValue: ", fmt.Sprintf("0x%08x", temp_data))

	} else {
		log.Println("NtpServerVlanEnable: False")
		log.Println("NtpServerVlanValue: NA")

	}

}
