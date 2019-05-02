package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// FlowMeasurement
const FlowMeasurementID zcl.ClusterID = 1028

var FlowMeasurementCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredFlowAttr:    func() zcl.Attr { return new(MeasuredFlow) },
		MinMeasuredFlowAttr: func() zcl.Attr { return new(MinMeasuredFlow) },
		MaxMeasuredFlowAttr: func() zcl.Attr { return new(MaxMeasuredFlow) },
		FlowToleranceAttr:   func() zcl.Attr { return new(FlowTolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const MeasuredFlowAttr zcl.AttrID = 0

type MeasuredFlow zcl.Zu16

func (a MeasuredFlow) ID() zcl.AttrID         { return MeasuredFlowAttr }
func (a MeasuredFlow) Cluster() zcl.ClusterID { return FlowMeasurementID }
func (a *MeasuredFlow) Value() *MeasuredFlow  { return a }
func (a MeasuredFlow) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredFlow(*nt)
	return br, err
}

func (a MeasuredFlow) Readable() bool   { return true }
func (a MeasuredFlow) Writable() bool   { return false }
func (a MeasuredFlow) Reportable() bool { return true }
func (a MeasuredFlow) SceneIndex() int  { return -1 }

func (a MeasuredFlow) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MinMeasuredFlowAttr zcl.AttrID = 1

type MinMeasuredFlow zcl.Zu16

func (a MinMeasuredFlow) ID() zcl.AttrID           { return MinMeasuredFlowAttr }
func (a MinMeasuredFlow) Cluster() zcl.ClusterID   { return FlowMeasurementID }
func (a *MinMeasuredFlow) Value() *MinMeasuredFlow { return a }
func (a MinMeasuredFlow) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MinMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinMeasuredFlow(*nt)
	return br, err
}

func (a MinMeasuredFlow) Readable() bool   { return true }
func (a MinMeasuredFlow) Writable() bool   { return false }
func (a MinMeasuredFlow) Reportable() bool { return false }
func (a MinMeasuredFlow) SceneIndex() int  { return -1 }

func (a MinMeasuredFlow) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MaxMeasuredFlowAttr zcl.AttrID = 2

type MaxMeasuredFlow zcl.Zu16

func (a MaxMeasuredFlow) ID() zcl.AttrID           { return MaxMeasuredFlowAttr }
func (a MaxMeasuredFlow) Cluster() zcl.ClusterID   { return FlowMeasurementID }
func (a *MaxMeasuredFlow) Value() *MaxMeasuredFlow { return a }
func (a MaxMeasuredFlow) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxMeasuredFlow(*nt)
	return br, err
}

func (a MaxMeasuredFlow) Readable() bool   { return true }
func (a MaxMeasuredFlow) Writable() bool   { return false }
func (a MaxMeasuredFlow) Reportable() bool { return false }
func (a MaxMeasuredFlow) SceneIndex() int  { return -1 }

func (a MaxMeasuredFlow) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const FlowToleranceAttr zcl.AttrID = 3

type FlowTolerance zcl.Zu16

func (a FlowTolerance) ID() zcl.AttrID         { return FlowToleranceAttr }
func (a FlowTolerance) Cluster() zcl.ClusterID { return FlowMeasurementID }
func (a *FlowTolerance) Value() *FlowTolerance { return a }
func (a FlowTolerance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *FlowTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = FlowTolerance(*nt)
	return br, err
}

func (a FlowTolerance) Readable() bool   { return true }
func (a FlowTolerance) Writable() bool   { return false }
func (a FlowTolerance) Reportable() bool { return true }
func (a FlowTolerance) SceneIndex() int  { return -1 }

func (a FlowTolerance) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}
