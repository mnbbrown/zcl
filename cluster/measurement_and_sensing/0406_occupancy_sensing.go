package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// OccupancySensing
const OccupancySensingID zcl.ClusterID = 1030

var OccupancySensingCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		OccupancyAttr:                               func() zcl.Attr { return new(Occupancy) },
		OccupancySensorTypeAttr:                     func() zcl.Attr { return new(OccupancySensorType) },
		PirOccupiedToUnoccupiedDelayAttr:            func() zcl.Attr { return new(PirOccupiedToUnoccupiedDelay) },
		PirUnoccupiedToOccupiedDelayAttr:            func() zcl.Attr { return new(PirUnoccupiedToOccupiedDelay) },
		PirUnoccupiedToOccupiedThresholdAttr:        func() zcl.Attr { return new(PirUnoccupiedToOccupiedThreshold) },
		UltrasonicOccupiedToUnoccupiedDelayAttr:     func() zcl.Attr { return new(UltrasonicOccupiedToUnoccupiedDelay) },
		UltrasonicUnoccupiedToOccupiedDelayAttr:     func() zcl.Attr { return new(UltrasonicUnoccupiedToOccupiedDelay) },
		UltrasonicUnoccupiedToOccupiedThresholdAttr: func() zcl.Attr { return new(UltrasonicUnoccupiedToOccupiedThreshold) },
		SensitivityAttr:                             func() zcl.Attr { return new(Sensitivity) },
		SensitivityMaxAttr:                          func() zcl.Attr { return new(SensitivityMax) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const OccupancyAttr zcl.AttrID = 0

type Occupancy zcl.Zbmp8

func (a Occupancy) ID() zcl.AttrID         { return OccupancyAttr }
func (a Occupancy) Cluster() zcl.ClusterID { return OccupancySensingID }
func (a *Occupancy) Value() *Occupancy     { return a }
func (a Occupancy) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Occupancy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Occupancy(*nt)
	return br, err
}

func (a Occupancy) Readable() bool   { return true }
func (a Occupancy) Writable() bool   { return false }
func (a Occupancy) Reportable() bool { return true }
func (a Occupancy) SceneIndex() int  { return -1 }

func (a Occupancy) String() string {
	return zcl.Sprintf("%s", zcl.Zbmp8(a))
}

const OccupancySensorTypeAttr zcl.AttrID = 1

type OccupancySensorType zcl.Zenum8

func (a OccupancySensorType) ID() zcl.AttrID               { return OccupancySensorTypeAttr }
func (a OccupancySensorType) Cluster() zcl.ClusterID       { return OccupancySensingID }
func (a *OccupancySensorType) Value() *OccupancySensorType { return a }
func (a OccupancySensorType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *OccupancySensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = OccupancySensorType(*nt)
	return br, err
}

func (a OccupancySensorType) Readable() bool   { return true }
func (a OccupancySensorType) Writable() bool   { return false }
func (a OccupancySensorType) Reportable() bool { return false }
func (a OccupancySensorType) SceneIndex() int  { return -1 }

func (a OccupancySensorType) String() string {
	switch a {
	case 0x00:
		return "PIR"
	case 0x01:
		return "Ultrasonic"
	case 0x02:
		return "PIR and ultrasonic"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsPir checks if OccupancySensorType equals the value for PIR (0x00)
func (a OccupancySensorType) IsPir() bool { return a == 0x00 }

// SetPir sets OccupancySensorType to PIR (0x00)
func (a *OccupancySensorType) SetPir() { *a = 0x00 }

// IsUltrasonic checks if OccupancySensorType equals the value for Ultrasonic (0x01)
func (a OccupancySensorType) IsUltrasonic() bool { return a == 0x01 }

// SetUltrasonic sets OccupancySensorType to Ultrasonic (0x01)
func (a *OccupancySensorType) SetUltrasonic() { *a = 0x01 }

// IsPirAndUltrasonic checks if OccupancySensorType equals the value for PIR and ultrasonic (0x02)
func (a OccupancySensorType) IsPirAndUltrasonic() bool { return a == 0x02 }

// SetPirAndUltrasonic sets OccupancySensorType to PIR and ultrasonic (0x02)
func (a *OccupancySensorType) SetPirAndUltrasonic() { *a = 0x02 }

const PirOccupiedToUnoccupiedDelayAttr zcl.AttrID = 16

type PirOccupiedToUnoccupiedDelay zcl.Zu16

func (a PirOccupiedToUnoccupiedDelay) ID() zcl.AttrID                        { return PirOccupiedToUnoccupiedDelayAttr }
func (a PirOccupiedToUnoccupiedDelay) Cluster() zcl.ClusterID                { return OccupancySensingID }
func (a *PirOccupiedToUnoccupiedDelay) Value() *PirOccupiedToUnoccupiedDelay { return a }
func (a PirOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PirOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PirOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (a PirOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (a PirOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (a PirOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (a PirOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

func (a PirOccupiedToUnoccupiedDelay) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PirUnoccupiedToOccupiedDelayAttr zcl.AttrID = 17

type PirUnoccupiedToOccupiedDelay zcl.Zu16

func (a PirUnoccupiedToOccupiedDelay) ID() zcl.AttrID                        { return PirUnoccupiedToOccupiedDelayAttr }
func (a PirUnoccupiedToOccupiedDelay) Cluster() zcl.ClusterID                { return OccupancySensingID }
func (a *PirUnoccupiedToOccupiedDelay) Value() *PirUnoccupiedToOccupiedDelay { return a }
func (a PirUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PirUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PirUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (a PirUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (a PirUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (a PirUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (a PirUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

func (a PirUnoccupiedToOccupiedDelay) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PirUnoccupiedToOccupiedThresholdAttr zcl.AttrID = 18

type PirUnoccupiedToOccupiedThreshold zcl.Zu8

func (a PirUnoccupiedToOccupiedThreshold) ID() zcl.AttrID                            { return PirUnoccupiedToOccupiedThresholdAttr }
func (a PirUnoccupiedToOccupiedThreshold) Cluster() zcl.ClusterID                    { return OccupancySensingID }
func (a *PirUnoccupiedToOccupiedThreshold) Value() *PirUnoccupiedToOccupiedThreshold { return a }
func (a PirUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PirUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PirUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (a PirUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (a PirUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (a PirUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (a PirUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

func (a PirUnoccupiedToOccupiedThreshold) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const UltrasonicOccupiedToUnoccupiedDelayAttr zcl.AttrID = 32

type UltrasonicOccupiedToUnoccupiedDelay zcl.Zu16

func (a UltrasonicOccupiedToUnoccupiedDelay) ID() zcl.AttrID {
	return UltrasonicOccupiedToUnoccupiedDelayAttr
}
func (a UltrasonicOccupiedToUnoccupiedDelay) Cluster() zcl.ClusterID                       { return OccupancySensingID }
func (a *UltrasonicOccupiedToUnoccupiedDelay) Value() *UltrasonicOccupiedToUnoccupiedDelay { return a }
func (a UltrasonicOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *UltrasonicOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = UltrasonicOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (a UltrasonicOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (a UltrasonicOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (a UltrasonicOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (a UltrasonicOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

func (a UltrasonicOccupiedToUnoccupiedDelay) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const UltrasonicUnoccupiedToOccupiedDelayAttr zcl.AttrID = 33

type UltrasonicUnoccupiedToOccupiedDelay zcl.Zu16

func (a UltrasonicUnoccupiedToOccupiedDelay) ID() zcl.AttrID {
	return UltrasonicUnoccupiedToOccupiedDelayAttr
}
func (a UltrasonicUnoccupiedToOccupiedDelay) Cluster() zcl.ClusterID                       { return OccupancySensingID }
func (a *UltrasonicUnoccupiedToOccupiedDelay) Value() *UltrasonicUnoccupiedToOccupiedDelay { return a }
func (a UltrasonicUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *UltrasonicUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = UltrasonicUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (a UltrasonicUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (a UltrasonicUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

func (a UltrasonicUnoccupiedToOccupiedDelay) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const UltrasonicUnoccupiedToOccupiedThresholdAttr zcl.AttrID = 34

type UltrasonicUnoccupiedToOccupiedThreshold zcl.Zu8

func (a UltrasonicUnoccupiedToOccupiedThreshold) ID() zcl.AttrID {
	return UltrasonicUnoccupiedToOccupiedThresholdAttr
}
func (a UltrasonicUnoccupiedToOccupiedThreshold) Cluster() zcl.ClusterID { return OccupancySensingID }
func (a *UltrasonicUnoccupiedToOccupiedThreshold) Value() *UltrasonicUnoccupiedToOccupiedThreshold {
	return a
}
func (a UltrasonicUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *UltrasonicUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = UltrasonicUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (a UltrasonicUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (a UltrasonicUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (a UltrasonicUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

func (a UltrasonicUnoccupiedToOccupiedThreshold) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const SensitivityAttr zcl.AttrID = 48

type Sensitivity zcl.Zu8

func (a Sensitivity) ID() zcl.AttrID         { return SensitivityAttr }
func (a Sensitivity) Cluster() zcl.ClusterID { return OccupancySensingID }
func (a *Sensitivity) Value() *Sensitivity   { return a }
func (a Sensitivity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Sensitivity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Sensitivity(*nt)
	return br, err
}

func (a Sensitivity) Readable() bool   { return true }
func (a Sensitivity) Writable() bool   { return true }
func (a Sensitivity) Reportable() bool { return false }
func (a Sensitivity) SceneIndex() int  { return -1 }

func (a Sensitivity) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const SensitivityMaxAttr zcl.AttrID = 49

type SensitivityMax zcl.Zu8

func (a SensitivityMax) ID() zcl.AttrID          { return SensitivityMaxAttr }
func (a SensitivityMax) Cluster() zcl.ClusterID  { return OccupancySensingID }
func (a *SensitivityMax) Value() *SensitivityMax { return a }
func (a SensitivityMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *SensitivityMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SensitivityMax(*nt)
	return br, err
}

func (a SensitivityMax) Readable() bool   { return true }
func (a SensitivityMax) Writable() bool   { return true }
func (a SensitivityMax) Reportable() bool { return false }
func (a SensitivityMax) SceneIndex() int  { return -1 }

func (a SensitivityMax) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}
