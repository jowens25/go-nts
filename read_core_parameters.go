package main

import (
	"log"
)

func read_core_parameters(core Core) {

	switch core.CoreType {

	case types.ConfSlaveCoreType:
		log.Print("Core type: ConfSlaveCoreType")

	case types.ClkClockCoreType:
		read_ClkClockCoreType(core)

	case types.ClkSignalGeneratorCoreType:
		log.Print("Core type: ClkSignalGeneratorCoreType")

	case types.ClkSignalTimestamperCoreType:
		log.Print("Core type: ClkSignalTimestamperCoreType")

	case types.IrigSlaveCoreType:
		log.Print("Core type: IrigSlaveCoreType")

	case types.IrigMasterCoreType:
		log.Print("Core type: IrigMasterCoreType")

	case types.PpsSlaveCoreType:
		log.Print("Core type: PpsSlaveCoreType")

	case types.PpsMasterCoreType:
		log.Print("Core type: PpsMasterCoreType")

	case types.PtpOrdinaryClockCoreType:
		log.Print("Core type: PtpOrdinaryClockCoreType")

	case types.PtpTransparentClockCoreType:
		log.Print("Core type: PtpTransparentClockCoreType")

	case types.PtpHybridClockCoreType:
		log.Print("Core type: PtpHybridClockCoreType")

	case types.RedHsrPrpCoreType:
		log.Print("Core type: RedHsrPrpCoreType")

	case types.RtcSlaveCoreType:
		log.Print("Core type: RtcSlaveCoreType")

	case types.RtcMasterCoreType:
		log.Print("Core type: RtcMasterCoreType")

	case types.TodSlaveCoreType:
		log.Print("Core type: TodSlaveCoreType ")

	case types.TodMasterCoreType:
		log.Print("Core type: TodMasterCoreType")

	case types.TapSlaveCoreType:
		log.Print("Core type: TapSlaveCoreType")

	case types.DcfSlaveCoreType:
		log.Print("Core type: DcfSlaveCoreType")

	case types.DcfMasterCoreType:
		log.Print("Core type: DcfMasterCoreType")

	case types.RedTsnCoreType:
		log.Print("Core type: RedTsnCoreType")

	case types.TsnIicCoreType:
		log.Print("Core type: TsnIicCoreType")

	case types.NtpServerCoreType:
		read_NtpServerCoreType(core)
	case types.NtpClientCoreType:
		log.Print("Core type: NtpClientCoreType")

	case types.ClkFrequencyGeneratorCoreType:
		log.Print("Core type: ClkFrequencyGeneratorCoreType")

	case types.SynceNodeCoreType:
		log.Print("Core type: SynceNodeCoreType")

	case types.PpsClkToPpsCoreType:
		log.Print("Core type: PpsClkToPpsCoreType")

	case types.PtpServerCoreType:
		log.Print("Core type: PtpServerCoreType")

	case types.PtpClientCoreType:
		log.Print("Core type: PtpClientCoreType")

	case types.PhyConfigurationCoreType:
		log.Print("Core type: PhyConfigurationCoreType")

	case types.I2cConfigurationCoreType:
		log.Print("Core type: I2cConfigurationCoreType")

	case types.IoConfigurationCoreType:
		log.Print("Core type: IoConfigurationCoreType")

	case types.EthernetTestplatformType:
		log.Print("Core type: EthernetTestplatformType")

	case types.MinSwitchCoreType:
		log.Print("Core type: MinSwitchCoreType")

	case types.ConfExtCoreType:
		log.Print("Core type: ConfExtCoreType")

	default:
		log.Print("Core not found error, type: ", core.CoreType)

	}

}
