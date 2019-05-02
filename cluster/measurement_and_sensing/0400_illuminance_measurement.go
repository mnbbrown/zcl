// The server cluster provides an interface to illuminance measurement functionality, including configuration and provision of notifications of illuminance measurements.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// IlluminanceMeasurement
const IlluminanceMeasurementID zcl.ClusterID = 1024

var IlluminanceMeasurementCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredIlluminanceAttr:    func() zcl.Attr { return new(MeasuredIlluminance) },
		MinMeasuredIlluminanceAttr: func() zcl.Attr { return new(MinMeasuredIlluminance) },
		MaxMeasuredIlluminanceAttr: func() zcl.Attr { return new(MaxMeasuredIlluminance) },
		IlluminanceToleranceAttr:   func() zcl.Attr { return new(IlluminanceTolerance) },
		LightSensorTypeAttr:        func() zcl.Attr { return new(LightSensorType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{
		HysteresisAttr:   func() zcl.Attr { return new(Hysteresis) },
		MaxUpSpeedAttr:   func() zcl.Attr { return new(MaxUpSpeed) },
		MaxDownSpeedAttr: func() zcl.Attr { return new(MaxDownSpeed) },
		TargetValueAttr:  func() zcl.Attr { return new(TargetValue) },
		StartupTypeAttr:  func() zcl.Attr { return new(StartupType) },
	},
	SceneAttr: []zcl.AttrID{},
}

const MeasuredIlluminanceAttr zcl.AttrID = 0

type MeasuredIlluminance zcl.Zu16

func (a MeasuredIlluminance) ID() zcl.AttrID               { return MeasuredIlluminanceAttr }
func (a MeasuredIlluminance) Cluster() zcl.ClusterID       { return IlluminanceMeasurementID }
func (a *MeasuredIlluminance) Value() *MeasuredIlluminance { return a }
func (a MeasuredIlluminance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredIlluminance(*nt)
	return br, err
}

func (a MeasuredIlluminance) Readable() bool   { return true }
func (a MeasuredIlluminance) Writable() bool   { return false }
func (a MeasuredIlluminance) Reportable() bool { return false }
func (a MeasuredIlluminance) SceneIndex() int  { return -1 }

func (a MeasuredIlluminance) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MinMeasuredIlluminanceAttr zcl.AttrID = 1

type MinMeasuredIlluminance zcl.Zu16

func (a MinMeasuredIlluminance) ID() zcl.AttrID                  { return MinMeasuredIlluminanceAttr }
func (a MinMeasuredIlluminance) Cluster() zcl.ClusterID          { return IlluminanceMeasurementID }
func (a *MinMeasuredIlluminance) Value() *MinMeasuredIlluminance { return a }
func (a MinMeasuredIlluminance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MinMeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinMeasuredIlluminance(*nt)
	return br, err
}

func (a MinMeasuredIlluminance) Readable() bool   { return true }
func (a MinMeasuredIlluminance) Writable() bool   { return false }
func (a MinMeasuredIlluminance) Reportable() bool { return true }
func (a MinMeasuredIlluminance) SceneIndex() int  { return -1 }

func (a MinMeasuredIlluminance) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MaxMeasuredIlluminanceAttr zcl.AttrID = 2

type MaxMeasuredIlluminance zcl.Zu16

func (a MaxMeasuredIlluminance) ID() zcl.AttrID                  { return MaxMeasuredIlluminanceAttr }
func (a MaxMeasuredIlluminance) Cluster() zcl.ClusterID          { return IlluminanceMeasurementID }
func (a *MaxMeasuredIlluminance) Value() *MaxMeasuredIlluminance { return a }
func (a MaxMeasuredIlluminance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxMeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxMeasuredIlluminance(*nt)
	return br, err
}

func (a MaxMeasuredIlluminance) Readable() bool   { return true }
func (a MaxMeasuredIlluminance) Writable() bool   { return false }
func (a MaxMeasuredIlluminance) Reportable() bool { return false }
func (a MaxMeasuredIlluminance) SceneIndex() int  { return -1 }

func (a MaxMeasuredIlluminance) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const IlluminanceToleranceAttr zcl.AttrID = 3

type IlluminanceTolerance zcl.Zu16

func (a IlluminanceTolerance) ID() zcl.AttrID                { return IlluminanceToleranceAttr }
func (a IlluminanceTolerance) Cluster() zcl.ClusterID        { return IlluminanceMeasurementID }
func (a *IlluminanceTolerance) Value() *IlluminanceTolerance { return a }
func (a IlluminanceTolerance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *IlluminanceTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IlluminanceTolerance(*nt)
	return br, err
}

func (a IlluminanceTolerance) Readable() bool   { return true }
func (a IlluminanceTolerance) Writable() bool   { return false }
func (a IlluminanceTolerance) Reportable() bool { return false }
func (a IlluminanceTolerance) SceneIndex() int  { return -1 }

func (a IlluminanceTolerance) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const LightSensorTypeAttr zcl.AttrID = 4

type LightSensorType zcl.Zenum8

func (a LightSensorType) ID() zcl.AttrID           { return LightSensorTypeAttr }
func (a LightSensorType) Cluster() zcl.ClusterID   { return IlluminanceMeasurementID }
func (a *LightSensorType) Value() *LightSensorType { return a }
func (a LightSensorType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *LightSensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LightSensorType(*nt)
	return br, err
}

func (a LightSensorType) Readable() bool   { return true }
func (a LightSensorType) Writable() bool   { return false }
func (a LightSensorType) Reportable() bool { return false }
func (a LightSensorType) SceneIndex() int  { return -1 }

func (a LightSensorType) String() string {
	switch a {
	case 0x00:
		return "Photodiode"
	case 0x01:
		return "CMOS"
	case 0xFF:
		return "Unknown"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsPhotodiode checks if LightSensorType equals the value for Photodiode (0x00)
func (a LightSensorType) IsPhotodiode() bool { return a == 0x00 }

// SetPhotodiode sets LightSensorType to Photodiode (0x00)
func (a *LightSensorType) SetPhotodiode() { *a = 0x00 }

// IsCmos checks if LightSensorType equals the value for CMOS (0x01)
func (a LightSensorType) IsCmos() bool { return a == 0x01 }

// SetCmos sets LightSensorType to CMOS (0x01)
func (a *LightSensorType) SetCmos() { *a = 0x01 }

// IsUnknown checks if LightSensorType equals the value for Unknown (0xFF)
func (a LightSensorType) IsUnknown() bool { return a == 0xFF }

// SetUnknown sets LightSensorType to Unknown (0xFF)
func (a *LightSensorType) SetUnknown() { *a = 0xFF }

const HysteresisAttr zcl.AttrID = 61440

type Hysteresis zcl.Zu16

func (a Hysteresis) ID() zcl.AttrID         { return HysteresisAttr }
func (a Hysteresis) Cluster() zcl.ClusterID { return IlluminanceMeasurementID }
func (a *Hysteresis) Value() *Hysteresis    { return a }
func (a Hysteresis) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Hysteresis) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Hysteresis(*nt)
	return br, err
}

func (a Hysteresis) Readable() bool   { return true }
func (a Hysteresis) Writable() bool   { return true }
func (a Hysteresis) Reportable() bool { return false }
func (a Hysteresis) SceneIndex() int  { return -1 }

func (a Hysteresis) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MaxUpSpeedAttr zcl.AttrID = 61441

type MaxUpSpeed zcl.Zu16

func (a MaxUpSpeed) ID() zcl.AttrID         { return MaxUpSpeedAttr }
func (a MaxUpSpeed) Cluster() zcl.ClusterID { return IlluminanceMeasurementID }
func (a *MaxUpSpeed) Value() *MaxUpSpeed    { return a }
func (a MaxUpSpeed) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxUpSpeed) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxUpSpeed(*nt)
	return br, err
}

func (a MaxUpSpeed) Readable() bool   { return true }
func (a MaxUpSpeed) Writable() bool   { return true }
func (a MaxUpSpeed) Reportable() bool { return false }
func (a MaxUpSpeed) SceneIndex() int  { return -1 }

func (a MaxUpSpeed) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MaxDownSpeedAttr zcl.AttrID = 61442

type MaxDownSpeed zcl.Zu16

func (a MaxDownSpeed) ID() zcl.AttrID         { return MaxDownSpeedAttr }
func (a MaxDownSpeed) Cluster() zcl.ClusterID { return IlluminanceMeasurementID }
func (a *MaxDownSpeed) Value() *MaxDownSpeed  { return a }
func (a MaxDownSpeed) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxDownSpeed) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxDownSpeed(*nt)
	return br, err
}

func (a MaxDownSpeed) Readable() bool   { return true }
func (a MaxDownSpeed) Writable() bool   { return true }
func (a MaxDownSpeed) Reportable() bool { return false }
func (a MaxDownSpeed) SceneIndex() int  { return -1 }

func (a MaxDownSpeed) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const TargetValueAttr zcl.AttrID = 61443

type TargetValue zcl.Zu16

func (a TargetValue) ID() zcl.AttrID         { return TargetValueAttr }
func (a TargetValue) Cluster() zcl.ClusterID { return IlluminanceMeasurementID }
func (a *TargetValue) Value() *TargetValue   { return a }
func (a TargetValue) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *TargetValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TargetValue(*nt)
	return br, err
}

func (a TargetValue) Readable() bool   { return true }
func (a TargetValue) Writable() bool   { return true }
func (a TargetValue) Reportable() bool { return false }
func (a TargetValue) SceneIndex() int  { return -1 }

func (a TargetValue) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const StartupTypeAttr zcl.AttrID = 61444

type StartupType zcl.Zenum8

func (a StartupType) ID() zcl.AttrID         { return StartupTypeAttr }
func (a StartupType) Cluster() zcl.ClusterID { return IlluminanceMeasurementID }
func (a *StartupType) Value() *StartupType   { return a }
func (a StartupType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *StartupType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartupType(*nt)
	return br, err
}

func (a StartupType) Readable() bool   { return true }
func (a StartupType) Writable() bool   { return true }
func (a StartupType) Reportable() bool { return false }
func (a StartupType) SceneIndex() int  { return -1 }

func (a StartupType) String() string {
	switch a {
	case 0x00:
		return "Default Level"
	case 0x01:
		return "Zero Level"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsDefaultLevel checks if StartupType equals the value for Default Level (0x00)
func (a StartupType) IsDefaultLevel() bool { return a == 0x00 }

// SetDefaultLevel sets StartupType to Default Level (0x00)
func (a *StartupType) SetDefaultLevel() { *a = 0x00 }

// IsZeroLevel checks if StartupType equals the value for Zero Level (0x01)
func (a StartupType) IsZeroLevel() bool { return a == 0x01 }

// SetZeroLevel sets StartupType to Zero Level (0x01)
func (a *StartupType) SetZeroLevel() { *a = 0x01 }
