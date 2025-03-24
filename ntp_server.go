package main

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
}
