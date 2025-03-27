package main

import (
	"flag"
	"fmt"
)

type Core struct {
	CoreType       int64
	InstanceNumber int64
	BaseAddrLReg   int64
	BaseAddrHReg   int64
	IrqMaskReg     int64
}

type CoreConfig struct {
	BlockSize       int64
	TypeInstanceReg int64
	BaseAddrLReg    int64
	BaseAddrHReg    int64
	IrqMaskReg      int64
	flags           *flag.FlagSet
}

var coreConfig = CoreConfig{
	BlockSize:       16,
	TypeInstanceReg: 0x00000000,
	BaseAddrLReg:    0x00000004,
	BaseAddrHReg:    0x00000008,
	IrqMaskReg:      0x0000000C,

	flags: flag.NewFlagSet("coreConfig", flag.ExitOnError),
}

func init() {
	coreConfig.flags.Bool("ls", false, "`list` core configuration")
}

func parseCoreConfigProperties(name string, value string) {

	fmt.Println("BlockSize", coreConfig.BlockSize)
	fmt.Println("TypeInstanceReg", coreConfig.TypeInstanceReg)
	fmt.Println("BaseAddrLReg", coreConfig.BaseAddrLReg)
	fmt.Println("BaseAddrHReg", coreConfig.BaseAddrHReg)
	fmt.Println("IrqMaskReg", coreConfig.IrqMaskReg)

}
