package main

type Clock struct {
	ControlReg               int64
	StatusReg                int64
	SelectReg                int64
	VersionReg               int64
	TimeValueLReg            int64
	TimeValueHReg            int64
	TimeAdjValueLReg         int64
	TimeAdjValueHReg         int64
	OffsetAdjValueReg        int64
	OffsetAdjIntervalReg     int64
	DriftAdjValueReg         int64
	DriftAdjIntervalReg      int64
	InSyncThresholdReg       int64
	ServoOffsetFactorPReg    int64
	ServoOffsetFactorIReg    int64
	ServoDriftFactorPReg     int64
	ServoDriftFactorIReg     int64
	StatusOffsetReg          int64
	StatusDriftReg           int64
	StatusOffsetFractionsReg int64
	StatusDriftFractionsReg  int64
}
