package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// IlluminanceLevelSensing

func NewIlluminanceLevelSensingServer(profile zcl.ProfileID) *IlluminanceLevelSensingServer {
	return &IlluminanceLevelSensingServer{p: profile}
}
func NewIlluminanceLevelSensingClient(profile zcl.ProfileID) *IlluminanceLevelSensingClient {
	return &IlluminanceLevelSensingClient{p: profile}
}

const IlluminanceLevelSensingCluster zcl.ClusterID = 1025

type IlluminanceLevelSensingServer struct {
	p zcl.ProfileID

	LevelStatus            *LevelStatus
	IlluminanceSensorType  *IlluminanceSensorType
	IlluminanceTargetLevel *IlluminanceTargetLevel
}

type IlluminanceLevelSensingClient struct {
	p zcl.ProfileID
}

/*
var IlluminanceLevelSensingServer = map[zcl.CommandID]func() zcl.Command{
}

var IlluminanceLevelSensingClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type LevelStatus zcl.Zenum8

func (a LevelStatus) ID() zcl.AttrID         { return 0 }
func (a LevelStatus) Cluster() zcl.ClusterID { return IlluminanceLevelSensingCluster }
func (a *LevelStatus) Value() *LevelStatus   { return a }
func (a LevelStatus) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *LevelStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LevelStatus(*nt)
	return br, err
}

func (a LevelStatus) Readable() bool   { return true }
func (a LevelStatus) Writable() bool   { return false }
func (a LevelStatus) Reportable() bool { return true }
func (a LevelStatus) SceneIndex() int  { return -1 }

func (a LevelStatus) String() string {
	switch a {
	case 0x00:
		return "Illuminance on target"
	case 0x01:
		return "Illuminance below target"
	case 0x02:
		return "Illuminance above target"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsIlluminanceOnTarget checks if LevelStatus equals the value for Illuminance on target (0x00)
func (a LevelStatus) IsIlluminanceOnTarget() bool { return a == 0x00 }

// SetIlluminanceOnTarget sets LevelStatus to Illuminance on target (0x00)
func (a *LevelStatus) SetIlluminanceOnTarget() { *a = 0x00 }

// IsIlluminanceBelowTarget checks if LevelStatus equals the value for Illuminance below target (0x01)
func (a LevelStatus) IsIlluminanceBelowTarget() bool { return a == 0x01 }

// SetIlluminanceBelowTarget sets LevelStatus to Illuminance below target (0x01)
func (a *LevelStatus) SetIlluminanceBelowTarget() { *a = 0x01 }

// IsIlluminanceAboveTarget checks if LevelStatus equals the value for Illuminance above target (0x02)
func (a LevelStatus) IsIlluminanceAboveTarget() bool { return a == 0x02 }

// SetIlluminanceAboveTarget sets LevelStatus to Illuminance above target (0x02)
func (a *LevelStatus) SetIlluminanceAboveTarget() { *a = 0x02 }

type IlluminanceSensorType zcl.Zenum8

func (a IlluminanceSensorType) ID() zcl.AttrID                 { return 1 }
func (a IlluminanceSensorType) Cluster() zcl.ClusterID         { return IlluminanceLevelSensingCluster }
func (a *IlluminanceSensorType) Value() *IlluminanceSensorType { return a }
func (a IlluminanceSensorType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *IlluminanceSensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IlluminanceSensorType(*nt)
	return br, err
}

func (a IlluminanceSensorType) Readable() bool   { return true }
func (a IlluminanceSensorType) Writable() bool   { return false }
func (a IlluminanceSensorType) Reportable() bool { return false }
func (a IlluminanceSensorType) SceneIndex() int  { return -1 }

func (a IlluminanceSensorType) String() string {
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

// IsPhotodiode checks if IlluminanceSensorType equals the value for Photodiode (0x00)
func (a IlluminanceSensorType) IsPhotodiode() bool { return a == 0x00 }

// SetPhotodiode sets IlluminanceSensorType to Photodiode (0x00)
func (a *IlluminanceSensorType) SetPhotodiode() { *a = 0x00 }

// IsCmos checks if IlluminanceSensorType equals the value for CMOS (0x01)
func (a IlluminanceSensorType) IsCmos() bool { return a == 0x01 }

// SetCmos sets IlluminanceSensorType to CMOS (0x01)
func (a *IlluminanceSensorType) SetCmos() { *a = 0x01 }

// IsUnknown checks if IlluminanceSensorType equals the value for Unknown (0xFF)
func (a IlluminanceSensorType) IsUnknown() bool { return a == 0xFF }

// SetUnknown sets IlluminanceSensorType to Unknown (0xFF)
func (a *IlluminanceSensorType) SetUnknown() { *a = 0xFF }

type IlluminanceTargetLevel zcl.Zu16

func (a IlluminanceTargetLevel) ID() zcl.AttrID                  { return 16 }
func (a IlluminanceTargetLevel) Cluster() zcl.ClusterID          { return IlluminanceLevelSensingCluster }
func (a *IlluminanceTargetLevel) Value() *IlluminanceTargetLevel { return a }
func (a IlluminanceTargetLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *IlluminanceTargetLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IlluminanceTargetLevel(*nt)
	return br, err
}

func (a IlluminanceTargetLevel) Readable() bool   { return true }
func (a IlluminanceTargetLevel) Writable() bool   { return true }
func (a IlluminanceTargetLevel) Reportable() bool { return false }
func (a IlluminanceTargetLevel) SceneIndex() int  { return -1 }

func (a IlluminanceTargetLevel) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
