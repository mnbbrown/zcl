// Attributes for determining information about a device’s internal temperature, and for configuring under/over temperature alarms.
package general

import (
	"neotor.se/zcl"
)

// DeviceTemperatureConfiguration
const DeviceTemperatureConfigurationID zcl.ClusterID = 2

var DeviceTemperatureConfigurationCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CurrentTemperatureAttr:     func() zcl.Attr { return new(CurrentTemperature) },
		MinTempExperiencedAttr:     func() zcl.Attr { return new(MinTempExperienced) },
		MaxTempExperiencedAttr:     func() zcl.Attr { return new(MaxTempExperienced) },
		OverTempTotalDwellAttr:     func() zcl.Attr { return new(OverTempTotalDwell) },
		DeviceTempAlarmMaskAttr:    func() zcl.Attr { return new(DeviceTempAlarmMask) },
		LowTempThresholdAttr:       func() zcl.Attr { return new(LowTempThreshold) },
		HighTempThresholdAttr:      func() zcl.Attr { return new(HighTempThreshold) },
		LowTempDwellTripPointAttr:  func() zcl.Attr { return new(LowTempDwellTripPoint) },
		HighTempDwellTripPointAttr: func() zcl.Attr { return new(HighTempDwellTripPoint) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// CurrentTemperature is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
type CurrentTemperature zcl.Zs16

const CurrentTemperatureAttr zcl.AttrID = 0

func (a CurrentTemperature) ID() zcl.AttrID              { return CurrentTemperatureAttr }
func (a CurrentTemperature) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationID }
func (a *CurrentTemperature) Value() *CurrentTemperature { return a }
func (a CurrentTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *CurrentTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentTemperature(*nt)
	return br, err
}

func (a CurrentTemperature) Readable() bool   { return true }
func (a CurrentTemperature) Writable() bool   { return false }
func (a CurrentTemperature) Reportable() bool { return false }
func (a CurrentTemperature) SceneIndex() int  { return -1 }

func (a CurrentTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

// MinTempExperienced is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
type MinTempExperienced zcl.Zs16

const MinTempExperiencedAttr zcl.AttrID = 1

func (a MinTempExperienced) ID() zcl.AttrID              { return MinTempExperiencedAttr }
func (a MinTempExperienced) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationID }
func (a *MinTempExperienced) Value() *MinTempExperienced { return a }
func (a MinTempExperienced) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MinTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinTempExperienced(*nt)
	return br, err
}

func (a MinTempExperienced) Readable() bool   { return true }
func (a MinTempExperienced) Writable() bool   { return false }
func (a MinTempExperienced) Reportable() bool { return false }
func (a MinTempExperienced) SceneIndex() int  { return -1 }

func (a MinTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

// MaxTempExperienced is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
type MaxTempExperienced zcl.Zs16

const MaxTempExperiencedAttr zcl.AttrID = 2

func (a MaxTempExperienced) ID() zcl.AttrID              { return MaxTempExperiencedAttr }
func (a MaxTempExperienced) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationID }
func (a *MaxTempExperienced) Value() *MaxTempExperienced { return a }
func (a MaxTempExperienced) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MaxTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxTempExperienced(*nt)
	return br, err
}

func (a MaxTempExperienced) Readable() bool   { return true }
func (a MaxTempExperienced) Writable() bool   { return false }
func (a MaxTempExperienced) Reportable() bool { return false }
func (a MaxTempExperienced) SceneIndex() int  { return -1 }

func (a MaxTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

// OverTempTotalDwell is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
// Total time the device has spent above the tmperature specified by High Temp Threshold
type OverTempTotalDwell zcl.Zu16

const OverTempTotalDwellAttr zcl.AttrID = 3

func (a OverTempTotalDwell) ID() zcl.AttrID              { return OverTempTotalDwellAttr }
func (a OverTempTotalDwell) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationID }
func (a *OverTempTotalDwell) Value() *OverTempTotalDwell { return a }
func (a OverTempTotalDwell) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *OverTempTotalDwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OverTempTotalDwell(*nt)
	return br, err
}

func (a OverTempTotalDwell) Readable() bool   { return false }
func (a OverTempTotalDwell) Writable() bool   { return false }
func (a OverTempTotalDwell) Reportable() bool { return false }
func (a OverTempTotalDwell) SceneIndex() int  { return -1 }

func (a OverTempTotalDwell) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// DeviceTempAlarmMask is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
type DeviceTempAlarmMask zcl.Zbmp8

const DeviceTempAlarmMaskAttr zcl.AttrID = 16

func (a DeviceTempAlarmMask) ID() zcl.AttrID               { return DeviceTempAlarmMaskAttr }
func (a DeviceTempAlarmMask) Cluster() zcl.ClusterID       { return DeviceTemperatureConfigurationID }
func (a *DeviceTempAlarmMask) Value() *DeviceTempAlarmMask { return a }
func (a DeviceTempAlarmMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *DeviceTempAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceTempAlarmMask(*nt)
	return br, err
}

func (a DeviceTempAlarmMask) Readable() bool   { return true }
func (a DeviceTempAlarmMask) Writable() bool   { return true }
func (a DeviceTempAlarmMask) Reportable() bool { return false }
func (a DeviceTempAlarmMask) SceneIndex() int  { return -1 }

func (a DeviceTempAlarmMask) String() string {
	var bstr []string
	if a.IsTemperatureTooLow() {
		bstr = append(bstr, "Temperature too low")
	}
	if a.IsTemperatureTooHigh() {
		bstr = append(bstr, "Temperature too high")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a DeviceTempAlarmMask) IsTemperatureTooLow() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *DeviceTempAlarmMask) SetTemperatureTooLow(b bool) {
	*a = DeviceTempAlarmMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a DeviceTempAlarmMask) IsTemperatureTooHigh() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *DeviceTempAlarmMask) SetTemperatureTooHigh(b bool) {
	*a = DeviceTempAlarmMask(zcl.BitmapSet([]byte(*a), 1, b))
}

// LowTempThreshold is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
// If the current temperature drops below the threshold for longer
// than the time specified by low temp dwell, an alarm will be triggered
//
type LowTempThreshold zcl.Zs16

const LowTempThresholdAttr zcl.AttrID = 17

func (a LowTempThreshold) ID() zcl.AttrID            { return LowTempThresholdAttr }
func (a LowTempThreshold) Cluster() zcl.ClusterID    { return DeviceTemperatureConfigurationID }
func (a *LowTempThreshold) Value() *LowTempThreshold { return a }
func (a LowTempThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *LowTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempThreshold(*nt)
	return br, err
}

func (a LowTempThreshold) Readable() bool   { return true }
func (a LowTempThreshold) Writable() bool   { return true }
func (a LowTempThreshold) Reportable() bool { return false }
func (a LowTempThreshold) SceneIndex() int  { return -1 }

func (a LowTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

// HighTempThreshold is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
// If the current temperature goes above the threshold for longer
// than the time specified by high temp dwell, an alarm will be triggered
//
type HighTempThreshold zcl.Zs16

const HighTempThresholdAttr zcl.AttrID = 18

func (a HighTempThreshold) ID() zcl.AttrID             { return HighTempThresholdAttr }
func (a HighTempThreshold) Cluster() zcl.ClusterID     { return DeviceTemperatureConfigurationID }
func (a *HighTempThreshold) Value() *HighTempThreshold { return a }
func (a HighTempThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *HighTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempThreshold(*nt)
	return br, err
}

func (a HighTempThreshold) Readable() bool   { return true }
func (a HighTempThreshold) Writable() bool   { return true }
func (a HighTempThreshold) Reportable() bool { return false }
func (a HighTempThreshold) SceneIndex() int  { return -1 }

func (a HighTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

// LowTempDwellTripPoint is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
type LowTempDwellTripPoint zcl.Zu24

const LowTempDwellTripPointAttr zcl.AttrID = 19

func (a LowTempDwellTripPoint) ID() zcl.AttrID                 { return LowTempDwellTripPointAttr }
func (a LowTempDwellTripPoint) Cluster() zcl.ClusterID         { return DeviceTemperatureConfigurationID }
func (a *LowTempDwellTripPoint) Value() *LowTempDwellTripPoint { return a }
func (a LowTempDwellTripPoint) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *LowTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempDwellTripPoint(*nt)
	return br, err
}

func (a LowTempDwellTripPoint) Readable() bool   { return false }
func (a LowTempDwellTripPoint) Writable() bool   { return false }
func (a LowTempDwellTripPoint) Reportable() bool { return false }
func (a LowTempDwellTripPoint) SceneIndex() int  { return -1 }

func (a LowTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}

// HighTempDwellTripPoint is an autogenerated attribute in the DeviceTemperatureConfiguration cluster
type HighTempDwellTripPoint zcl.Zu24

const HighTempDwellTripPointAttr zcl.AttrID = 20

func (a HighTempDwellTripPoint) ID() zcl.AttrID                  { return HighTempDwellTripPointAttr }
func (a HighTempDwellTripPoint) Cluster() zcl.ClusterID          { return DeviceTemperatureConfigurationID }
func (a *HighTempDwellTripPoint) Value() *HighTempDwellTripPoint { return a }
func (a HighTempDwellTripPoint) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *HighTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempDwellTripPoint(*nt)
	return br, err
}

func (a HighTempDwellTripPoint) Readable() bool   { return false }
func (a HighTempDwellTripPoint) Writable() bool   { return false }
func (a HighTempDwellTripPoint) Reportable() bool { return false }
func (a HighTempDwellTripPoint) SceneIndex() int  { return -1 }

func (a HighTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}
