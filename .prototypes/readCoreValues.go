package main

import (
	"fmt"
)

func readCoreValues(core Core) {

	switch core.CoreType {

	case types.ConfSlaveCoreType:
		fmt.Print("Core type: ConfSlaveCoreType")

	case types.ClkClockCoreType:
		read_ClkClockCoreType(core)

	case types.ClkSignalGeneratorCoreType:
		fmt.Print("Core type: ClkSignalGeneratorCoreType")

	case types.ClkSignalTimestamperCoreType:
		fmt.Print("Core type: ClkSignalTimestamperCoreType")

	case types.IrigSlaveCoreType:
		fmt.Print("Core type: IrigSlaveCoreType")

	case types.IrigMasterCoreType:
		fmt.Print("Core type: IrigMasterCoreType")

	case types.PpsSlaveCoreType:
		fmt.Print("Core type: PpsSlaveCoreType")

	case types.PpsMasterCoreType:
		fmt.Print("Core type: PpsMasterCoreType")

	case types.PtpOrdinaryClockCoreType:
		fmt.Print("Core type: PtpOrdinaryClockCoreType")

	case types.PtpTransparentClockCoreType:
		fmt.Print("Core type: PtpTransparentClockCoreType")

	case types.PtpHybridClockCoreType:
		fmt.Print("Core type: PtpHybridClockCoreType")

	case types.RedHsrPrpCoreType:
		fmt.Print("Core type: RedHsrPrpCoreType")

	case types.RtcSlaveCoreType:
		fmt.Print("Core type: RtcSlaveCoreType")

	case types.RtcMasterCoreType:
		fmt.Print("Core type: RtcMasterCoreType")

	case types.TodSlaveCoreType:
		fmt.Print("Core type: TodSlaveCoreType ")

	case types.TodMasterCoreType:
		fmt.Print("Core type: TodMasterCoreType")

	case types.TapSlaveCoreType:
		fmt.Print("Core type: TapSlaveCoreType")

	case types.DcfSlaveCoreType:
		fmt.Print("Core type: DcfSlaveCoreType")

	case types.DcfMasterCoreType:
		fmt.Print("Core type: DcfMasterCoreType")

	case types.RedTsnCoreType:
		fmt.Print("Core type: RedTsnCoreType")

	case types.TsnIicCoreType:
		fmt.Print("Core type: TsnIicCoreType")

	case types.NtpServerCoreType:
		readNtpServer(core)
	case types.NtpClientCoreType:
		fmt.Print("Core type: NtpClientCoreType")

	case types.ClkFrequencyGeneratorCoreType:
		fmt.Print("Core type: ClkFrequencyGeneratorCoreType")

	case types.SynceNodeCoreType:
		fmt.Print("Core type: SynceNodeCoreType")

	case types.PpsClkToPpsCoreType:
		fmt.Print("Core type: PpsClkToPpsCoreType")

	case types.PtpServerCoreType:
		fmt.Print("Core type: PtpServerCoreType")

	case types.PtpClientCoreType:
		fmt.Print("Core type: PtpClientCoreType")

	case types.PhyConfigurationCoreType:
		fmt.Print("Core type: PhyConfigurationCoreType")

	case types.I2cConfigurationCoreType:
		fmt.Print("Core type: I2cConfigurationCoreType")

	case types.IoConfigurationCoreType:
		fmt.Print("Core type: IoConfigurationCoreType")

	case types.EthernetTestplatformType:
		fmt.Print("Core type: EthernetTestplatformType")

	case types.MinSwitchCoreType:
		fmt.Print("Core type: MinSwitchCoreType")

	case types.ConfExtCoreType:
		fmt.Print("Core type: ConfExtCoreType")

	default:
		fmt.Print("Core not found error, type: ", core.CoreType)

	}

}
