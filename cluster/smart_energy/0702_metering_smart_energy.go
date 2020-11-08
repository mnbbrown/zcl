package smart_energy

import "hemtjan.st/zcl"

// MeteringSmartEnergy
const MeteringSmartEnergyID zcl.ClusterID = 1794

var MeteringSmartEnergyCluster = zcl.Cluster{
	Name:      "Metering (Smart Energy)",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CurrentSummationDeliveredAttr:     func() zcl.Attr { return new(CurrentSummationDelivered) },
		CurrentSummationReceivedAttr:      func() zcl.Attr { return new(CurrentSummationReceived) },
		CurrentMaxDemandDeliveredAttr:     func() zcl.Attr { return new(CurrentMaxDemandDelivered) },
		CurrentMaxDemandReceivedAttr:      func() zcl.Attr { return new(CurrentMaxDemandReceived) },
		DftSummationAttr:                  func() zcl.Attr { return new(DftSummation) },
		DailyFreezeTimeAttr:               func() zcl.Attr { return new(DailyFreezeTime) },
		PowerFactorAttr:                   func() zcl.Attr { return new(PowerFactor) },
		ReadingSnapShotTimeAttr:           func() zcl.Attr { return new(ReadingSnapShotTime) },
		CurrentMaxDemandDeliveredTimeAttr: func() zcl.Attr { return new(CurrentMaxDemandDeliveredTime) },
		CurrentMaxDemandReceivedTimeAttr:  func() zcl.Attr { return new(CurrentMaxDemandReceivedTime) },
		DefaultUpdatePeriodAttr:           func() zcl.Attr { return new(DefaultUpdatePeriod) },
		FastPollUpdatePeriodAttr:          func() zcl.Attr { return new(FastPollUpdatePeriod) },
		InstantaneousDemandAttr:           func() zcl.Attr { return new(InstantaneousDemand) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
