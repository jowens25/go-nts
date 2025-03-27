package main

import (
	"fmt"
	"log"
)

func read_NtpServerCoreType(core Core) {

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

	// mode & server config
	if read_reg(core.BaseAddrLReg+ntpServer.ConfigModeReg, &temp_data) == 0 {
		if ((temp_data >> 0) & 0x00000003) == 1 {
			log.Println("NtpServerIpModeValue: IPv4")
		} else if ((temp_data >> 0) & 0x00000003) == 2 {
			log.Println("NtpServerIpModeValue: IPv6")
		} else {
			log.Println("NtpServerIpModeValue: NA")
		}

		if (temp_data & 0x00000010) == 0 {
			log.Println("NtpServerUnicastMode: false")
		} else {
			log.Println("NtpServerUnicastMode: true")
		}

		if (temp_data & 0x00000020) == 0 {
			log.Println("NtpServerMulticastMode: false")
		} else {
			log.Println("NtpServerMulticastMode: true")
		}

		if (temp_data & 0x00000040) == 0 {
			log.Println("NtpServerBroadcastMode: false")
		} else {
			log.Println("NtpServerBroadcastMode: true")
		}

		log.Println("NtpServerPrecisionValue: ", int8(((temp_data >> 8) & 0x000000FF)))
		log.Println("NtpServerPollIntervalValue: ", ((temp_data >> 16) & 0x000000FF))
		log.Println("NtpServerStratumValue ", ((temp_data >> 24) & 0x000000FF))

	} else {
		log.Println("NtpServerIpModeValue: NA")
		log.Println("NtpServerUnicastMode: false")
		log.Println("NtpServerMulticastMode: false")
		log.Println("NtpServerBroadcastMode: false")

		log.Println("NtpServerPrecisionValue: ", "NA")
		log.Println("NtpServerPollIntervalValue: ", "NA")
		log.Println("NtpServerStratumValue ", "NA")

	}

	// reference id // no ref on UI??
	if read_reg(core.BaseAddrLReg+ntpServer.ConfigReferenceIdReg, &temp_data) == 0 {
		var temp_string []byte
		temp_string = append(temp_string, byte(((temp_data >> 24) & 0x000000FF)))
		temp_string = append(temp_string, byte(((temp_data >> 16) & 0x000000FF)))
		temp_string = append(temp_string, byte(((temp_data >> 8) & 0x000000FF)))
		temp_string = append(temp_string, byte(((temp_data >> 0) & 0x000000FF)))
		log.Println("NtpServerReferenceIdValue: ", fmt.Sprintf("0x%08x", temp_string)) // TODO
	} else {
		log.Println("NtpServerReferenceIdValue: NA")
	}

}
