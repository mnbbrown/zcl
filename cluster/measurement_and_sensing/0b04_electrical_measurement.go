// Provides a mechanism for querying data about the electrical properties as measured by the device.
package measurement_and_sensing

import (
	"neotor.se/zcl/cluster/zcl"
)

// ElectricalMeasurement
// Provides a mechanism for querying data about the electrical properties as measured by the device.

func NewElectricalMeasurementServer(profile zcl.ProfileID) *ElectricalMeasurementServer {
	return &ElectricalMeasurementServer{p: profile}
}
func NewElectricalMeasurementClient(profile zcl.ProfileID) *ElectricalMeasurementClient {
	return &ElectricalMeasurementClient{p: profile}
}

const ElectricalMeasurementCluster zcl.ClusterID = 2820

type ElectricalMeasurementServer struct {
	p zcl.ProfileID

	MeasurementType                          *MeasurementType
	DcVoltage                                *DcVoltage
	DcVoltageMin                             *DcVoltageMin
	DcVoltageMax                             *DcVoltageMax
	DcCurrent                                *DcCurrent
	DcCurrentMin                             *DcCurrentMin
	DcCurrentMax                             *DcCurrentMax
	DcPower                                  *DcPower
	DcPowerMin                               *DcPowerMin
	DcPowerMax                               *DcPowerMax
	DcVoltageMultiplier                      *DcVoltageMultiplier
	DcVoltageDivisor                         *DcVoltageDivisor
	DcCurrentMultiplier                      *DcCurrentMultiplier
	DcCurrentDivisor                         *DcCurrentDivisor
	DcPowerMultiplier                        *DcPowerMultiplier
	DcPowerDivisor                           *DcPowerDivisor
	AcFrequency                              *AcFrequency
	AcFrequencyMin                           *AcFrequencyMin
	AcFrequencyMax                           *AcFrequencyMax
	NeutralCurrent                           *NeutralCurrent
	TotalActivePower                         *TotalActivePower
	TotalReactivePower                       *TotalReactivePower
	TotalApparentPower                       *TotalApparentPower
	Measured1StHarmonicCurrent               *Measured1StHarmonicCurrent
	Measured3RdHarmonicCurrent               *Measured3RdHarmonicCurrent
	Measured5ThHarmonicCurrent               *Measured5ThHarmonicCurrent
	Measured7ThHarmonicCurrent               *Measured7ThHarmonicCurrent
	Measured9ThHarmonicCurrent               *Measured9ThHarmonicCurrent
	Measured11ThHarmonicCurrent              *Measured11ThHarmonicCurrent
	MeasuredPhase1StHarmonicCurrent          *MeasuredPhase1StHarmonicCurrent
	MeasuredPhase3RdHarmonicCurrent          *MeasuredPhase3RdHarmonicCurrent
	MeasuredPhase5ThHarmonicCurrent          *MeasuredPhase5ThHarmonicCurrent
	MeasuredPhase7ThHarmonicCurrent          *MeasuredPhase7ThHarmonicCurrent
	MeasuredPhase9ThHarmonicCurrent          *MeasuredPhase9ThHarmonicCurrent
	MeasuredPhase11ThHarmonicCurrent         *MeasuredPhase11ThHarmonicCurrent
	AcFrequencyMultiplier                    *AcFrequencyMultiplier
	AcFrequencyDivisor                       *AcFrequencyDivisor
	PowerMultiplier                          *PowerMultiplier
	PowerDivisor                             *PowerDivisor
	HarmonicCurrentMultiplier                *HarmonicCurrentMultiplier
	PhaseHarmonicCurrentMultiplier           *PhaseHarmonicCurrentMultiplier
	LineCurrent                              *LineCurrent
	ActiveCurrent                            *ActiveCurrent
	ReactiveCurrent                          *ReactiveCurrent
	RmsVoltage                               *RmsVoltage
	RmsVoltageMin                            *RmsVoltageMin
	RmsVoltageMax                            *RmsVoltageMax
	RmsCurrent                               *RmsCurrent
	RmsCurrentMin                            *RmsCurrentMin
	RmsCurrentMax                            *RmsCurrentMax
	ActivePower                              *ActivePower
	ActivePowerMin                           *ActivePowerMin
	ActivePowerMax                           *ActivePowerMax
	ReactivePower                            *ReactivePower
	ApparentPower                            *ApparentPower
	PowerFactor                              *PowerFactor
	AverageRmsVoltageMeasurementPeriod       *AverageRmsVoltageMeasurementPeriod
	AverageRmsOvervoltageCounter             *AverageRmsOvervoltageCounter
	AverageRmsUndervoltageCounter            *AverageRmsUndervoltageCounter
	RmsExtremeOvervoltagePeriod              *RmsExtremeOvervoltagePeriod
	RmsExtremeUndervoltagePeriod             *RmsExtremeUndervoltagePeriod
	RmsVoltageSagPeriod                      *RmsVoltageSagPeriod
	RmsVoltageSwellPeriod                    *RmsVoltageSwellPeriod
	AcVoltageMultiplier                      *AcVoltageMultiplier
	AcVoltageDivisor                         *AcVoltageDivisor
	AcCurrentMultiplier                      *AcCurrentMultiplier
	AcCurrentDivisor                         *AcCurrentDivisor
	AcPowerMultiplier                        *AcPowerMultiplier
	AcPowerDivisor                           *AcPowerDivisor
	DcOverloadAlarmsMask                     *DcOverloadAlarmsMask
	DcVoltageOverload                        *DcVoltageOverload
	DcCurrentOverload                        *DcCurrentOverload
	AcOverloadAlarmsMask                     *AcOverloadAlarmsMask
	AcVoltageOverload                        *AcVoltageOverload
	AcCurrentOverload                        *AcCurrentOverload
	AcActivePowerOverload                    *AcActivePowerOverload
	AcReactivePowerOverload                  *AcReactivePowerOverload
	AverageRmsOvervoltage                    *AverageRmsOvervoltage
	AverageRmsUndervoltage                   *AverageRmsUndervoltage
	RmsExtremeOvervoltage                    *RmsExtremeOvervoltage
	RmsExtremeUndervoltage                   *RmsExtremeUndervoltage
	RmsVoltageSag                            *RmsVoltageSag
	RmsVoltageSwell                          *RmsVoltageSwell
	LineCurrentPhaseB                        *LineCurrentPhaseB
	ActiveCurrentPhaseB                      *ActiveCurrentPhaseB
	ReactiveCurrentPhaseB                    *ReactiveCurrentPhaseB
	RmsVoltagePhaseB                         *RmsVoltagePhaseB
	RmsVoltageMinPhaseB                      *RmsVoltageMinPhaseB
	RmsVoltageMaxPhaseB                      *RmsVoltageMaxPhaseB
	RmsCurrentPhaseB                         *RmsCurrentPhaseB
	RmsCurrentMinPhaseB                      *RmsCurrentMinPhaseB
	RmsCurrentMaxPhaseB                      *RmsCurrentMaxPhaseB
	ActivePowerPhaseB                        *ActivePowerPhaseB
	ActivePowerMinPhaseB                     *ActivePowerMinPhaseB
	ActivePowerMaxPhaseB                     *ActivePowerMaxPhaseB
	ReactivePowerPhaseB                      *ReactivePowerPhaseB
	ApparentPowerPhaseB                      *ApparentPowerPhaseB
	PowerFactorPhaseB                        *PowerFactorPhaseB
	AverageRmsVoltageMeasurementPeriodPhaseB *AverageRmsVoltageMeasurementPeriodPhaseB
	AverageRmsOvervoltageCounterPhaseB       *AverageRmsOvervoltageCounterPhaseB
	AverageRmsUndervoltageCounterPhaseB      *AverageRmsUndervoltageCounterPhaseB
	RmsExtremeOvervoltagePeriodPhaseB        *RmsExtremeOvervoltagePeriodPhaseB
	RmsExtremeUndervoltagePeriodPhaseB       *RmsExtremeUndervoltagePeriodPhaseB
	RmsVoltageSagPeriodPhaseB                *RmsVoltageSagPeriodPhaseB
	RmsVoltageSwellPeriodPhaseB              *RmsVoltageSwellPeriodPhaseB
	LineCurrentPhaseC                        *LineCurrentPhaseC
	ActiveCurrentPhaseC                      *ActiveCurrentPhaseC
	ReactiveCurrentPhaseC                    *ReactiveCurrentPhaseC
	RmsVoltagePhaseC                         *RmsVoltagePhaseC
	RmsVoltageMinPhaseC                      *RmsVoltageMinPhaseC
	RmsVoltageMaxPhaseC                      *RmsVoltageMaxPhaseC
	RmsCurrentPhaseC                         *RmsCurrentPhaseC
	RmsCurrentMinPhaseC                      *RmsCurrentMinPhaseC
	RmsCurrentMaxPhaseC                      *RmsCurrentMaxPhaseC
	ActivePowerPhaseC                        *ActivePowerPhaseC
	ActivePowerMinPhaseC                     *ActivePowerMinPhaseC
	ActivePowerMaxPhaseC                     *ActivePowerMaxPhaseC
	ReactivePowerPhaseC                      *ReactivePowerPhaseC
	ApparentPowerPhaseC                      *ApparentPowerPhaseC
	PowerFactorPhaseC                        *PowerFactorPhaseC
	AverageRmsVoltageMeasurementPeriodPhaseC *AverageRmsVoltageMeasurementPeriodPhaseC
	AverageRmsOvervoltageCounterPhaseC       *AverageRmsOvervoltageCounterPhaseC
	AverageRmsUndervoltageCounterPhaseC      *AverageRmsUndervoltageCounterPhaseC
	RmsExtremeOvervoltagePeriodPhaseC        *RmsExtremeOvervoltagePeriodPhaseC
	RmsExtremeUndervoltagePeriodPhaseC       *RmsExtremeUndervoltagePeriodPhaseC
	RmsVoltageSagPeriodPhaseC                *RmsVoltageSagPeriodPhaseC
	RmsVoltageSwellPeriodPhaseC              *RmsVoltageSwellPeriodPhaseC
}

func (s *ElectricalMeasurementServer) GetProfileInfoResponseCommand() *GetProfileInfoResponseCommand {
	return new(GetProfileInfoResponseCommand)
}
func (s *ElectricalMeasurementServer) GetMeasurementProfileResponseCommand() *GetMeasurementProfileResponseCommand {
	return new(GetMeasurementProfileResponseCommand)
}

type ElectricalMeasurementClient struct {
	p zcl.ProfileID
}

func (s *ElectricalMeasurementClient) GetProfileInfoCommand() *GetProfileInfoCommand {
	return new(GetProfileInfoCommand)
}
func (s *ElectricalMeasurementClient) GetMeasurementProfileCommand() *GetMeasurementProfileCommand {
	return new(GetMeasurementProfileCommand)
}

/*
var ElectricalMeasurementServer = map[zcl.CommandID]func() zcl.Command{
    GetProfileInfoResponseCommandID: func() zcl.Command { return new(GetProfileInfoResponseCommand) },
    GetMeasurementProfileResponseCommandID: func() zcl.Command { return new(GetMeasurementProfileResponseCommand) },
}

var ElectricalMeasurementClient = map[zcl.CommandID]func() zcl.Command{
    GetProfileInfoCommandID: func() zcl.Command { return new(GetProfileInfoCommand) },
    GetMeasurementProfileCommandID: func() zcl.Command { return new(GetMeasurementProfileCommand) },
}
*/

type GetProfileInfoResponseCommand struct {
	ProfileCount          zcl.Zu8
	ProfileIntervalPeriod zcl.Zenum8
	MaxNumberOfIntervals  zcl.Zu8
	ListOfAttributes      zcl.Zarray
}

const GetProfileInfoResponseCommandCommand zcl.CommandID = 0

func (v *GetProfileInfoResponseCommand) Values() []zcl.Val {
	return []zcl.Val{
		&v.ProfileCount,
		&v.ProfileIntervalPeriod,
		&v.MaxNumberOfIntervals,
		&v.ListOfAttributes,
	}
}

func (v GetProfileInfoResponseCommand) ID() zcl.CommandID {
	return GetProfileInfoResponseCommandCommand
}

func (v GetProfileInfoResponseCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}

func (v GetProfileInfoResponseCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetProfileInfoResponseCommand) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ProfileCount.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ProfileIntervalPeriod.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.MaxNumberOfIntervals.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ListOfAttributes.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetProfileInfoResponseCommand) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ProfileCount).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileIntervalPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxNumberOfIntervals).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ListOfAttributes).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type GetMeasurementProfileResponseCommand struct {
	StartTime                  zcl.Zutc
	Status                     zcl.Zenum8
	ProfileIntervalPeriod      zcl.Zenum8
	NumberOfIntervalsDelivered zcl.Zu8
	AttributeId                zcl.Zu16
	Intervals                  zcl.Zarray
}

const GetMeasurementProfileResponseCommandCommand zcl.CommandID = 1

func (v *GetMeasurementProfileResponseCommand) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartTime,
		&v.Status,
		&v.ProfileIntervalPeriod,
		&v.NumberOfIntervalsDelivered,
		&v.AttributeId,
		&v.Intervals,
	}
}

func (v GetMeasurementProfileResponseCommand) ID() zcl.CommandID {
	return GetMeasurementProfileResponseCommandCommand
}

func (v GetMeasurementProfileResponseCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}

func (v GetMeasurementProfileResponseCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetMeasurementProfileResponseCommand) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StartTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ProfileIntervalPeriod.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberOfIntervalsDelivered.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.AttributeId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Intervals.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetMeasurementProfileResponseCommand) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StartTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileIntervalPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberOfIntervalsDelivered).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.AttributeId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Intervals).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type GetProfileInfoCommand struct {
}

const GetProfileInfoCommandCommand zcl.CommandID = 0

func (v *GetProfileInfoCommand) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v GetProfileInfoCommand) ID() zcl.CommandID {
	return GetProfileInfoCommandCommand
}

func (v GetProfileInfoCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}

func (v GetProfileInfoCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetProfileInfoCommand) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *GetProfileInfoCommand) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

type GetMeasurementProfileCommand struct {
	AttributeId       zcl.Zu16
	StartTime         zcl.Zutc
	NumberOfIntervals zcl.Zu8
}

const GetMeasurementProfileCommandCommand zcl.CommandID = 1

func (v *GetMeasurementProfileCommand) Values() []zcl.Val {
	return []zcl.Val{
		&v.AttributeId,
		&v.StartTime,
		&v.NumberOfIntervals,
	}
}

func (v GetMeasurementProfileCommand) ID() zcl.CommandID {
	return GetMeasurementProfileCommandCommand
}

func (v GetMeasurementProfileCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}

func (v GetMeasurementProfileCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetMeasurementProfileCommand) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.AttributeId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StartTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberOfIntervals.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetMeasurementProfileCommand) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.AttributeId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberOfIntervals).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type MeasurementType zcl.Zbmp32

func (a MeasurementType) ID() zcl.AttrID           { return 0 }
func (a MeasurementType) Cluster() zcl.ClusterID   { return ElectricalMeasurementCluster }
func (a *MeasurementType) Value() *MeasurementType { return a }
func (a MeasurementType) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp32(a).MarshalZcl()
}
func (a *MeasurementType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasurementType(*nt)
	return br, err
}

func (a MeasurementType) Readable() bool   { return true }
func (a MeasurementType) Writable() bool   { return false }
func (a MeasurementType) Reportable() bool { return false }
func (a MeasurementType) SceneIndex() int  { return -1 }

func (a MeasurementType) String() string {

	var bstr []string
	if a.IsActiveMeasurementAc() {
		bstr = append(bstr, "Active measurement (AC)")
	}
	if a.IsReactiveMeasurementAc() {
		bstr = append(bstr, "Reactive measurement (AC)")
	}
	if a.IsApparentMeasurementAc() {
		bstr = append(bstr, "Apparent measurement (AC)")
	}
	if a.IsPhaseAMeasurement() {
		bstr = append(bstr, "Phase A measurement")
	}
	if a.IsPhaseBMeasurement() {
		bstr = append(bstr, "Phase B measurement")
	}
	if a.IsPhaseCMeasurement() {
		bstr = append(bstr, "Phase C measurement")
	}
	if a.IsDcMeasurement() {
		bstr = append(bstr, "DC measurement")
	}
	if a.IsHarmonicsMeasurement() {
		bstr = append(bstr, "Harmonics measurement")
	}
	if a.IsPowerQualityMeasurement() {
		bstr = append(bstr, "Power quality measurement")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a MeasurementType) IsActiveMeasurementAc() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *MeasurementType) SetActiveMeasurementAc(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a MeasurementType) IsReactiveMeasurementAc() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *MeasurementType) SetReactiveMeasurementAc(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a MeasurementType) IsApparentMeasurementAc() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *MeasurementType) SetApparentMeasurementAc(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a MeasurementType) IsPhaseAMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *MeasurementType) SetPhaseAMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a MeasurementType) IsPhaseBMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *MeasurementType) SetPhaseBMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a MeasurementType) IsPhaseCMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *MeasurementType) SetPhaseCMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a MeasurementType) IsDcMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *MeasurementType) SetDcMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a MeasurementType) IsHarmonicsMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *MeasurementType) SetHarmonicsMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 7, b))
}

func (a MeasurementType) IsPowerQualityMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *MeasurementType) SetPowerQualityMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 8, b))
}

type DcVoltage zcl.Zs16

func (a DcVoltage) ID() zcl.AttrID         { return 256 }
func (a DcVoltage) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcVoltage) Value() *DcVoltage     { return a }
func (a DcVoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltage(*nt)
	return br, err
}

func (a DcVoltage) Readable() bool   { return true }
func (a DcVoltage) Writable() bool   { return false }
func (a DcVoltage) Reportable() bool { return false }
func (a DcVoltage) SceneIndex() int  { return -1 }

func (a DcVoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcVoltageMin zcl.Zs16

func (a DcVoltageMin) ID() zcl.AttrID         { return 257 }
func (a DcVoltageMin) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcVoltageMin) Value() *DcVoltageMin  { return a }
func (a DcVoltageMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMin(*nt)
	return br, err
}

func (a DcVoltageMin) Readable() bool   { return true }
func (a DcVoltageMin) Writable() bool   { return false }
func (a DcVoltageMin) Reportable() bool { return false }
func (a DcVoltageMin) SceneIndex() int  { return -1 }

func (a DcVoltageMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcVoltageMax zcl.Zs16

func (a DcVoltageMax) ID() zcl.AttrID         { return 258 }
func (a DcVoltageMax) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcVoltageMax) Value() *DcVoltageMax  { return a }
func (a DcVoltageMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMax(*nt)
	return br, err
}

func (a DcVoltageMax) Readable() bool   { return true }
func (a DcVoltageMax) Writable() bool   { return false }
func (a DcVoltageMax) Reportable() bool { return false }
func (a DcVoltageMax) SceneIndex() int  { return -1 }

func (a DcVoltageMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcCurrent zcl.Zs16

func (a DcCurrent) ID() zcl.AttrID         { return 259 }
func (a DcCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcCurrent) Value() *DcCurrent     { return a }
func (a DcCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrent(*nt)
	return br, err
}

func (a DcCurrent) Readable() bool   { return true }
func (a DcCurrent) Writable() bool   { return false }
func (a DcCurrent) Reportable() bool { return false }
func (a DcCurrent) SceneIndex() int  { return -1 }

func (a DcCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcCurrentMin zcl.Zs16

func (a DcCurrentMin) ID() zcl.AttrID         { return 260 }
func (a DcCurrentMin) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcCurrentMin) Value() *DcCurrentMin  { return a }
func (a DcCurrentMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMin(*nt)
	return br, err
}

func (a DcCurrentMin) Readable() bool   { return true }
func (a DcCurrentMin) Writable() bool   { return false }
func (a DcCurrentMin) Reportable() bool { return false }
func (a DcCurrentMin) SceneIndex() int  { return -1 }

func (a DcCurrentMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcCurrentMax zcl.Zs16

func (a DcCurrentMax) ID() zcl.AttrID         { return 261 }
func (a DcCurrentMax) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcCurrentMax) Value() *DcCurrentMax  { return a }
func (a DcCurrentMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMax(*nt)
	return br, err
}

func (a DcCurrentMax) Readable() bool   { return true }
func (a DcCurrentMax) Writable() bool   { return false }
func (a DcCurrentMax) Reportable() bool { return false }
func (a DcCurrentMax) SceneIndex() int  { return -1 }

func (a DcCurrentMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcPower zcl.Zs16

func (a DcPower) ID() zcl.AttrID         { return 262 }
func (a DcPower) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcPower) Value() *DcPower       { return a }
func (a DcPower) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPower(*nt)
	return br, err
}

func (a DcPower) Readable() bool   { return true }
func (a DcPower) Writable() bool   { return false }
func (a DcPower) Reportable() bool { return false }
func (a DcPower) SceneIndex() int  { return -1 }

func (a DcPower) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcPowerMin zcl.Zs16

func (a DcPowerMin) ID() zcl.AttrID         { return 263 }
func (a DcPowerMin) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcPowerMin) Value() *DcPowerMin    { return a }
func (a DcPowerMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcPowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMin(*nt)
	return br, err
}

func (a DcPowerMin) Readable() bool   { return true }
func (a DcPowerMin) Writable() bool   { return false }
func (a DcPowerMin) Reportable() bool { return false }
func (a DcPowerMin) SceneIndex() int  { return -1 }

func (a DcPowerMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcPowerMax zcl.Zs16

func (a DcPowerMax) ID() zcl.AttrID         { return 264 }
func (a DcPowerMax) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *DcPowerMax) Value() *DcPowerMax    { return a }
func (a DcPowerMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcPowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMax(*nt)
	return br, err
}

func (a DcPowerMax) Readable() bool   { return true }
func (a DcPowerMax) Writable() bool   { return false }
func (a DcPowerMax) Reportable() bool { return false }
func (a DcPowerMax) SceneIndex() int  { return -1 }

func (a DcPowerMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcVoltageMultiplier zcl.Zu16

func (a DcVoltageMultiplier) ID() zcl.AttrID               { return 512 }
func (a DcVoltageMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *DcVoltageMultiplier) Value() *DcVoltageMultiplier { return a }
func (a DcVoltageMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMultiplier(*nt)
	return br, err
}

func (a DcVoltageMultiplier) Readable() bool   { return true }
func (a DcVoltageMultiplier) Writable() bool   { return false }
func (a DcVoltageMultiplier) Reportable() bool { return false }
func (a DcVoltageMultiplier) SceneIndex() int  { return -1 }

func (a DcVoltageMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DcVoltageDivisor zcl.Zu16

func (a DcVoltageDivisor) ID() zcl.AttrID            { return 513 }
func (a DcVoltageDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *DcVoltageDivisor) Value() *DcVoltageDivisor { return a }
func (a DcVoltageDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageDivisor(*nt)
	return br, err
}

func (a DcVoltageDivisor) Readable() bool   { return true }
func (a DcVoltageDivisor) Writable() bool   { return false }
func (a DcVoltageDivisor) Reportable() bool { return false }
func (a DcVoltageDivisor) SceneIndex() int  { return -1 }

func (a DcVoltageDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DcCurrentMultiplier zcl.Zu16

func (a DcCurrentMultiplier) ID() zcl.AttrID               { return 514 }
func (a DcCurrentMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *DcCurrentMultiplier) Value() *DcCurrentMultiplier { return a }
func (a DcCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMultiplier(*nt)
	return br, err
}

func (a DcCurrentMultiplier) Readable() bool   { return true }
func (a DcCurrentMultiplier) Writable() bool   { return false }
func (a DcCurrentMultiplier) Reportable() bool { return false }
func (a DcCurrentMultiplier) SceneIndex() int  { return -1 }

func (a DcCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DcCurrentDivisor zcl.Zu16

func (a DcCurrentDivisor) ID() zcl.AttrID            { return 515 }
func (a DcCurrentDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *DcCurrentDivisor) Value() *DcCurrentDivisor { return a }
func (a DcCurrentDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentDivisor(*nt)
	return br, err
}

func (a DcCurrentDivisor) Readable() bool   { return true }
func (a DcCurrentDivisor) Writable() bool   { return false }
func (a DcCurrentDivisor) Reportable() bool { return false }
func (a DcCurrentDivisor) SceneIndex() int  { return -1 }

func (a DcCurrentDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DcPowerMultiplier zcl.Zu16

func (a DcPowerMultiplier) ID() zcl.AttrID             { return 516 }
func (a DcPowerMultiplier) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *DcPowerMultiplier) Value() *DcPowerMultiplier { return a }
func (a DcPowerMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMultiplier(*nt)
	return br, err
}

func (a DcPowerMultiplier) Readable() bool   { return true }
func (a DcPowerMultiplier) Writable() bool   { return false }
func (a DcPowerMultiplier) Reportable() bool { return false }
func (a DcPowerMultiplier) SceneIndex() int  { return -1 }

func (a DcPowerMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DcPowerDivisor zcl.Zu16

func (a DcPowerDivisor) ID() zcl.AttrID          { return 517 }
func (a DcPowerDivisor) Cluster() zcl.ClusterID  { return ElectricalMeasurementCluster }
func (a *DcPowerDivisor) Value() *DcPowerDivisor { return a }
func (a DcPowerDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerDivisor(*nt)
	return br, err
}

func (a DcPowerDivisor) Readable() bool   { return true }
func (a DcPowerDivisor) Writable() bool   { return false }
func (a DcPowerDivisor) Reportable() bool { return false }
func (a DcPowerDivisor) SceneIndex() int  { return -1 }

func (a DcPowerDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcFrequency zcl.Zu16

func (a AcFrequency) ID() zcl.AttrID         { return 768 }
func (a AcFrequency) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *AcFrequency) Value() *AcFrequency   { return a }
func (a AcFrequency) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequency(*nt)
	return br, err
}

func (a AcFrequency) Readable() bool   { return true }
func (a AcFrequency) Writable() bool   { return false }
func (a AcFrequency) Reportable() bool { return false }
func (a AcFrequency) SceneIndex() int  { return -1 }

func (a AcFrequency) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcFrequencyMin zcl.Zu16

func (a AcFrequencyMin) ID() zcl.AttrID          { return 769 }
func (a AcFrequencyMin) Cluster() zcl.ClusterID  { return ElectricalMeasurementCluster }
func (a *AcFrequencyMin) Value() *AcFrequencyMin { return a }
func (a AcFrequencyMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMin(*nt)
	return br, err
}

func (a AcFrequencyMin) Readable() bool   { return true }
func (a AcFrequencyMin) Writable() bool   { return false }
func (a AcFrequencyMin) Reportable() bool { return false }
func (a AcFrequencyMin) SceneIndex() int  { return -1 }

func (a AcFrequencyMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcFrequencyMax zcl.Zu16

func (a AcFrequencyMax) ID() zcl.AttrID          { return 770 }
func (a AcFrequencyMax) Cluster() zcl.ClusterID  { return ElectricalMeasurementCluster }
func (a *AcFrequencyMax) Value() *AcFrequencyMax { return a }
func (a AcFrequencyMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMax(*nt)
	return br, err
}

func (a AcFrequencyMax) Readable() bool   { return true }
func (a AcFrequencyMax) Writable() bool   { return false }
func (a AcFrequencyMax) Reportable() bool { return false }
func (a AcFrequencyMax) SceneIndex() int  { return -1 }

func (a AcFrequencyMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type NeutralCurrent zcl.Zu16

func (a NeutralCurrent) ID() zcl.AttrID          { return 771 }
func (a NeutralCurrent) Cluster() zcl.ClusterID  { return ElectricalMeasurementCluster }
func (a *NeutralCurrent) Value() *NeutralCurrent { return a }
func (a NeutralCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NeutralCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeutralCurrent(*nt)
	return br, err
}

func (a NeutralCurrent) Readable() bool   { return true }
func (a NeutralCurrent) Writable() bool   { return false }
func (a NeutralCurrent) Reportable() bool { return false }
func (a NeutralCurrent) SceneIndex() int  { return -1 }

func (a NeutralCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type TotalActivePower zcl.Zs32

func (a TotalActivePower) ID() zcl.AttrID            { return 772 }
func (a TotalActivePower) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *TotalActivePower) Value() *TotalActivePower { return a }
func (a TotalActivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs32(a).MarshalZcl()
}
func (a *TotalActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalActivePower(*nt)
	return br, err
}

func (a TotalActivePower) Readable() bool   { return true }
func (a TotalActivePower) Writable() bool   { return false }
func (a TotalActivePower) Reportable() bool { return false }
func (a TotalActivePower) SceneIndex() int  { return -1 }

func (a TotalActivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs32(a))
}

type TotalReactivePower zcl.Zs32

func (a TotalReactivePower) ID() zcl.AttrID              { return 773 }
func (a TotalReactivePower) Cluster() zcl.ClusterID      { return ElectricalMeasurementCluster }
func (a *TotalReactivePower) Value() *TotalReactivePower { return a }
func (a TotalReactivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs32(a).MarshalZcl()
}
func (a *TotalReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalReactivePower(*nt)
	return br, err
}

func (a TotalReactivePower) Readable() bool   { return true }
func (a TotalReactivePower) Writable() bool   { return false }
func (a TotalReactivePower) Reportable() bool { return false }
func (a TotalReactivePower) SceneIndex() int  { return -1 }

func (a TotalReactivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs32(a))
}

type TotalApparentPower zcl.Zu32

func (a TotalApparentPower) ID() zcl.AttrID              { return 774 }
func (a TotalApparentPower) Cluster() zcl.ClusterID      { return ElectricalMeasurementCluster }
func (a *TotalApparentPower) Value() *TotalApparentPower { return a }
func (a TotalApparentPower) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *TotalApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalApparentPower(*nt)
	return br, err
}

func (a TotalApparentPower) Readable() bool   { return true }
func (a TotalApparentPower) Writable() bool   { return false }
func (a TotalApparentPower) Reportable() bool { return false }
func (a TotalApparentPower) SceneIndex() int  { return -1 }

func (a TotalApparentPower) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

type Measured1StHarmonicCurrent zcl.Zs16

func (a Measured1StHarmonicCurrent) ID() zcl.AttrID                      { return 775 }
func (a Measured1StHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementCluster }
func (a *Measured1StHarmonicCurrent) Value() *Measured1StHarmonicCurrent { return a }
func (a Measured1StHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured1StHarmonicCurrent(*nt)
	return br, err
}

func (a Measured1StHarmonicCurrent) Readable() bool   { return true }
func (a Measured1StHarmonicCurrent) Writable() bool   { return false }
func (a Measured1StHarmonicCurrent) Reportable() bool { return false }
func (a Measured1StHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured1StHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type Measured3RdHarmonicCurrent zcl.Zs16

func (a Measured3RdHarmonicCurrent) ID() zcl.AttrID                      { return 776 }
func (a Measured3RdHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementCluster }
func (a *Measured3RdHarmonicCurrent) Value() *Measured3RdHarmonicCurrent { return a }
func (a Measured3RdHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured3RdHarmonicCurrent(*nt)
	return br, err
}

func (a Measured3RdHarmonicCurrent) Readable() bool   { return true }
func (a Measured3RdHarmonicCurrent) Writable() bool   { return false }
func (a Measured3RdHarmonicCurrent) Reportable() bool { return false }
func (a Measured3RdHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured3RdHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type Measured5ThHarmonicCurrent zcl.Zs16

func (a Measured5ThHarmonicCurrent) ID() zcl.AttrID                      { return 777 }
func (a Measured5ThHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementCluster }
func (a *Measured5ThHarmonicCurrent) Value() *Measured5ThHarmonicCurrent { return a }
func (a Measured5ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured5ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured5ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured5ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured5ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured5ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured5ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type Measured7ThHarmonicCurrent zcl.Zs16

func (a Measured7ThHarmonicCurrent) ID() zcl.AttrID                      { return 778 }
func (a Measured7ThHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementCluster }
func (a *Measured7ThHarmonicCurrent) Value() *Measured7ThHarmonicCurrent { return a }
func (a Measured7ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured7ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured7ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured7ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured7ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured7ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured7ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type Measured9ThHarmonicCurrent zcl.Zs16

func (a Measured9ThHarmonicCurrent) ID() zcl.AttrID                      { return 779 }
func (a Measured9ThHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementCluster }
func (a *Measured9ThHarmonicCurrent) Value() *Measured9ThHarmonicCurrent { return a }
func (a Measured9ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured9ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured9ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured9ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured9ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured9ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured9ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type Measured11ThHarmonicCurrent zcl.Zs16

func (a Measured11ThHarmonicCurrent) ID() zcl.AttrID                       { return 780 }
func (a Measured11ThHarmonicCurrent) Cluster() zcl.ClusterID               { return ElectricalMeasurementCluster }
func (a *Measured11ThHarmonicCurrent) Value() *Measured11ThHarmonicCurrent { return a }
func (a Measured11ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured11ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured11ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured11ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured11ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured11ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured11ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured11ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MeasuredPhase1StHarmonicCurrent zcl.Zs16

func (a MeasuredPhase1StHarmonicCurrent) ID() zcl.AttrID                           { return 781 }
func (a MeasuredPhase1StHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementCluster }
func (a *MeasuredPhase1StHarmonicCurrent) Value() *MeasuredPhase1StHarmonicCurrent { return a }
func (a MeasuredPhase1StHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase1StHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase1StHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase1StHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase1StHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase1StHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase1StHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MeasuredPhase3RdHarmonicCurrent zcl.Zs16

func (a MeasuredPhase3RdHarmonicCurrent) ID() zcl.AttrID                           { return 782 }
func (a MeasuredPhase3RdHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementCluster }
func (a *MeasuredPhase3RdHarmonicCurrent) Value() *MeasuredPhase3RdHarmonicCurrent { return a }
func (a MeasuredPhase3RdHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase3RdHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase3RdHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase3RdHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase3RdHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase3RdHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase3RdHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MeasuredPhase5ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase5ThHarmonicCurrent) ID() zcl.AttrID                           { return 783 }
func (a MeasuredPhase5ThHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementCluster }
func (a *MeasuredPhase5ThHarmonicCurrent) Value() *MeasuredPhase5ThHarmonicCurrent { return a }
func (a MeasuredPhase5ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase5ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase5ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase5ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase5ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase5ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase5ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MeasuredPhase7ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase7ThHarmonicCurrent) ID() zcl.AttrID                           { return 784 }
func (a MeasuredPhase7ThHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementCluster }
func (a *MeasuredPhase7ThHarmonicCurrent) Value() *MeasuredPhase7ThHarmonicCurrent { return a }
func (a MeasuredPhase7ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase7ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase7ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase7ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase7ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase7ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase7ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MeasuredPhase9ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase9ThHarmonicCurrent) ID() zcl.AttrID                           { return 785 }
func (a MeasuredPhase9ThHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementCluster }
func (a *MeasuredPhase9ThHarmonicCurrent) Value() *MeasuredPhase9ThHarmonicCurrent { return a }
func (a MeasuredPhase9ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase9ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase9ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase9ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase9ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase9ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase9ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MeasuredPhase11ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase11ThHarmonicCurrent) ID() zcl.AttrID                            { return 786 }
func (a MeasuredPhase11ThHarmonicCurrent) Cluster() zcl.ClusterID                    { return ElectricalMeasurementCluster }
func (a *MeasuredPhase11ThHarmonicCurrent) Value() *MeasuredPhase11ThHarmonicCurrent { return a }
func (a MeasuredPhase11ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase11ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase11ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase11ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase11ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase11ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase11ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase11ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type AcFrequencyMultiplier zcl.Zu16

func (a AcFrequencyMultiplier) ID() zcl.AttrID                 { return 1024 }
func (a AcFrequencyMultiplier) Cluster() zcl.ClusterID         { return ElectricalMeasurementCluster }
func (a *AcFrequencyMultiplier) Value() *AcFrequencyMultiplier { return a }
func (a AcFrequencyMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMultiplier(*nt)
	return br, err
}

func (a AcFrequencyMultiplier) Readable() bool   { return true }
func (a AcFrequencyMultiplier) Writable() bool   { return false }
func (a AcFrequencyMultiplier) Reportable() bool { return false }
func (a AcFrequencyMultiplier) SceneIndex() int  { return -1 }

func (a AcFrequencyMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcFrequencyDivisor zcl.Zu16

func (a AcFrequencyDivisor) ID() zcl.AttrID              { return 1025 }
func (a AcFrequencyDivisor) Cluster() zcl.ClusterID      { return ElectricalMeasurementCluster }
func (a *AcFrequencyDivisor) Value() *AcFrequencyDivisor { return a }
func (a AcFrequencyDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyDivisor(*nt)
	return br, err
}

func (a AcFrequencyDivisor) Readable() bool   { return true }
func (a AcFrequencyDivisor) Writable() bool   { return false }
func (a AcFrequencyDivisor) Reportable() bool { return false }
func (a AcFrequencyDivisor) SceneIndex() int  { return -1 }

func (a AcFrequencyDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type PowerMultiplier zcl.Zu32

func (a PowerMultiplier) ID() zcl.AttrID           { return 1026 }
func (a PowerMultiplier) Cluster() zcl.ClusterID   { return ElectricalMeasurementCluster }
func (a *PowerMultiplier) Value() *PowerMultiplier { return a }
func (a PowerMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *PowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerMultiplier(*nt)
	return br, err
}

func (a PowerMultiplier) Readable() bool   { return true }
func (a PowerMultiplier) Writable() bool   { return false }
func (a PowerMultiplier) Reportable() bool { return false }
func (a PowerMultiplier) SceneIndex() int  { return -1 }

func (a PowerMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

type PowerDivisor zcl.Zu32

func (a PowerDivisor) ID() zcl.AttrID         { return 1027 }
func (a PowerDivisor) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *PowerDivisor) Value() *PowerDivisor  { return a }
func (a PowerDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *PowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerDivisor(*nt)
	return br, err
}

func (a PowerDivisor) Readable() bool   { return true }
func (a PowerDivisor) Writable() bool   { return false }
func (a PowerDivisor) Reportable() bool { return false }
func (a PowerDivisor) SceneIndex() int  { return -1 }

func (a PowerDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

type HarmonicCurrentMultiplier zcl.Zs8

func (a HarmonicCurrentMultiplier) ID() zcl.AttrID                     { return 1028 }
func (a HarmonicCurrentMultiplier) Cluster() zcl.ClusterID             { return ElectricalMeasurementCluster }
func (a *HarmonicCurrentMultiplier) Value() *HarmonicCurrentMultiplier { return a }
func (a HarmonicCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *HarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = HarmonicCurrentMultiplier(*nt)
	return br, err
}

func (a HarmonicCurrentMultiplier) Readable() bool   { return true }
func (a HarmonicCurrentMultiplier) Writable() bool   { return false }
func (a HarmonicCurrentMultiplier) Reportable() bool { return false }
func (a HarmonicCurrentMultiplier) SceneIndex() int  { return -1 }

func (a HarmonicCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

type PhaseHarmonicCurrentMultiplier zcl.Zs8

func (a PhaseHarmonicCurrentMultiplier) ID() zcl.AttrID                          { return 1029 }
func (a PhaseHarmonicCurrentMultiplier) Cluster() zcl.ClusterID                  { return ElectricalMeasurementCluster }
func (a *PhaseHarmonicCurrentMultiplier) Value() *PhaseHarmonicCurrentMultiplier { return a }
func (a PhaseHarmonicCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PhaseHarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhaseHarmonicCurrentMultiplier(*nt)
	return br, err
}

func (a PhaseHarmonicCurrentMultiplier) Readable() bool   { return true }
func (a PhaseHarmonicCurrentMultiplier) Writable() bool   { return false }
func (a PhaseHarmonicCurrentMultiplier) Reportable() bool { return false }
func (a PhaseHarmonicCurrentMultiplier) SceneIndex() int  { return -1 }

func (a PhaseHarmonicCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

type LineCurrent zcl.Zu16

func (a LineCurrent) ID() zcl.AttrID         { return 1281 }
func (a LineCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *LineCurrent) Value() *LineCurrent   { return a }
func (a LineCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *LineCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrent(*nt)
	return br, err
}

func (a LineCurrent) Readable() bool   { return true }
func (a LineCurrent) Writable() bool   { return false }
func (a LineCurrent) Reportable() bool { return false }
func (a LineCurrent) SceneIndex() int  { return -1 }

func (a LineCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ActiveCurrent zcl.Zs16

func (a ActiveCurrent) ID() zcl.AttrID         { return 1282 }
func (a ActiveCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *ActiveCurrent) Value() *ActiveCurrent { return a }
func (a ActiveCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrent(*nt)
	return br, err
}

func (a ActiveCurrent) Readable() bool   { return true }
func (a ActiveCurrent) Writable() bool   { return false }
func (a ActiveCurrent) Reportable() bool { return false }
func (a ActiveCurrent) SceneIndex() int  { return -1 }

func (a ActiveCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ReactiveCurrent zcl.Zs16

func (a ReactiveCurrent) ID() zcl.AttrID           { return 1283 }
func (a ReactiveCurrent) Cluster() zcl.ClusterID   { return ElectricalMeasurementCluster }
func (a *ReactiveCurrent) Value() *ReactiveCurrent { return a }
func (a ReactiveCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrent(*nt)
	return br, err
}

func (a ReactiveCurrent) Readable() bool   { return true }
func (a ReactiveCurrent) Writable() bool   { return false }
func (a ReactiveCurrent) Reportable() bool { return false }
func (a ReactiveCurrent) SceneIndex() int  { return -1 }

func (a ReactiveCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type RmsVoltage zcl.Zu16

func (a RmsVoltage) ID() zcl.AttrID         { return 1285 }
func (a RmsVoltage) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *RmsVoltage) Value() *RmsVoltage    { return a }
func (a RmsVoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltage(*nt)
	return br, err
}

func (a RmsVoltage) Readable() bool   { return true }
func (a RmsVoltage) Writable() bool   { return false }
func (a RmsVoltage) Reportable() bool { return false }
func (a RmsVoltage) SceneIndex() int  { return -1 }

func (a RmsVoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageMin zcl.Zu16

func (a RmsVoltageMin) ID() zcl.AttrID         { return 1286 }
func (a RmsVoltageMin) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *RmsVoltageMin) Value() *RmsVoltageMin { return a }
func (a RmsVoltageMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMin(*nt)
	return br, err
}

func (a RmsVoltageMin) Readable() bool   { return true }
func (a RmsVoltageMin) Writable() bool   { return false }
func (a RmsVoltageMin) Reportable() bool { return false }
func (a RmsVoltageMin) SceneIndex() int  { return -1 }

func (a RmsVoltageMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageMax zcl.Zu16

func (a RmsVoltageMax) ID() zcl.AttrID         { return 1287 }
func (a RmsVoltageMax) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *RmsVoltageMax) Value() *RmsVoltageMax { return a }
func (a RmsVoltageMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMax(*nt)
	return br, err
}

func (a RmsVoltageMax) Readable() bool   { return true }
func (a RmsVoltageMax) Writable() bool   { return false }
func (a RmsVoltageMax) Reportable() bool { return false }
func (a RmsVoltageMax) SceneIndex() int  { return -1 }

func (a RmsVoltageMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrent zcl.Zu16

func (a RmsCurrent) ID() zcl.AttrID         { return 1288 }
func (a RmsCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *RmsCurrent) Value() *RmsCurrent    { return a }
func (a RmsCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrent(*nt)
	return br, err
}

func (a RmsCurrent) Readable() bool   { return true }
func (a RmsCurrent) Writable() bool   { return false }
func (a RmsCurrent) Reportable() bool { return false }
func (a RmsCurrent) SceneIndex() int  { return -1 }

func (a RmsCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentMin zcl.Zu16

func (a RmsCurrentMin) ID() zcl.AttrID         { return 1289 }
func (a RmsCurrentMin) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *RmsCurrentMin) Value() *RmsCurrentMin { return a }
func (a RmsCurrentMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMin(*nt)
	return br, err
}

func (a RmsCurrentMin) Readable() bool   { return true }
func (a RmsCurrentMin) Writable() bool   { return false }
func (a RmsCurrentMin) Reportable() bool { return false }
func (a RmsCurrentMin) SceneIndex() int  { return -1 }

func (a RmsCurrentMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentMax zcl.Zu16

func (a RmsCurrentMax) ID() zcl.AttrID         { return 1290 }
func (a RmsCurrentMax) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *RmsCurrentMax) Value() *RmsCurrentMax { return a }
func (a RmsCurrentMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMax(*nt)
	return br, err
}

func (a RmsCurrentMax) Readable() bool   { return true }
func (a RmsCurrentMax) Writable() bool   { return false }
func (a RmsCurrentMax) Reportable() bool { return false }
func (a RmsCurrentMax) SceneIndex() int  { return -1 }

func (a RmsCurrentMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ActivePower zcl.Zs16

func (a ActivePower) ID() zcl.AttrID         { return 1291 }
func (a ActivePower) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *ActivePower) Value() *ActivePower   { return a }
func (a ActivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePower(*nt)
	return br, err
}

func (a ActivePower) Readable() bool   { return true }
func (a ActivePower) Writable() bool   { return false }
func (a ActivePower) Reportable() bool { return false }
func (a ActivePower) SceneIndex() int  { return -1 }

func (a ActivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ActivePowerMin zcl.Zs16

func (a ActivePowerMin) ID() zcl.AttrID          { return 1292 }
func (a ActivePowerMin) Cluster() zcl.ClusterID  { return ElectricalMeasurementCluster }
func (a *ActivePowerMin) Value() *ActivePowerMin { return a }
func (a ActivePowerMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMin(*nt)
	return br, err
}

func (a ActivePowerMin) Readable() bool   { return true }
func (a ActivePowerMin) Writable() bool   { return false }
func (a ActivePowerMin) Reportable() bool { return false }
func (a ActivePowerMin) SceneIndex() int  { return -1 }

func (a ActivePowerMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ActivePowerMax zcl.Zs16

func (a ActivePowerMax) ID() zcl.AttrID          { return 1293 }
func (a ActivePowerMax) Cluster() zcl.ClusterID  { return ElectricalMeasurementCluster }
func (a *ActivePowerMax) Value() *ActivePowerMax { return a }
func (a ActivePowerMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMax(*nt)
	return br, err
}

func (a ActivePowerMax) Readable() bool   { return true }
func (a ActivePowerMax) Writable() bool   { return false }
func (a ActivePowerMax) Reportable() bool { return false }
func (a ActivePowerMax) SceneIndex() int  { return -1 }

func (a ActivePowerMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ReactivePower zcl.Zs16

func (a ReactivePower) ID() zcl.AttrID         { return 1294 }
func (a ReactivePower) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *ReactivePower) Value() *ReactivePower { return a }
func (a ReactivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePower(*nt)
	return br, err
}

func (a ReactivePower) Readable() bool   { return true }
func (a ReactivePower) Writable() bool   { return false }
func (a ReactivePower) Reportable() bool { return false }
func (a ReactivePower) SceneIndex() int  { return -1 }

func (a ReactivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ApparentPower zcl.Zu16

func (a ApparentPower) ID() zcl.AttrID         { return 1295 }
func (a ApparentPower) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *ApparentPower) Value() *ApparentPower { return a }
func (a ApparentPower) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPower(*nt)
	return br, err
}

func (a ApparentPower) Readable() bool   { return true }
func (a ApparentPower) Writable() bool   { return false }
func (a ApparentPower) Reportable() bool { return false }
func (a ApparentPower) SceneIndex() int  { return -1 }

func (a ApparentPower) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type PowerFactor zcl.Zs8

func (a PowerFactor) ID() zcl.AttrID         { return 1296 }
func (a PowerFactor) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *PowerFactor) Value() *PowerFactor   { return a }
func (a PowerFactor) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PowerFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactor(*nt)
	return br, err
}

func (a PowerFactor) Readable() bool   { return true }
func (a PowerFactor) Writable() bool   { return false }
func (a PowerFactor) Reportable() bool { return false }
func (a PowerFactor) SceneIndex() int  { return -1 }

func (a PowerFactor) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

type AverageRmsVoltageMeasurementPeriod zcl.Zu16

func (a AverageRmsVoltageMeasurementPeriod) ID() zcl.AttrID { return 1297 }
func (a AverageRmsVoltageMeasurementPeriod) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *AverageRmsVoltageMeasurementPeriod) Value() *AverageRmsVoltageMeasurementPeriod { return a }
func (a AverageRmsVoltageMeasurementPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsVoltageMeasurementPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsVoltageMeasurementPeriod(*nt)
	return br, err
}

func (a AverageRmsVoltageMeasurementPeriod) Readable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriod) Writable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriod) Reportable() bool { return false }
func (a AverageRmsVoltageMeasurementPeriod) SceneIndex() int  { return -1 }

func (a AverageRmsVoltageMeasurementPeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AverageRmsOvervoltageCounter zcl.Zu16

func (a AverageRmsOvervoltageCounter) ID() zcl.AttrID                        { return 1298 }
func (a AverageRmsOvervoltageCounter) Cluster() zcl.ClusterID                { return ElectricalMeasurementCluster }
func (a *AverageRmsOvervoltageCounter) Value() *AverageRmsOvervoltageCounter { return a }
func (a AverageRmsOvervoltageCounter) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsOvervoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltageCounter(*nt)
	return br, err
}

func (a AverageRmsOvervoltageCounter) Readable() bool   { return true }
func (a AverageRmsOvervoltageCounter) Writable() bool   { return true }
func (a AverageRmsOvervoltageCounter) Reportable() bool { return false }
func (a AverageRmsOvervoltageCounter) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltageCounter) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AverageRmsUndervoltageCounter zcl.Zu16

func (a AverageRmsUndervoltageCounter) ID() zcl.AttrID                         { return 1299 }
func (a AverageRmsUndervoltageCounter) Cluster() zcl.ClusterID                 { return ElectricalMeasurementCluster }
func (a *AverageRmsUndervoltageCounter) Value() *AverageRmsUndervoltageCounter { return a }
func (a AverageRmsUndervoltageCounter) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsUndervoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltageCounter(*nt)
	return br, err
}

func (a AverageRmsUndervoltageCounter) Readable() bool   { return true }
func (a AverageRmsUndervoltageCounter) Writable() bool   { return true }
func (a AverageRmsUndervoltageCounter) Reportable() bool { return false }
func (a AverageRmsUndervoltageCounter) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltageCounter) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsExtremeOvervoltagePeriod zcl.Zu16

func (a RmsExtremeOvervoltagePeriod) ID() zcl.AttrID                       { return 1300 }
func (a RmsExtremeOvervoltagePeriod) Cluster() zcl.ClusterID               { return ElectricalMeasurementCluster }
func (a *RmsExtremeOvervoltagePeriod) Value() *RmsExtremeOvervoltagePeriod { return a }
func (a RmsExtremeOvervoltagePeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeOvervoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltagePeriod(*nt)
	return br, err
}

func (a RmsExtremeOvervoltagePeriod) Readable() bool   { return true }
func (a RmsExtremeOvervoltagePeriod) Writable() bool   { return true }
func (a RmsExtremeOvervoltagePeriod) Reportable() bool { return false }
func (a RmsExtremeOvervoltagePeriod) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltagePeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsExtremeUndervoltagePeriod zcl.Zu16

func (a RmsExtremeUndervoltagePeriod) ID() zcl.AttrID                        { return 1301 }
func (a RmsExtremeUndervoltagePeriod) Cluster() zcl.ClusterID                { return ElectricalMeasurementCluster }
func (a *RmsExtremeUndervoltagePeriod) Value() *RmsExtremeUndervoltagePeriod { return a }
func (a RmsExtremeUndervoltagePeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeUndervoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltagePeriod(*nt)
	return br, err
}

func (a RmsExtremeUndervoltagePeriod) Readable() bool   { return true }
func (a RmsExtremeUndervoltagePeriod) Writable() bool   { return true }
func (a RmsExtremeUndervoltagePeriod) Reportable() bool { return false }
func (a RmsExtremeUndervoltagePeriod) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltagePeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageSagPeriod zcl.Zu16

func (a RmsVoltageSagPeriod) ID() zcl.AttrID               { return 1302 }
func (a RmsVoltageSagPeriod) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsVoltageSagPeriod) Value() *RmsVoltageSagPeriod { return a }
func (a RmsVoltageSagPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSagPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriod(*nt)
	return br, err
}

func (a RmsVoltageSagPeriod) Readable() bool   { return true }
func (a RmsVoltageSagPeriod) Writable() bool   { return true }
func (a RmsVoltageSagPeriod) Reportable() bool { return false }
func (a RmsVoltageSagPeriod) SceneIndex() int  { return -1 }

func (a RmsVoltageSagPeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageSwellPeriod zcl.Zu16

func (a RmsVoltageSwellPeriod) ID() zcl.AttrID                 { return 1303 }
func (a RmsVoltageSwellPeriod) Cluster() zcl.ClusterID         { return ElectricalMeasurementCluster }
func (a *RmsVoltageSwellPeriod) Value() *RmsVoltageSwellPeriod { return a }
func (a RmsVoltageSwellPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSwellPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriod(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriod) Readable() bool   { return true }
func (a RmsVoltageSwellPeriod) Writable() bool   { return true }
func (a RmsVoltageSwellPeriod) Reportable() bool { return false }
func (a RmsVoltageSwellPeriod) SceneIndex() int  { return -1 }

func (a RmsVoltageSwellPeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcVoltageMultiplier zcl.Zu16

func (a AcVoltageMultiplier) ID() zcl.AttrID               { return 1536 }
func (a AcVoltageMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *AcVoltageMultiplier) Value() *AcVoltageMultiplier { return a }
func (a AcVoltageMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageMultiplier(*nt)
	return br, err
}

func (a AcVoltageMultiplier) Readable() bool   { return true }
func (a AcVoltageMultiplier) Writable() bool   { return false }
func (a AcVoltageMultiplier) Reportable() bool { return false }
func (a AcVoltageMultiplier) SceneIndex() int  { return -1 }

func (a AcVoltageMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcVoltageDivisor zcl.Zu16

func (a AcVoltageDivisor) ID() zcl.AttrID            { return 1537 }
func (a AcVoltageDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *AcVoltageDivisor) Value() *AcVoltageDivisor { return a }
func (a AcVoltageDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageDivisor(*nt)
	return br, err
}

func (a AcVoltageDivisor) Readable() bool   { return true }
func (a AcVoltageDivisor) Writable() bool   { return false }
func (a AcVoltageDivisor) Reportable() bool { return false }
func (a AcVoltageDivisor) SceneIndex() int  { return -1 }

func (a AcVoltageDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcCurrentMultiplier zcl.Zu16

func (a AcCurrentMultiplier) ID() zcl.AttrID               { return 1538 }
func (a AcCurrentMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *AcCurrentMultiplier) Value() *AcCurrentMultiplier { return a }
func (a AcCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentMultiplier(*nt)
	return br, err
}

func (a AcCurrentMultiplier) Readable() bool   { return true }
func (a AcCurrentMultiplier) Writable() bool   { return false }
func (a AcCurrentMultiplier) Reportable() bool { return false }
func (a AcCurrentMultiplier) SceneIndex() int  { return -1 }

func (a AcCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcCurrentDivisor zcl.Zu16

func (a AcCurrentDivisor) ID() zcl.AttrID            { return 1539 }
func (a AcCurrentDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *AcCurrentDivisor) Value() *AcCurrentDivisor { return a }
func (a AcCurrentDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentDivisor(*nt)
	return br, err
}

func (a AcCurrentDivisor) Readable() bool   { return true }
func (a AcCurrentDivisor) Writable() bool   { return false }
func (a AcCurrentDivisor) Reportable() bool { return false }
func (a AcCurrentDivisor) SceneIndex() int  { return -1 }

func (a AcCurrentDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcPowerMultiplier zcl.Zu16

func (a AcPowerMultiplier) ID() zcl.AttrID             { return 1540 }
func (a AcPowerMultiplier) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *AcPowerMultiplier) Value() *AcPowerMultiplier { return a }
func (a AcPowerMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcPowerMultiplier(*nt)
	return br, err
}

func (a AcPowerMultiplier) Readable() bool   { return true }
func (a AcPowerMultiplier) Writable() bool   { return false }
func (a AcPowerMultiplier) Reportable() bool { return false }
func (a AcPowerMultiplier) SceneIndex() int  { return -1 }

func (a AcPowerMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AcPowerDivisor zcl.Zu16

func (a AcPowerDivisor) ID() zcl.AttrID          { return 1541 }
func (a AcPowerDivisor) Cluster() zcl.ClusterID  { return ElectricalMeasurementCluster }
func (a *AcPowerDivisor) Value() *AcPowerDivisor { return a }
func (a AcPowerDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcPowerDivisor(*nt)
	return br, err
}

func (a AcPowerDivisor) Readable() bool   { return true }
func (a AcPowerDivisor) Writable() bool   { return false }
func (a AcPowerDivisor) Reportable() bool { return false }
func (a AcPowerDivisor) SceneIndex() int  { return -1 }

func (a AcPowerDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DcOverloadAlarmsMask zcl.Zbmp8

func (a DcOverloadAlarmsMask) ID() zcl.AttrID                { return 1792 }
func (a DcOverloadAlarmsMask) Cluster() zcl.ClusterID        { return ElectricalMeasurementCluster }
func (a *DcOverloadAlarmsMask) Value() *DcOverloadAlarmsMask { return a }
func (a DcOverloadAlarmsMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *DcOverloadAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DcOverloadAlarmsMask(*nt)
	return br, err
}

func (a DcOverloadAlarmsMask) Readable() bool   { return true }
func (a DcOverloadAlarmsMask) Writable() bool   { return true }
func (a DcOverloadAlarmsMask) Reportable() bool { return false }
func (a DcOverloadAlarmsMask) SceneIndex() int  { return -1 }

func (a DcOverloadAlarmsMask) String() string {

	var bstr []string
	if a.IsVoltageOverload() {
		bstr = append(bstr, "Voltage Overload")
	}
	if a.IsCurrentOverload() {
		bstr = append(bstr, "Current Overload")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a DcOverloadAlarmsMask) IsVoltageOverload() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *DcOverloadAlarmsMask) SetVoltageOverload(b bool) {
	*a = DcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a DcOverloadAlarmsMask) IsCurrentOverload() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *DcOverloadAlarmsMask) SetCurrentOverload(b bool) {
	*a = DcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 1, b))
}

type DcVoltageOverload zcl.Zs16

func (a DcVoltageOverload) ID() zcl.AttrID             { return 1793 }
func (a DcVoltageOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *DcVoltageOverload) Value() *DcVoltageOverload { return a }
func (a DcVoltageOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageOverload(*nt)
	return br, err
}

func (a DcVoltageOverload) Readable() bool   { return true }
func (a DcVoltageOverload) Writable() bool   { return false }
func (a DcVoltageOverload) Reportable() bool { return false }
func (a DcVoltageOverload) SceneIndex() int  { return -1 }

func (a DcVoltageOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type DcCurrentOverload zcl.Zs16

func (a DcCurrentOverload) ID() zcl.AttrID             { return 1794 }
func (a DcCurrentOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *DcCurrentOverload) Value() *DcCurrentOverload { return a }
func (a DcCurrentOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentOverload(*nt)
	return br, err
}

func (a DcCurrentOverload) Readable() bool   { return true }
func (a DcCurrentOverload) Writable() bool   { return false }
func (a DcCurrentOverload) Reportable() bool { return false }
func (a DcCurrentOverload) SceneIndex() int  { return -1 }

func (a DcCurrentOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type AcOverloadAlarmsMask zcl.Zbmp16

func (a AcOverloadAlarmsMask) ID() zcl.AttrID                { return 2048 }
func (a AcOverloadAlarmsMask) Cluster() zcl.ClusterID        { return ElectricalMeasurementCluster }
func (a *AcOverloadAlarmsMask) Value() *AcOverloadAlarmsMask { return a }
func (a AcOverloadAlarmsMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *AcOverloadAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcOverloadAlarmsMask(*nt)
	return br, err
}

func (a AcOverloadAlarmsMask) Readable() bool   { return true }
func (a AcOverloadAlarmsMask) Writable() bool   { return true }
func (a AcOverloadAlarmsMask) Reportable() bool { return false }
func (a AcOverloadAlarmsMask) SceneIndex() int  { return -1 }

func (a AcOverloadAlarmsMask) String() string {

	var bstr []string
	if a.IsVoltageOverload() {
		bstr = append(bstr, "Voltage Overload")
	}
	if a.IsCurrentOverload() {
		bstr = append(bstr, "Current Overload")
	}
	if a.IsActivePowerOverload() {
		bstr = append(bstr, "Active Power Overload")
	}
	if a.IsReactivePowerOverload() {
		bstr = append(bstr, "Reactive Power Overload")
	}
	if a.IsAverageRmsOvervoltage() {
		bstr = append(bstr, "Average RMS Overvoltage")
	}
	if a.IsAverageRmsUndervoltage() {
		bstr = append(bstr, "Average RMS Undervoltage")
	}
	if a.IsRmsExtremeOvervoltage() {
		bstr = append(bstr, "RMS Extreme Overvoltage")
	}
	if a.IsRmsExtremeUndervoltage() {
		bstr = append(bstr, "RMS Extreme Undervoltage")
	}
	if a.IsRmsVoltageSag() {
		bstr = append(bstr, "RMS Voltage Sag")
	}
	if a.IsRmsVoltageSwell() {
		bstr = append(bstr, "RMS Voltage Swell")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a AcOverloadAlarmsMask) IsVoltageOverload() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *AcOverloadAlarmsMask) SetVoltageOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a AcOverloadAlarmsMask) IsCurrentOverload() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AcOverloadAlarmsMask) SetCurrentOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AcOverloadAlarmsMask) IsActivePowerOverload() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AcOverloadAlarmsMask) SetActivePowerOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AcOverloadAlarmsMask) IsReactivePowerOverload() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *AcOverloadAlarmsMask) SetReactivePowerOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a AcOverloadAlarmsMask) IsAverageRmsOvervoltage() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *AcOverloadAlarmsMask) SetAverageRmsOvervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a AcOverloadAlarmsMask) IsAverageRmsUndervoltage() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *AcOverloadAlarmsMask) SetAverageRmsUndervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a AcOverloadAlarmsMask) IsRmsExtremeOvervoltage() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *AcOverloadAlarmsMask) SetRmsExtremeOvervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a AcOverloadAlarmsMask) IsRmsExtremeUndervoltage() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *AcOverloadAlarmsMask) SetRmsExtremeUndervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 7, b))
}

func (a AcOverloadAlarmsMask) IsRmsVoltageSag() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *AcOverloadAlarmsMask) SetRmsVoltageSag(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 8, b))
}

func (a AcOverloadAlarmsMask) IsRmsVoltageSwell() bool {
	return zcl.BitmapTest([]byte(a), 9)
}
func (a *AcOverloadAlarmsMask) SetRmsVoltageSwell(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 9, b))
}

type AcVoltageOverload zcl.Zs16

func (a AcVoltageOverload) ID() zcl.AttrID             { return 2049 }
func (a AcVoltageOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *AcVoltageOverload) Value() *AcVoltageOverload { return a }
func (a AcVoltageOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageOverload(*nt)
	return br, err
}

func (a AcVoltageOverload) Readable() bool   { return true }
func (a AcVoltageOverload) Writable() bool   { return false }
func (a AcVoltageOverload) Reportable() bool { return false }
func (a AcVoltageOverload) SceneIndex() int  { return -1 }

func (a AcVoltageOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type AcCurrentOverload zcl.Zs16

func (a AcCurrentOverload) ID() zcl.AttrID             { return 2050 }
func (a AcCurrentOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *AcCurrentOverload) Value() *AcCurrentOverload { return a }
func (a AcCurrentOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentOverload(*nt)
	return br, err
}

func (a AcCurrentOverload) Readable() bool   { return true }
func (a AcCurrentOverload) Writable() bool   { return false }
func (a AcCurrentOverload) Reportable() bool { return false }
func (a AcCurrentOverload) SceneIndex() int  { return -1 }

func (a AcCurrentOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type AcActivePowerOverload zcl.Zs16

func (a AcActivePowerOverload) ID() zcl.AttrID                 { return 2051 }
func (a AcActivePowerOverload) Cluster() zcl.ClusterID         { return ElectricalMeasurementCluster }
func (a *AcActivePowerOverload) Value() *AcActivePowerOverload { return a }
func (a AcActivePowerOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcActivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcActivePowerOverload(*nt)
	return br, err
}

func (a AcActivePowerOverload) Readable() bool   { return true }
func (a AcActivePowerOverload) Writable() bool   { return false }
func (a AcActivePowerOverload) Reportable() bool { return false }
func (a AcActivePowerOverload) SceneIndex() int  { return -1 }

func (a AcActivePowerOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type AcReactivePowerOverload zcl.Zs16

func (a AcReactivePowerOverload) ID() zcl.AttrID                   { return 2052 }
func (a AcReactivePowerOverload) Cluster() zcl.ClusterID           { return ElectricalMeasurementCluster }
func (a *AcReactivePowerOverload) Value() *AcReactivePowerOverload { return a }
func (a AcReactivePowerOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcReactivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcReactivePowerOverload(*nt)
	return br, err
}

func (a AcReactivePowerOverload) Readable() bool   { return true }
func (a AcReactivePowerOverload) Writable() bool   { return false }
func (a AcReactivePowerOverload) Reportable() bool { return false }
func (a AcReactivePowerOverload) SceneIndex() int  { return -1 }

func (a AcReactivePowerOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type AverageRmsOvervoltage zcl.Zs16

func (a AverageRmsOvervoltage) ID() zcl.AttrID                 { return 2053 }
func (a AverageRmsOvervoltage) Cluster() zcl.ClusterID         { return ElectricalMeasurementCluster }
func (a *AverageRmsOvervoltage) Value() *AverageRmsOvervoltage { return a }
func (a AverageRmsOvervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AverageRmsOvervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltage(*nt)
	return br, err
}

func (a AverageRmsOvervoltage) Readable() bool   { return true }
func (a AverageRmsOvervoltage) Writable() bool   { return false }
func (a AverageRmsOvervoltage) Reportable() bool { return false }
func (a AverageRmsOvervoltage) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type AverageRmsUndervoltage zcl.Zs16

func (a AverageRmsUndervoltage) ID() zcl.AttrID                  { return 2054 }
func (a AverageRmsUndervoltage) Cluster() zcl.ClusterID          { return ElectricalMeasurementCluster }
func (a *AverageRmsUndervoltage) Value() *AverageRmsUndervoltage { return a }
func (a AverageRmsUndervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AverageRmsUndervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltage(*nt)
	return br, err
}

func (a AverageRmsUndervoltage) Readable() bool   { return true }
func (a AverageRmsUndervoltage) Writable() bool   { return false }
func (a AverageRmsUndervoltage) Reportable() bool { return false }
func (a AverageRmsUndervoltage) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type RmsExtremeOvervoltage zcl.Zs16

func (a RmsExtremeOvervoltage) ID() zcl.AttrID                 { return 2055 }
func (a RmsExtremeOvervoltage) Cluster() zcl.ClusterID         { return ElectricalMeasurementCluster }
func (a *RmsExtremeOvervoltage) Value() *RmsExtremeOvervoltage { return a }
func (a RmsExtremeOvervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsExtremeOvervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltage(*nt)
	return br, err
}

func (a RmsExtremeOvervoltage) Readable() bool   { return true }
func (a RmsExtremeOvervoltage) Writable() bool   { return false }
func (a RmsExtremeOvervoltage) Reportable() bool { return false }
func (a RmsExtremeOvervoltage) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type RmsExtremeUndervoltage zcl.Zs16

func (a RmsExtremeUndervoltage) ID() zcl.AttrID                  { return 2056 }
func (a RmsExtremeUndervoltage) Cluster() zcl.ClusterID          { return ElectricalMeasurementCluster }
func (a *RmsExtremeUndervoltage) Value() *RmsExtremeUndervoltage { return a }
func (a RmsExtremeUndervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsExtremeUndervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltage(*nt)
	return br, err
}

func (a RmsExtremeUndervoltage) Readable() bool   { return true }
func (a RmsExtremeUndervoltage) Writable() bool   { return false }
func (a RmsExtremeUndervoltage) Reportable() bool { return false }
func (a RmsExtremeUndervoltage) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type RmsVoltageSag zcl.Zs16

func (a RmsVoltageSag) ID() zcl.AttrID         { return 2057 }
func (a RmsVoltageSag) Cluster() zcl.ClusterID { return ElectricalMeasurementCluster }
func (a *RmsVoltageSag) Value() *RmsVoltageSag { return a }
func (a RmsVoltageSag) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsVoltageSag) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSag(*nt)
	return br, err
}

func (a RmsVoltageSag) Readable() bool   { return true }
func (a RmsVoltageSag) Writable() bool   { return false }
func (a RmsVoltageSag) Reportable() bool { return false }
func (a RmsVoltageSag) SceneIndex() int  { return -1 }

func (a RmsVoltageSag) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type RmsVoltageSwell zcl.Zs16

func (a RmsVoltageSwell) ID() zcl.AttrID           { return 2058 }
func (a RmsVoltageSwell) Cluster() zcl.ClusterID   { return ElectricalMeasurementCluster }
func (a *RmsVoltageSwell) Value() *RmsVoltageSwell { return a }
func (a RmsVoltageSwell) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsVoltageSwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwell(*nt)
	return br, err
}

func (a RmsVoltageSwell) Readable() bool   { return true }
func (a RmsVoltageSwell) Writable() bool   { return false }
func (a RmsVoltageSwell) Reportable() bool { return false }
func (a RmsVoltageSwell) SceneIndex() int  { return -1 }

func (a RmsVoltageSwell) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type LineCurrentPhaseB zcl.Zu16

func (a LineCurrentPhaseB) ID() zcl.AttrID             { return 2305 }
func (a LineCurrentPhaseB) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *LineCurrentPhaseB) Value() *LineCurrentPhaseB { return a }
func (a LineCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *LineCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrentPhaseB(*nt)
	return br, err
}

func (a LineCurrentPhaseB) Readable() bool   { return true }
func (a LineCurrentPhaseB) Writable() bool   { return false }
func (a LineCurrentPhaseB) Reportable() bool { return false }
func (a LineCurrentPhaseB) SceneIndex() int  { return -1 }

func (a LineCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ActiveCurrentPhaseB zcl.Zs16

func (a ActiveCurrentPhaseB) ID() zcl.AttrID               { return 2306 }
func (a ActiveCurrentPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *ActiveCurrentPhaseB) Value() *ActiveCurrentPhaseB { return a }
func (a ActiveCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActiveCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrentPhaseB(*nt)
	return br, err
}

func (a ActiveCurrentPhaseB) Readable() bool   { return true }
func (a ActiveCurrentPhaseB) Writable() bool   { return false }
func (a ActiveCurrentPhaseB) Reportable() bool { return false }
func (a ActiveCurrentPhaseB) SceneIndex() int  { return -1 }

func (a ActiveCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ReactiveCurrentPhaseB zcl.Zs16

func (a ReactiveCurrentPhaseB) ID() zcl.AttrID                 { return 2307 }
func (a ReactiveCurrentPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementCluster }
func (a *ReactiveCurrentPhaseB) Value() *ReactiveCurrentPhaseB { return a }
func (a ReactiveCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactiveCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrentPhaseB(*nt)
	return br, err
}

func (a ReactiveCurrentPhaseB) Readable() bool   { return true }
func (a ReactiveCurrentPhaseB) Writable() bool   { return false }
func (a ReactiveCurrentPhaseB) Reportable() bool { return false }
func (a ReactiveCurrentPhaseB) SceneIndex() int  { return -1 }

func (a ReactiveCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type RmsVoltagePhaseB zcl.Zu16

func (a RmsVoltagePhaseB) ID() zcl.AttrID            { return 2309 }
func (a RmsVoltagePhaseB) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *RmsVoltagePhaseB) Value() *RmsVoltagePhaseB { return a }
func (a RmsVoltagePhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltagePhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltagePhaseB(*nt)
	return br, err
}

func (a RmsVoltagePhaseB) Readable() bool   { return true }
func (a RmsVoltagePhaseB) Writable() bool   { return false }
func (a RmsVoltagePhaseB) Reportable() bool { return false }
func (a RmsVoltagePhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltagePhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageMinPhaseB zcl.Zu16

func (a RmsVoltageMinPhaseB) ID() zcl.AttrID               { return 2310 }
func (a RmsVoltageMinPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsVoltageMinPhaseB) Value() *RmsVoltageMinPhaseB { return a }
func (a RmsVoltageMinPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMinPhaseB(*nt)
	return br, err
}

func (a RmsVoltageMinPhaseB) Readable() bool   { return true }
func (a RmsVoltageMinPhaseB) Writable() bool   { return false }
func (a RmsVoltageMinPhaseB) Reportable() bool { return false }
func (a RmsVoltageMinPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageMinPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageMaxPhaseB zcl.Zu16

func (a RmsVoltageMaxPhaseB) ID() zcl.AttrID               { return 2311 }
func (a RmsVoltageMaxPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsVoltageMaxPhaseB) Value() *RmsVoltageMaxPhaseB { return a }
func (a RmsVoltageMaxPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMaxPhaseB(*nt)
	return br, err
}

func (a RmsVoltageMaxPhaseB) Readable() bool   { return true }
func (a RmsVoltageMaxPhaseB) Writable() bool   { return false }
func (a RmsVoltageMaxPhaseB) Reportable() bool { return false }
func (a RmsVoltageMaxPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageMaxPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentPhaseB zcl.Zu16

func (a RmsCurrentPhaseB) ID() zcl.AttrID            { return 2312 }
func (a RmsCurrentPhaseB) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *RmsCurrentPhaseB) Value() *RmsCurrentPhaseB { return a }
func (a RmsCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentPhaseB(*nt)
	return br, err
}

func (a RmsCurrentPhaseB) Readable() bool   { return true }
func (a RmsCurrentPhaseB) Writable() bool   { return false }
func (a RmsCurrentPhaseB) Reportable() bool { return false }
func (a RmsCurrentPhaseB) SceneIndex() int  { return -1 }

func (a RmsCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentMinPhaseB zcl.Zu16

func (a RmsCurrentMinPhaseB) ID() zcl.AttrID               { return 2313 }
func (a RmsCurrentMinPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsCurrentMinPhaseB) Value() *RmsCurrentMinPhaseB { return a }
func (a RmsCurrentMinPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMinPhaseB(*nt)
	return br, err
}

func (a RmsCurrentMinPhaseB) Readable() bool   { return true }
func (a RmsCurrentMinPhaseB) Writable() bool   { return false }
func (a RmsCurrentMinPhaseB) Reportable() bool { return false }
func (a RmsCurrentMinPhaseB) SceneIndex() int  { return -1 }

func (a RmsCurrentMinPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentMaxPhaseB zcl.Zu16

func (a RmsCurrentMaxPhaseB) ID() zcl.AttrID               { return 2314 }
func (a RmsCurrentMaxPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsCurrentMaxPhaseB) Value() *RmsCurrentMaxPhaseB { return a }
func (a RmsCurrentMaxPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMaxPhaseB(*nt)
	return br, err
}

func (a RmsCurrentMaxPhaseB) Readable() bool   { return true }
func (a RmsCurrentMaxPhaseB) Writable() bool   { return false }
func (a RmsCurrentMaxPhaseB) Reportable() bool { return false }
func (a RmsCurrentMaxPhaseB) SceneIndex() int  { return -1 }

func (a RmsCurrentMaxPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ActivePowerPhaseB zcl.Zs16

func (a ActivePowerPhaseB) ID() zcl.AttrID             { return 2315 }
func (a ActivePowerPhaseB) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *ActivePowerPhaseB) Value() *ActivePowerPhaseB { return a }
func (a ActivePowerPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerPhaseB(*nt)
	return br, err
}

func (a ActivePowerPhaseB) Readable() bool   { return true }
func (a ActivePowerPhaseB) Writable() bool   { return false }
func (a ActivePowerPhaseB) Reportable() bool { return false }
func (a ActivePowerPhaseB) SceneIndex() int  { return -1 }

func (a ActivePowerPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ActivePowerMinPhaseB zcl.Zs16

func (a ActivePowerMinPhaseB) ID() zcl.AttrID                { return 2316 }
func (a ActivePowerMinPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementCluster }
func (a *ActivePowerMinPhaseB) Value() *ActivePowerMinPhaseB { return a }
func (a ActivePowerMinPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMinPhaseB(*nt)
	return br, err
}

func (a ActivePowerMinPhaseB) Readable() bool   { return true }
func (a ActivePowerMinPhaseB) Writable() bool   { return false }
func (a ActivePowerMinPhaseB) Reportable() bool { return false }
func (a ActivePowerMinPhaseB) SceneIndex() int  { return -1 }

func (a ActivePowerMinPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ActivePowerMaxPhaseB zcl.Zs16

func (a ActivePowerMaxPhaseB) ID() zcl.AttrID                { return 2317 }
func (a ActivePowerMaxPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementCluster }
func (a *ActivePowerMaxPhaseB) Value() *ActivePowerMaxPhaseB { return a }
func (a ActivePowerMaxPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMaxPhaseB(*nt)
	return br, err
}

func (a ActivePowerMaxPhaseB) Readable() bool   { return true }
func (a ActivePowerMaxPhaseB) Writable() bool   { return false }
func (a ActivePowerMaxPhaseB) Reportable() bool { return false }
func (a ActivePowerMaxPhaseB) SceneIndex() int  { return -1 }

func (a ActivePowerMaxPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ReactivePowerPhaseB zcl.Zs16

func (a ReactivePowerPhaseB) ID() zcl.AttrID               { return 2318 }
func (a ReactivePowerPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *ReactivePowerPhaseB) Value() *ReactivePowerPhaseB { return a }
func (a ReactivePowerPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactivePowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePowerPhaseB(*nt)
	return br, err
}

func (a ReactivePowerPhaseB) Readable() bool   { return true }
func (a ReactivePowerPhaseB) Writable() bool   { return false }
func (a ReactivePowerPhaseB) Reportable() bool { return false }
func (a ReactivePowerPhaseB) SceneIndex() int  { return -1 }

func (a ReactivePowerPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ApparentPowerPhaseB zcl.Zu16

func (a ApparentPowerPhaseB) ID() zcl.AttrID               { return 2319 }
func (a ApparentPowerPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *ApparentPowerPhaseB) Value() *ApparentPowerPhaseB { return a }
func (a ApparentPowerPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApparentPowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPowerPhaseB(*nt)
	return br, err
}

func (a ApparentPowerPhaseB) Readable() bool   { return true }
func (a ApparentPowerPhaseB) Writable() bool   { return false }
func (a ApparentPowerPhaseB) Reportable() bool { return false }
func (a ApparentPowerPhaseB) SceneIndex() int  { return -1 }

func (a ApparentPowerPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type PowerFactorPhaseB zcl.Zs8

func (a PowerFactorPhaseB) ID() zcl.AttrID             { return 2320 }
func (a PowerFactorPhaseB) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *PowerFactorPhaseB) Value() *PowerFactorPhaseB { return a }
func (a PowerFactorPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PowerFactorPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactorPhaseB(*nt)
	return br, err
}

func (a PowerFactorPhaseB) Readable() bool   { return true }
func (a PowerFactorPhaseB) Writable() bool   { return false }
func (a PowerFactorPhaseB) Reportable() bool { return false }
func (a PowerFactorPhaseB) SceneIndex() int  { return -1 }

func (a PowerFactorPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

type AverageRmsVoltageMeasurementPeriodPhaseB zcl.Zu16

func (a AverageRmsVoltageMeasurementPeriodPhaseB) ID() zcl.AttrID { return 2321 }
func (a AverageRmsVoltageMeasurementPeriodPhaseB) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *AverageRmsVoltageMeasurementPeriodPhaseB) Value() *AverageRmsVoltageMeasurementPeriodPhaseB {
	return a
}
func (a AverageRmsVoltageMeasurementPeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsVoltageMeasurementPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsVoltageMeasurementPeriodPhaseB(*nt)
	return br, err
}

func (a AverageRmsVoltageMeasurementPeriodPhaseB) Readable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseB) Writable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseB) Reportable() bool { return false }
func (a AverageRmsVoltageMeasurementPeriodPhaseB) SceneIndex() int  { return -1 }

func (a AverageRmsVoltageMeasurementPeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AverageRmsOvervoltageCounterPhaseB zcl.Zu16

func (a AverageRmsOvervoltageCounterPhaseB) ID() zcl.AttrID { return 2322 }
func (a AverageRmsOvervoltageCounterPhaseB) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *AverageRmsOvervoltageCounterPhaseB) Value() *AverageRmsOvervoltageCounterPhaseB { return a }
func (a AverageRmsOvervoltageCounterPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsOvervoltageCounterPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltageCounterPhaseB(*nt)
	return br, err
}

func (a AverageRmsOvervoltageCounterPhaseB) Readable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseB) Writable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseB) Reportable() bool { return false }
func (a AverageRmsOvervoltageCounterPhaseB) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltageCounterPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AverageRmsUndervoltageCounterPhaseB zcl.Zu16

func (a AverageRmsUndervoltageCounterPhaseB) ID() zcl.AttrID { return 2323 }
func (a AverageRmsUndervoltageCounterPhaseB) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *AverageRmsUndervoltageCounterPhaseB) Value() *AverageRmsUndervoltageCounterPhaseB { return a }
func (a AverageRmsUndervoltageCounterPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsUndervoltageCounterPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltageCounterPhaseB(*nt)
	return br, err
}

func (a AverageRmsUndervoltageCounterPhaseB) Readable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseB) Writable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseB) Reportable() bool { return false }
func (a AverageRmsUndervoltageCounterPhaseB) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltageCounterPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsExtremeOvervoltagePeriodPhaseB zcl.Zu16

func (a RmsExtremeOvervoltagePeriodPhaseB) ID() zcl.AttrID { return 2324 }
func (a RmsExtremeOvervoltagePeriodPhaseB) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *RmsExtremeOvervoltagePeriodPhaseB) Value() *RmsExtremeOvervoltagePeriodPhaseB { return a }
func (a RmsExtremeOvervoltagePeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeOvervoltagePeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltagePeriodPhaseB(*nt)
	return br, err
}

func (a RmsExtremeOvervoltagePeriodPhaseB) Readable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseB) Writable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseB) Reportable() bool { return false }
func (a RmsExtremeOvervoltagePeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltagePeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsExtremeUndervoltagePeriodPhaseB zcl.Zu16

func (a RmsExtremeUndervoltagePeriodPhaseB) ID() zcl.AttrID { return 2325 }
func (a RmsExtremeUndervoltagePeriodPhaseB) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *RmsExtremeUndervoltagePeriodPhaseB) Value() *RmsExtremeUndervoltagePeriodPhaseB { return a }
func (a RmsExtremeUndervoltagePeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeUndervoltagePeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltagePeriodPhaseB(*nt)
	return br, err
}

func (a RmsExtremeUndervoltagePeriodPhaseB) Readable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseB) Writable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseB) Reportable() bool { return false }
func (a RmsExtremeUndervoltagePeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltagePeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageSagPeriodPhaseB zcl.Zu16

func (a RmsVoltageSagPeriodPhaseB) ID() zcl.AttrID                     { return 2326 }
func (a RmsVoltageSagPeriodPhaseB) Cluster() zcl.ClusterID             { return ElectricalMeasurementCluster }
func (a *RmsVoltageSagPeriodPhaseB) Value() *RmsVoltageSagPeriodPhaseB { return a }
func (a RmsVoltageSagPeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSagPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriodPhaseB(*nt)
	return br, err
}

func (a RmsVoltageSagPeriodPhaseB) Readable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseB) Writable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseB) Reportable() bool { return false }
func (a RmsVoltageSagPeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageSagPeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageSwellPeriodPhaseB zcl.Zu16

func (a RmsVoltageSwellPeriodPhaseB) ID() zcl.AttrID                       { return 2327 }
func (a RmsVoltageSwellPeriodPhaseB) Cluster() zcl.ClusterID               { return ElectricalMeasurementCluster }
func (a *RmsVoltageSwellPeriodPhaseB) Value() *RmsVoltageSwellPeriodPhaseB { return a }
func (a RmsVoltageSwellPeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSwellPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriodPhaseB(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriodPhaseB) Readable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseB) Writable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseB) Reportable() bool { return false }
func (a RmsVoltageSwellPeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageSwellPeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type LineCurrentPhaseC zcl.Zu16

func (a LineCurrentPhaseC) ID() zcl.AttrID             { return 2561 }
func (a LineCurrentPhaseC) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *LineCurrentPhaseC) Value() *LineCurrentPhaseC { return a }
func (a LineCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *LineCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrentPhaseC(*nt)
	return br, err
}

func (a LineCurrentPhaseC) Readable() bool   { return true }
func (a LineCurrentPhaseC) Writable() bool   { return false }
func (a LineCurrentPhaseC) Reportable() bool { return false }
func (a LineCurrentPhaseC) SceneIndex() int  { return -1 }

func (a LineCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ActiveCurrentPhaseC zcl.Zs16

func (a ActiveCurrentPhaseC) ID() zcl.AttrID               { return 2562 }
func (a ActiveCurrentPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *ActiveCurrentPhaseC) Value() *ActiveCurrentPhaseC { return a }
func (a ActiveCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActiveCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrentPhaseC(*nt)
	return br, err
}

func (a ActiveCurrentPhaseC) Readable() bool   { return true }
func (a ActiveCurrentPhaseC) Writable() bool   { return false }
func (a ActiveCurrentPhaseC) Reportable() bool { return false }
func (a ActiveCurrentPhaseC) SceneIndex() int  { return -1 }

func (a ActiveCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ReactiveCurrentPhaseC zcl.Zs16

func (a ReactiveCurrentPhaseC) ID() zcl.AttrID                 { return 2563 }
func (a ReactiveCurrentPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementCluster }
func (a *ReactiveCurrentPhaseC) Value() *ReactiveCurrentPhaseC { return a }
func (a ReactiveCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactiveCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrentPhaseC(*nt)
	return br, err
}

func (a ReactiveCurrentPhaseC) Readable() bool   { return true }
func (a ReactiveCurrentPhaseC) Writable() bool   { return false }
func (a ReactiveCurrentPhaseC) Reportable() bool { return false }
func (a ReactiveCurrentPhaseC) SceneIndex() int  { return -1 }

func (a ReactiveCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type RmsVoltagePhaseC zcl.Zu16

func (a RmsVoltagePhaseC) ID() zcl.AttrID            { return 2565 }
func (a RmsVoltagePhaseC) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *RmsVoltagePhaseC) Value() *RmsVoltagePhaseC { return a }
func (a RmsVoltagePhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltagePhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltagePhaseC(*nt)
	return br, err
}

func (a RmsVoltagePhaseC) Readable() bool   { return true }
func (a RmsVoltagePhaseC) Writable() bool   { return false }
func (a RmsVoltagePhaseC) Reportable() bool { return false }
func (a RmsVoltagePhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltagePhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageMinPhaseC zcl.Zu16

func (a RmsVoltageMinPhaseC) ID() zcl.AttrID               { return 2566 }
func (a RmsVoltageMinPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsVoltageMinPhaseC) Value() *RmsVoltageMinPhaseC { return a }
func (a RmsVoltageMinPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMinPhaseC(*nt)
	return br, err
}

func (a RmsVoltageMinPhaseC) Readable() bool   { return true }
func (a RmsVoltageMinPhaseC) Writable() bool   { return false }
func (a RmsVoltageMinPhaseC) Reportable() bool { return false }
func (a RmsVoltageMinPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageMinPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageMaxPhaseC zcl.Zu16

func (a RmsVoltageMaxPhaseC) ID() zcl.AttrID               { return 2567 }
func (a RmsVoltageMaxPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsVoltageMaxPhaseC) Value() *RmsVoltageMaxPhaseC { return a }
func (a RmsVoltageMaxPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMaxPhaseC(*nt)
	return br, err
}

func (a RmsVoltageMaxPhaseC) Readable() bool   { return true }
func (a RmsVoltageMaxPhaseC) Writable() bool   { return false }
func (a RmsVoltageMaxPhaseC) Reportable() bool { return false }
func (a RmsVoltageMaxPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageMaxPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentPhaseC zcl.Zu16

func (a RmsCurrentPhaseC) ID() zcl.AttrID            { return 2568 }
func (a RmsCurrentPhaseC) Cluster() zcl.ClusterID    { return ElectricalMeasurementCluster }
func (a *RmsCurrentPhaseC) Value() *RmsCurrentPhaseC { return a }
func (a RmsCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentPhaseC(*nt)
	return br, err
}

func (a RmsCurrentPhaseC) Readable() bool   { return true }
func (a RmsCurrentPhaseC) Writable() bool   { return false }
func (a RmsCurrentPhaseC) Reportable() bool { return false }
func (a RmsCurrentPhaseC) SceneIndex() int  { return -1 }

func (a RmsCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentMinPhaseC zcl.Zu16

func (a RmsCurrentMinPhaseC) ID() zcl.AttrID               { return 2569 }
func (a RmsCurrentMinPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsCurrentMinPhaseC) Value() *RmsCurrentMinPhaseC { return a }
func (a RmsCurrentMinPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMinPhaseC(*nt)
	return br, err
}

func (a RmsCurrentMinPhaseC) Readable() bool   { return true }
func (a RmsCurrentMinPhaseC) Writable() bool   { return false }
func (a RmsCurrentMinPhaseC) Reportable() bool { return false }
func (a RmsCurrentMinPhaseC) SceneIndex() int  { return -1 }

func (a RmsCurrentMinPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsCurrentMaxPhaseC zcl.Zu16

func (a RmsCurrentMaxPhaseC) ID() zcl.AttrID               { return 2570 }
func (a RmsCurrentMaxPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *RmsCurrentMaxPhaseC) Value() *RmsCurrentMaxPhaseC { return a }
func (a RmsCurrentMaxPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMaxPhaseC(*nt)
	return br, err
}

func (a RmsCurrentMaxPhaseC) Readable() bool   { return true }
func (a RmsCurrentMaxPhaseC) Writable() bool   { return false }
func (a RmsCurrentMaxPhaseC) Reportable() bool { return false }
func (a RmsCurrentMaxPhaseC) SceneIndex() int  { return -1 }

func (a RmsCurrentMaxPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ActivePowerPhaseC zcl.Zs16

func (a ActivePowerPhaseC) ID() zcl.AttrID             { return 2571 }
func (a ActivePowerPhaseC) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *ActivePowerPhaseC) Value() *ActivePowerPhaseC { return a }
func (a ActivePowerPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerPhaseC(*nt)
	return br, err
}

func (a ActivePowerPhaseC) Readable() bool   { return true }
func (a ActivePowerPhaseC) Writable() bool   { return false }
func (a ActivePowerPhaseC) Reportable() bool { return false }
func (a ActivePowerPhaseC) SceneIndex() int  { return -1 }

func (a ActivePowerPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ActivePowerMinPhaseC zcl.Zs16

func (a ActivePowerMinPhaseC) ID() zcl.AttrID                { return 2572 }
func (a ActivePowerMinPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementCluster }
func (a *ActivePowerMinPhaseC) Value() *ActivePowerMinPhaseC { return a }
func (a ActivePowerMinPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMinPhaseC(*nt)
	return br, err
}

func (a ActivePowerMinPhaseC) Readable() bool   { return true }
func (a ActivePowerMinPhaseC) Writable() bool   { return false }
func (a ActivePowerMinPhaseC) Reportable() bool { return false }
func (a ActivePowerMinPhaseC) SceneIndex() int  { return -1 }

func (a ActivePowerMinPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ActivePowerMaxPhaseC zcl.Zs16

func (a ActivePowerMaxPhaseC) ID() zcl.AttrID                { return 2573 }
func (a ActivePowerMaxPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementCluster }
func (a *ActivePowerMaxPhaseC) Value() *ActivePowerMaxPhaseC { return a }
func (a ActivePowerMaxPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMaxPhaseC(*nt)
	return br, err
}

func (a ActivePowerMaxPhaseC) Readable() bool   { return true }
func (a ActivePowerMaxPhaseC) Writable() bool   { return false }
func (a ActivePowerMaxPhaseC) Reportable() bool { return false }
func (a ActivePowerMaxPhaseC) SceneIndex() int  { return -1 }

func (a ActivePowerMaxPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ReactivePowerPhaseC zcl.Zs16

func (a ReactivePowerPhaseC) ID() zcl.AttrID               { return 2574 }
func (a ReactivePowerPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *ReactivePowerPhaseC) Value() *ReactivePowerPhaseC { return a }
func (a ReactivePowerPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactivePowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePowerPhaseC(*nt)
	return br, err
}

func (a ReactivePowerPhaseC) Readable() bool   { return true }
func (a ReactivePowerPhaseC) Writable() bool   { return false }
func (a ReactivePowerPhaseC) Reportable() bool { return false }
func (a ReactivePowerPhaseC) SceneIndex() int  { return -1 }

func (a ReactivePowerPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ApparentPowerPhaseC zcl.Zu16

func (a ApparentPowerPhaseC) ID() zcl.AttrID               { return 2575 }
func (a ApparentPowerPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementCluster }
func (a *ApparentPowerPhaseC) Value() *ApparentPowerPhaseC { return a }
func (a ApparentPowerPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApparentPowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPowerPhaseC(*nt)
	return br, err
}

func (a ApparentPowerPhaseC) Readable() bool   { return true }
func (a ApparentPowerPhaseC) Writable() bool   { return false }
func (a ApparentPowerPhaseC) Reportable() bool { return false }
func (a ApparentPowerPhaseC) SceneIndex() int  { return -1 }

func (a ApparentPowerPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type PowerFactorPhaseC zcl.Zs8

func (a PowerFactorPhaseC) ID() zcl.AttrID             { return 2576 }
func (a PowerFactorPhaseC) Cluster() zcl.ClusterID     { return ElectricalMeasurementCluster }
func (a *PowerFactorPhaseC) Value() *PowerFactorPhaseC { return a }
func (a PowerFactorPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PowerFactorPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactorPhaseC(*nt)
	return br, err
}

func (a PowerFactorPhaseC) Readable() bool   { return true }
func (a PowerFactorPhaseC) Writable() bool   { return false }
func (a PowerFactorPhaseC) Reportable() bool { return false }
func (a PowerFactorPhaseC) SceneIndex() int  { return -1 }

func (a PowerFactorPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

type AverageRmsVoltageMeasurementPeriodPhaseC zcl.Zu16

func (a AverageRmsVoltageMeasurementPeriodPhaseC) ID() zcl.AttrID { return 2577 }
func (a AverageRmsVoltageMeasurementPeriodPhaseC) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *AverageRmsVoltageMeasurementPeriodPhaseC) Value() *AverageRmsVoltageMeasurementPeriodPhaseC {
	return a
}
func (a AverageRmsVoltageMeasurementPeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsVoltageMeasurementPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsVoltageMeasurementPeriodPhaseC(*nt)
	return br, err
}

func (a AverageRmsVoltageMeasurementPeriodPhaseC) Readable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseC) Writable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseC) Reportable() bool { return false }
func (a AverageRmsVoltageMeasurementPeriodPhaseC) SceneIndex() int  { return -1 }

func (a AverageRmsVoltageMeasurementPeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AverageRmsOvervoltageCounterPhaseC zcl.Zu16

func (a AverageRmsOvervoltageCounterPhaseC) ID() zcl.AttrID { return 2578 }
func (a AverageRmsOvervoltageCounterPhaseC) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *AverageRmsOvervoltageCounterPhaseC) Value() *AverageRmsOvervoltageCounterPhaseC { return a }
func (a AverageRmsOvervoltageCounterPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsOvervoltageCounterPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltageCounterPhaseC(*nt)
	return br, err
}

func (a AverageRmsOvervoltageCounterPhaseC) Readable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseC) Writable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseC) Reportable() bool { return false }
func (a AverageRmsOvervoltageCounterPhaseC) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltageCounterPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AverageRmsUndervoltageCounterPhaseC zcl.Zu16

func (a AverageRmsUndervoltageCounterPhaseC) ID() zcl.AttrID { return 2579 }
func (a AverageRmsUndervoltageCounterPhaseC) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *AverageRmsUndervoltageCounterPhaseC) Value() *AverageRmsUndervoltageCounterPhaseC { return a }
func (a AverageRmsUndervoltageCounterPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsUndervoltageCounterPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltageCounterPhaseC(*nt)
	return br, err
}

func (a AverageRmsUndervoltageCounterPhaseC) Readable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseC) Writable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseC) Reportable() bool { return false }
func (a AverageRmsUndervoltageCounterPhaseC) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltageCounterPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsExtremeOvervoltagePeriodPhaseC zcl.Zu16

func (a RmsExtremeOvervoltagePeriodPhaseC) ID() zcl.AttrID { return 2580 }
func (a RmsExtremeOvervoltagePeriodPhaseC) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *RmsExtremeOvervoltagePeriodPhaseC) Value() *RmsExtremeOvervoltagePeriodPhaseC { return a }
func (a RmsExtremeOvervoltagePeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeOvervoltagePeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltagePeriodPhaseC(*nt)
	return br, err
}

func (a RmsExtremeOvervoltagePeriodPhaseC) Readable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseC) Writable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseC) Reportable() bool { return false }
func (a RmsExtremeOvervoltagePeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltagePeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsExtremeUndervoltagePeriodPhaseC zcl.Zu16

func (a RmsExtremeUndervoltagePeriodPhaseC) ID() zcl.AttrID { return 2581 }
func (a RmsExtremeUndervoltagePeriodPhaseC) Cluster() zcl.ClusterID {
	return ElectricalMeasurementCluster
}
func (a *RmsExtremeUndervoltagePeriodPhaseC) Value() *RmsExtremeUndervoltagePeriodPhaseC { return a }
func (a RmsExtremeUndervoltagePeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeUndervoltagePeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltagePeriodPhaseC(*nt)
	return br, err
}

func (a RmsExtremeUndervoltagePeriodPhaseC) Readable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseC) Writable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseC) Reportable() bool { return false }
func (a RmsExtremeUndervoltagePeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltagePeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageSagPeriodPhaseC zcl.Zu16

func (a RmsVoltageSagPeriodPhaseC) ID() zcl.AttrID                     { return 2582 }
func (a RmsVoltageSagPeriodPhaseC) Cluster() zcl.ClusterID             { return ElectricalMeasurementCluster }
func (a *RmsVoltageSagPeriodPhaseC) Value() *RmsVoltageSagPeriodPhaseC { return a }
func (a RmsVoltageSagPeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSagPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriodPhaseC(*nt)
	return br, err
}

func (a RmsVoltageSagPeriodPhaseC) Readable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseC) Writable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseC) Reportable() bool { return false }
func (a RmsVoltageSagPeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageSagPeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RmsVoltageSwellPeriodPhaseC zcl.Zu16

func (a RmsVoltageSwellPeriodPhaseC) ID() zcl.AttrID                       { return 2583 }
func (a RmsVoltageSwellPeriodPhaseC) Cluster() zcl.ClusterID               { return ElectricalMeasurementCluster }
func (a *RmsVoltageSwellPeriodPhaseC) Value() *RmsVoltageSwellPeriodPhaseC { return a }
func (a RmsVoltageSwellPeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSwellPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriodPhaseC(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriodPhaseC) Readable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseC) Writable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseC) Reportable() bool { return false }
func (a RmsVoltageSwellPeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageSwellPeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}