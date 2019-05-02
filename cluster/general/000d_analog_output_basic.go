// The Analog Output (Basic) cluster provides an interface for setting the value of an analog output (typically to the environment) and accessing various characteristics of that value.
package general

import (
	"neotor.se/zcl"
)

// AnalogOutputBasic
const AnalogOutputBasicID zcl.ClusterID = 13

var AnalogOutputBasicCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		AnalogOutputDescriptionAttr:       func() zcl.Attr { return new(AnalogOutputDescription) },
		AnalogOutputMaxPresentValueAttr:   func() zcl.Attr { return new(AnalogOutputMaxPresentValue) },
		AnalogOutputMinPresentValueAttr:   func() zcl.Attr { return new(AnalogOutputMinPresentValue) },
		AnalogOutputOutOfServiceAttr:      func() zcl.Attr { return new(AnalogOutputOutOfService) },
		AnalogOutputPresentValueAttr:      func() zcl.Attr { return new(AnalogOutputPresentValue) },
		AnalogOutputPriorityArrayAttr:     func() zcl.Attr { return new(AnalogOutputPriorityArray) },
		AnalogOutputReliabilityAttr:       func() zcl.Attr { return new(AnalogOutputReliability) },
		AnalogOutputRelinquishDefaultAttr: func() zcl.Attr { return new(AnalogOutputRelinquishDefault) },
		AnalogOutputResolutionAttr:        func() zcl.Attr { return new(AnalogOutputResolution) },
		AnalogOutputStatusFlagsAttr:       func() zcl.Attr { return new(AnalogOutputStatusFlags) },
		AnalogOutputEngineeringUnitsAttr:  func() zcl.Attr { return new(AnalogOutputEngineeringUnits) },
		AnalogOutputApplicationTypeAttr:   func() zcl.Attr { return new(AnalogOutputApplicationType) },
		AnalogOutputXiaomi0Xf000Attr:      func() zcl.Attr { return new(AnalogOutputXiaomi0Xf000) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// AnalogOutputDescription is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputDescription zcl.Zcstring

const AnalogOutputDescriptionAttr zcl.AttrID = 28

func (a AnalogOutputDescription) ID() zcl.AttrID                   { return AnalogOutputDescriptionAttr }
func (a AnalogOutputDescription) Cluster() zcl.ClusterID           { return AnalogOutputBasicID }
func (a *AnalogOutputDescription) Value() *AnalogOutputDescription { return a }
func (a AnalogOutputDescription) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *AnalogOutputDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputDescription(*nt)
	return br, err
}

func (a AnalogOutputDescription) Readable() bool   { return true }
func (a AnalogOutputDescription) Writable() bool   { return true }
func (a AnalogOutputDescription) Reportable() bool { return false }
func (a AnalogOutputDescription) SceneIndex() int  { return -1 }

func (a AnalogOutputDescription) String() string {
	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

// AnalogOutputMaxPresentValue is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputMaxPresentValue zcl.Zfloat

const AnalogOutputMaxPresentValueAttr zcl.AttrID = 65

func (a AnalogOutputMaxPresentValue) ID() zcl.AttrID                       { return AnalogOutputMaxPresentValueAttr }
func (a AnalogOutputMaxPresentValue) Cluster() zcl.ClusterID               { return AnalogOutputBasicID }
func (a *AnalogOutputMaxPresentValue) Value() *AnalogOutputMaxPresentValue { return a }
func (a AnalogOutputMaxPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputMaxPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputMaxPresentValue(*nt)
	return br, err
}

func (a AnalogOutputMaxPresentValue) Readable() bool   { return true }
func (a AnalogOutputMaxPresentValue) Writable() bool   { return true }
func (a AnalogOutputMaxPresentValue) Reportable() bool { return false }
func (a AnalogOutputMaxPresentValue) SceneIndex() int  { return -1 }

func (a AnalogOutputMaxPresentValue) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

// AnalogOutputMinPresentValue is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputMinPresentValue zcl.Zfloat

const AnalogOutputMinPresentValueAttr zcl.AttrID = 69

func (a AnalogOutputMinPresentValue) ID() zcl.AttrID                       { return AnalogOutputMinPresentValueAttr }
func (a AnalogOutputMinPresentValue) Cluster() zcl.ClusterID               { return AnalogOutputBasicID }
func (a *AnalogOutputMinPresentValue) Value() *AnalogOutputMinPresentValue { return a }
func (a AnalogOutputMinPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputMinPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputMinPresentValue(*nt)
	return br, err
}

func (a AnalogOutputMinPresentValue) Readable() bool   { return true }
func (a AnalogOutputMinPresentValue) Writable() bool   { return true }
func (a AnalogOutputMinPresentValue) Reportable() bool { return false }
func (a AnalogOutputMinPresentValue) SceneIndex() int  { return -1 }

func (a AnalogOutputMinPresentValue) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

// AnalogOutputOutOfService is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputOutOfService zcl.Zbool

const AnalogOutputOutOfServiceAttr zcl.AttrID = 81

func (a AnalogOutputOutOfService) ID() zcl.AttrID                    { return AnalogOutputOutOfServiceAttr }
func (a AnalogOutputOutOfService) Cluster() zcl.ClusterID            { return AnalogOutputBasicID }
func (a *AnalogOutputOutOfService) Value() *AnalogOutputOutOfService { return a }
func (a AnalogOutputOutOfService) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *AnalogOutputOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputOutOfService(*nt)
	return br, err
}

func (a AnalogOutputOutOfService) Readable() bool   { return true }
func (a AnalogOutputOutOfService) Writable() bool   { return true }
func (a AnalogOutputOutOfService) Reportable() bool { return false }
func (a AnalogOutputOutOfService) SceneIndex() int  { return -1 }

func (a AnalogOutputOutOfService) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

// AnalogOutputPresentValue is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputPresentValue zcl.Zfloat

const AnalogOutputPresentValueAttr zcl.AttrID = 85

func (a AnalogOutputPresentValue) ID() zcl.AttrID                    { return AnalogOutputPresentValueAttr }
func (a AnalogOutputPresentValue) Cluster() zcl.ClusterID            { return AnalogOutputBasicID }
func (a *AnalogOutputPresentValue) Value() *AnalogOutputPresentValue { return a }
func (a AnalogOutputPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputPresentValue(*nt)
	return br, err
}

func (a AnalogOutputPresentValue) Readable() bool   { return true }
func (a AnalogOutputPresentValue) Writable() bool   { return true }
func (a AnalogOutputPresentValue) Reportable() bool { return true }
func (a AnalogOutputPresentValue) SceneIndex() int  { return -1 }

func (a AnalogOutputPresentValue) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

// AnalogOutputPriorityArray is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputPriorityArray zcl.Zarray

const AnalogOutputPriorityArrayAttr zcl.AttrID = 87

func (a AnalogOutputPriorityArray) ID() zcl.AttrID                     { return AnalogOutputPriorityArrayAttr }
func (a AnalogOutputPriorityArray) Cluster() zcl.ClusterID             { return AnalogOutputBasicID }
func (a *AnalogOutputPriorityArray) Value() *AnalogOutputPriorityArray { return a }
func (a AnalogOutputPriorityArray) MarshalZcl() ([]byte, error) {
	return zcl.Zarray(a).MarshalZcl()
}
func (a *AnalogOutputPriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputPriorityArray(*nt)
	return br, err
}

func (a AnalogOutputPriorityArray) Readable() bool   { return true }
func (a AnalogOutputPriorityArray) Writable() bool   { return true }
func (a AnalogOutputPriorityArray) Reportable() bool { return false }
func (a AnalogOutputPriorityArray) SceneIndex() int  { return -1 }

func (a AnalogOutputPriorityArray) String() string {
	return zcl.Sprintf("%s", zcl.Zarray(a))
}

// AnalogOutputReliability is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputReliability zcl.Zenum8

const AnalogOutputReliabilityAttr zcl.AttrID = 103

func (a AnalogOutputReliability) ID() zcl.AttrID                   { return AnalogOutputReliabilityAttr }
func (a AnalogOutputReliability) Cluster() zcl.ClusterID           { return AnalogOutputBasicID }
func (a *AnalogOutputReliability) Value() *AnalogOutputReliability { return a }
func (a AnalogOutputReliability) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *AnalogOutputReliability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputReliability(*nt)
	return br, err
}

func (a AnalogOutputReliability) Readable() bool   { return true }
func (a AnalogOutputReliability) Writable() bool   { return true }
func (a AnalogOutputReliability) Reportable() bool { return false }
func (a AnalogOutputReliability) SceneIndex() int  { return -1 }

func (a AnalogOutputReliability) String() string {
	switch a {
	case 0x00:
		return "No fault detected"
	case 0x01:
		return "No Sensor"
	case 0x02:
		return "Over Range"
	case 0x03:
		return "Under Range"
	case 0x04:
		return "Open Loop"
	case 0x05:
		return "Shorted Loop"
	case 0x06:
		return "No Output"
	case 0x07:
		return "Unreliable (other)"
	case 0x08:
		return "Process Error"
	case 0x09:
		return "Multi state fault"
	case 0x0A:
		return "Configuration Error"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNoFaultDetected checks if AnalogOutputReliability equals the value for No fault detected (0x00)
func (a AnalogOutputReliability) IsNoFaultDetected() bool { return a == 0x00 }

// SetNoFaultDetected sets AnalogOutputReliability to No fault detected (0x00)
func (a *AnalogOutputReliability) SetNoFaultDetected() { *a = 0x00 }

// IsNoSensor checks if AnalogOutputReliability equals the value for No Sensor (0x01)
func (a AnalogOutputReliability) IsNoSensor() bool { return a == 0x01 }

// SetNoSensor sets AnalogOutputReliability to No Sensor (0x01)
func (a *AnalogOutputReliability) SetNoSensor() { *a = 0x01 }

// IsOverRange checks if AnalogOutputReliability equals the value for Over Range (0x02)
func (a AnalogOutputReliability) IsOverRange() bool { return a == 0x02 }

// SetOverRange sets AnalogOutputReliability to Over Range (0x02)
func (a *AnalogOutputReliability) SetOverRange() { *a = 0x02 }

// IsUnderRange checks if AnalogOutputReliability equals the value for Under Range (0x03)
func (a AnalogOutputReliability) IsUnderRange() bool { return a == 0x03 }

// SetUnderRange sets AnalogOutputReliability to Under Range (0x03)
func (a *AnalogOutputReliability) SetUnderRange() { *a = 0x03 }

// IsOpenLoop checks if AnalogOutputReliability equals the value for Open Loop (0x04)
func (a AnalogOutputReliability) IsOpenLoop() bool { return a == 0x04 }

// SetOpenLoop sets AnalogOutputReliability to Open Loop (0x04)
func (a *AnalogOutputReliability) SetOpenLoop() { *a = 0x04 }

// IsShortedLoop checks if AnalogOutputReliability equals the value for Shorted Loop (0x05)
func (a AnalogOutputReliability) IsShortedLoop() bool { return a == 0x05 }

// SetShortedLoop sets AnalogOutputReliability to Shorted Loop (0x05)
func (a *AnalogOutputReliability) SetShortedLoop() { *a = 0x05 }

// IsNoOutput checks if AnalogOutputReliability equals the value for No Output (0x06)
func (a AnalogOutputReliability) IsNoOutput() bool { return a == 0x06 }

// SetNoOutput sets AnalogOutputReliability to No Output (0x06)
func (a *AnalogOutputReliability) SetNoOutput() { *a = 0x06 }

// IsUnreliableOther checks if AnalogOutputReliability equals the value for Unreliable (other) (0x07)
func (a AnalogOutputReliability) IsUnreliableOther() bool { return a == 0x07 }

// SetUnreliableOther sets AnalogOutputReliability to Unreliable (other) (0x07)
func (a *AnalogOutputReliability) SetUnreliableOther() { *a = 0x07 }

// IsProcessError checks if AnalogOutputReliability equals the value for Process Error (0x08)
func (a AnalogOutputReliability) IsProcessError() bool { return a == 0x08 }

// SetProcessError sets AnalogOutputReliability to Process Error (0x08)
func (a *AnalogOutputReliability) SetProcessError() { *a = 0x08 }

// IsMultiStateFault checks if AnalogOutputReliability equals the value for Multi state fault (0x09)
func (a AnalogOutputReliability) IsMultiStateFault() bool { return a == 0x09 }

// SetMultiStateFault sets AnalogOutputReliability to Multi state fault (0x09)
func (a *AnalogOutputReliability) SetMultiStateFault() { *a = 0x09 }

// IsConfigurationError checks if AnalogOutputReliability equals the value for Configuration Error (0x0A)
func (a AnalogOutputReliability) IsConfigurationError() bool { return a == 0x0A }

// SetConfigurationError sets AnalogOutputReliability to Configuration Error (0x0A)
func (a *AnalogOutputReliability) SetConfigurationError() { *a = 0x0A }

// AnalogOutputRelinquishDefault is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputRelinquishDefault zcl.Zfloat

const AnalogOutputRelinquishDefaultAttr zcl.AttrID = 104

func (a AnalogOutputRelinquishDefault) ID() zcl.AttrID                         { return AnalogOutputRelinquishDefaultAttr }
func (a AnalogOutputRelinquishDefault) Cluster() zcl.ClusterID                 { return AnalogOutputBasicID }
func (a *AnalogOutputRelinquishDefault) Value() *AnalogOutputRelinquishDefault { return a }
func (a AnalogOutputRelinquishDefault) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputRelinquishDefault(*nt)
	return br, err
}

func (a AnalogOutputRelinquishDefault) Readable() bool   { return true }
func (a AnalogOutputRelinquishDefault) Writable() bool   { return true }
func (a AnalogOutputRelinquishDefault) Reportable() bool { return false }
func (a AnalogOutputRelinquishDefault) SceneIndex() int  { return -1 }

func (a AnalogOutputRelinquishDefault) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

// AnalogOutputResolution is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputResolution zcl.Zfloat

const AnalogOutputResolutionAttr zcl.AttrID = 106

func (a AnalogOutputResolution) ID() zcl.AttrID                  { return AnalogOutputResolutionAttr }
func (a AnalogOutputResolution) Cluster() zcl.ClusterID          { return AnalogOutputBasicID }
func (a *AnalogOutputResolution) Value() *AnalogOutputResolution { return a }
func (a AnalogOutputResolution) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputResolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputResolution(*nt)
	return br, err
}

func (a AnalogOutputResolution) Readable() bool   { return true }
func (a AnalogOutputResolution) Writable() bool   { return true }
func (a AnalogOutputResolution) Reportable() bool { return false }
func (a AnalogOutputResolution) SceneIndex() int  { return -1 }

func (a AnalogOutputResolution) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

// AnalogOutputStatusFlags is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputStatusFlags zcl.Zbmp8

const AnalogOutputStatusFlagsAttr zcl.AttrID = 111

func (a AnalogOutputStatusFlags) ID() zcl.AttrID                   { return AnalogOutputStatusFlagsAttr }
func (a AnalogOutputStatusFlags) Cluster() zcl.ClusterID           { return AnalogOutputBasicID }
func (a *AnalogOutputStatusFlags) Value() *AnalogOutputStatusFlags { return a }
func (a AnalogOutputStatusFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *AnalogOutputStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputStatusFlags(*nt)
	return br, err
}

func (a AnalogOutputStatusFlags) Readable() bool   { return true }
func (a AnalogOutputStatusFlags) Writable() bool   { return false }
func (a AnalogOutputStatusFlags) Reportable() bool { return true }
func (a AnalogOutputStatusFlags) SceneIndex() int  { return -1 }

func (a AnalogOutputStatusFlags) String() string {
	var bstr []string
	if a.IsInAlarm() {
		bstr = append(bstr, "In Alarm")
	}
	if a.IsFault() {
		bstr = append(bstr, "Fault")
	}
	if a.IsOveridden() {
		bstr = append(bstr, "Overidden")
	}
	if a.IsOutOfService() {
		bstr = append(bstr, "Out of Service")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a AnalogOutputStatusFlags) IsInAlarm() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *AnalogOutputStatusFlags) SetInAlarm(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a AnalogOutputStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AnalogOutputStatusFlags) SetFault(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AnalogOutputStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AnalogOutputStatusFlags) SetOveridden(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AnalogOutputStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *AnalogOutputStatusFlags) SetOutOfService(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 3, b))
}

// AnalogOutputEngineeringUnits is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputEngineeringUnits zcl.EngineeringUnit

const AnalogOutputEngineeringUnitsAttr zcl.AttrID = 117

func (a AnalogOutputEngineeringUnits) ID() zcl.AttrID                        { return AnalogOutputEngineeringUnitsAttr }
func (a AnalogOutputEngineeringUnits) Cluster() zcl.ClusterID                { return AnalogOutputBasicID }
func (a *AnalogOutputEngineeringUnits) Value() *AnalogOutputEngineeringUnits { return a }
func (a AnalogOutputEngineeringUnits) MarshalZcl() ([]byte, error) {
	return zcl.EngineeringUnit(a).MarshalZcl()
}
func (a *AnalogOutputEngineeringUnits) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.EngineeringUnit)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputEngineeringUnits(*nt)
	return br, err
}

func (a AnalogOutputEngineeringUnits) Readable() bool   { return true }
func (a AnalogOutputEngineeringUnits) Writable() bool   { return true }
func (a AnalogOutputEngineeringUnits) Reportable() bool { return false }
func (a AnalogOutputEngineeringUnits) SceneIndex() int  { return -1 }

func (a AnalogOutputEngineeringUnits) String() string {
	return zcl.Sprintf("%s", zcl.EngineeringUnit(a))
}

// AnalogOutputApplicationType is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputApplicationType zcl.Zu32

const AnalogOutputApplicationTypeAttr zcl.AttrID = 256

func (a AnalogOutputApplicationType) ID() zcl.AttrID                       { return AnalogOutputApplicationTypeAttr }
func (a AnalogOutputApplicationType) Cluster() zcl.ClusterID               { return AnalogOutputBasicID }
func (a *AnalogOutputApplicationType) Value() *AnalogOutputApplicationType { return a }
func (a AnalogOutputApplicationType) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *AnalogOutputApplicationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputApplicationType(*nt)
	return br, err
}

func (a AnalogOutputApplicationType) Readable() bool   { return true }
func (a AnalogOutputApplicationType) Writable() bool   { return false }
func (a AnalogOutputApplicationType) Reportable() bool { return false }
func (a AnalogOutputApplicationType) SceneIndex() int  { return -1 }

func (a AnalogOutputApplicationType) String() string {
	return zcl.Sprintf("%s", zcl.Zu32(a))
}

// AnalogOutputXiaomi0Xf000 is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputXiaomi0Xf000 zcl.Zu32

const AnalogOutputXiaomi0Xf000Attr zcl.AttrID = 61440

func (a AnalogOutputXiaomi0Xf000) ID() zcl.AttrID                    { return AnalogOutputXiaomi0Xf000Attr }
func (a AnalogOutputXiaomi0Xf000) Cluster() zcl.ClusterID            { return AnalogOutputBasicID }
func (a *AnalogOutputXiaomi0Xf000) Value() *AnalogOutputXiaomi0Xf000 { return a }
func (a AnalogOutputXiaomi0Xf000) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *AnalogOutputXiaomi0Xf000) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputXiaomi0Xf000(*nt)
	return br, err
}

func (a AnalogOutputXiaomi0Xf000) Readable() bool   { return true }
func (a AnalogOutputXiaomi0Xf000) Writable() bool   { return false }
func (a AnalogOutputXiaomi0Xf000) Reportable() bool { return false }
func (a AnalogOutputXiaomi0Xf000) SceneIndex() int  { return -1 }

func (a AnalogOutputXiaomi0Xf000) String() string {
	return zcl.Sprintf("%s", zcl.Zu32(a))
}
