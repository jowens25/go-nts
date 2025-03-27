package main

import "log"

type Core struct {
	CoreType       int64
	InstanceNumber int64
	BaseAddrLReg   int64
	BaseAddrHReg   int64
	IrqMaskReg     int64
}

type Config struct {
	BlockSize       int64
	TypeInstanceReg int64
	BaseAddrLReg    int64
	BaseAddrHReg    int64
	IrqMaskReg      int64
	Cores           []Core
}

var temp_config = Config{
	BlockSize:       16,
	TypeInstanceReg: 0x00000000,
	BaseAddrLReg:    0x00000004,
	BaseAddrHReg:    0x00000008,
	IrqMaskReg:      0x0000000C,
	Cores:           make([]Core, 0, 64),
}

func parseCoreConfigProperties(name string, value string) {

	switch name {
	case "blocksize":
		log.Println("block size case")

	default:
		log.Println("core config command not known: try -ls")

	}

}
