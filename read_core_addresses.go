package main

import "log"

func read_core_addresses() int {

	var temp_data int64 = 0x00000000

	for i := int64(0); i < 256; i++ {
		temp_core := Core{}

		type_addr := (0x00000000 + ((i * temp_config.BlockSize) + temp_config.TypeInstanceReg))
		//log.Println("i: ", i)
		if read_reg(type_addr, &temp_data) == 0 {
			//log.Println(temp_data)
			if (i == 0) && ((((temp_data >> int64(16)) & 0x0000FFFF) != types.ConfSlaveCoreType) || (((temp_data >> int64(0)) & 0x0000FFFF) != 1)) {

				log.Println("ERROR: not a conf block at the address expected")
				break

			} else if temp_data == 0 {
				break

			} else {
				temp_core.CoreType = ((temp_data >> 16) & 0x0000FFFF)
				temp_core.InstanceNumber = ((temp_data >> 0) & 0x0000FFFF)
			}

		} else {
			log.Fatal("Error in reading modules config")
		}

		low_addr := (0x00000000 + ((i * temp_config.BlockSize) + temp_config.BaseAddrLReg))
		if read_reg(low_addr, &temp_data) == 0 {
			temp_core.BaseAddrLReg = temp_data
		} else {
			break
		}

		high_addr := (0x00000000 + ((i * temp_config.BlockSize) + temp_config.BaseAddrHReg))
		if read_reg(high_addr, &temp_data) == 0 {
			temp_core.BaseAddrHReg = temp_data
		} else {
			break
		}

		interrupt_mask := (0x00000000 + ((i * temp_config.BlockSize) + temp_config.IrqMaskReg))
		if read_reg(interrupt_mask, &temp_data) == 0 {
			temp_core.IrqMaskReg = temp_data
		} else {
			break
		}

		temp_config.Cores = append(temp_config.Cores, temp_core)
		log.Println(temp_config)
		read_core_parameters(temp_core)
	}

	return 0
}
