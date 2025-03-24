package main

import (
	"log"
)

func read_core_parameters(core Core) {

	switch core.CoreType {

	case types.ConfSlaveCoreType:
		log.Println("We have a core type: ConfSlaveCoreType")

	case types.ClkClockCoreType:
		log.Println("We have a core type: ClkClockCoreType")
		read_ClkClockCoreType_values(core)
	/*
		case types.ClkSignalGeneratorCoreType:
			log.Println("We have a core type: ClkSignalGeneratorCoreType")

		case types.ClkSignalTimestamperCoreType:
			log.Println("We have a core type: ClkSignalTimestamperCoreType")

		case types.IrigSlaveCoreType:
			log.Println("We have a core type: IrigSlaveCoreType")

		case types.IrigMasterCoreType:
			log.Println("We have a core type: IrigMasterCoreType")

		case types.PpsSlaveCoreType:
			log.Println("We have a core type: PpsSlaveCoreType")

		case types.PpsMasterCoreType:
			log.Println("We have a core type: PpsMasterCoreType")

		case types.PtpOrdinaryClockCoreType:
			log.Println("We have a core type: PtpOrdinaryClockCoreType")

		case types.PtpTransparentClockCoreType:
			log.Println("We have a core type: PtpTransparentClockCoreType")

		case types.PtpHybridClockCoreType:
			log.Println("We have a core type: PtpHybridClockCoreType")

		case types.RedHsrPrpCoreType:
			log.Println("We have a core type: RedHsrPrpCoreType")

		case types.RtcSlaveCoreType:
			log.Println("We have a core type: RtcSlaveCoreType")

		case types.RtcMasterCoreType:
			log.Println("We have a core type: RtcMasterCoreType")

		case types.TodSlaveCoreType:
			log.Println("We have a core type: TodSlaveCoreType")

		case types.TodMasterCoreType:
			log.Println("We have a core type: TodMasterCoreType")

		case types.TapSlaveCoreType:
			log.Println("We have a core type: TapSlaveCoreType")

		case types.DcfSlaveCoreType:
			log.Println("We have a core type: DcfSlaveCoreType")

		case types.DcfMasterCoreType:
			log.Println("We have a core type: DcfMasterCoreType")

		case types.RedTsnCoreType:
			log.Println("We have a core type: RedTsnCoreType")

		case types.TsnIicCoreType:
			log.Println("We have a core type: TsnIicCoreType")

		case types.NtpServerCoreType:
			log.Println("We have a core type: NtpServerCoreType")

		case types.NtpClientCoreType:
			log.Println("We have a core type: NtpClientCoreType")

		case types.ClkFrequencyGeneratorCoreType:
			log.Println("We have a core type: ClkFrequencyGeneratorCoreType")

		case types.SynceNodeCoreType:
			log.Println("We have a core type: SynceNodeCoreType")

		case types.PpsClkToPpsCoreType:
			log.Println("We have a core type: PpsClkToPpsCoreType")

		case types.PtpServerCoreType:
			log.Println("We have a core type: PtpServerCoreType")

		case types.PtpClientCoreType:
			log.Println("We have a core type: PtpClientCoreType")

		case types.PhyConfigurationCoreType:
			log.Println("We have a core type: PhyConfigurationCoreType")

		case types.I2cConfigurationCoreType:
			log.Println("We have a core type: I2cConfigurationCoreType")

		case types.IoConfigurationCoreType:
			log.Println("We have a core type: IoConfigurationCoreType")

		case types.EthernetTestplatformType:
			log.Println("We have a core type: EthernetTestplatformType")

		case types.MinSwitchCoreType:
			log.Println("We have a core type: MinSwitchCoreType")

		case types.ConfExtCoreType:
			log.Println("We have a core type: ConfExtCoreType")
	*/
	default:
		log.Fatal("Core not found error, type: ", core.CoreType)
	}

}
