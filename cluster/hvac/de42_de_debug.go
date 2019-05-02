// Attributes and commands for debugging purposes.
package hvac

import (
	"neotor.se/zcl"
)

// DeDebug
const DeDebugID zcl.ClusterID = 56898

var DeDebugCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		DebugEnabledAttr:     func() zcl.Attr { return new(DebugEnabled) },
		DebugDestinationAttr: func() zcl.Attr { return new(DebugDestination) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const DebugEnabledAttr zcl.AttrID = 0

type DebugEnabled zcl.Zbool

func (a DebugEnabled) ID() zcl.AttrID         { return DebugEnabledAttr }
func (a DebugEnabled) Cluster() zcl.ClusterID { return DeDebugID }
func (a *DebugEnabled) Value() *DebugEnabled  { return a }
func (a DebugEnabled) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *DebugEnabled) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = DebugEnabled(*nt)
	return br, err
}

func (a DebugEnabled) Readable() bool   { return true }
func (a DebugEnabled) Writable() bool   { return true }
func (a DebugEnabled) Reportable() bool { return false }
func (a DebugEnabled) SceneIndex() int  { return -1 }

func (a DebugEnabled) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

const DebugDestinationAttr zcl.AttrID = 1

type DebugDestination zcl.Zu16

func (a DebugDestination) ID() zcl.AttrID            { return DebugDestinationAttr }
func (a DebugDestination) Cluster() zcl.ClusterID    { return DeDebugID }
func (a *DebugDestination) Value() *DebugDestination { return a }
func (a DebugDestination) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DebugDestination) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DebugDestination(*nt)
	return br, err
}

func (a DebugDestination) Readable() bool   { return true }
func (a DebugDestination) Writable() bool   { return true }
func (a DebugDestination) Reportable() bool { return false }
func (a DebugDestination) SceneIndex() int  { return -1 }

func (a DebugDestination) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}
