// Attributes and commands for configuring On/Off switching devices
package general

import (
	"neotor.se/zcl"
)

// OnOffSwitchConfiguration
const OnOffSwitchConfigurationID zcl.ClusterID = 7

var OnOffSwitchConfigurationCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		SwitchtypeAttr:    func() zcl.Attr { return new(Switchtype) },
		SwitchactionsAttr: func() zcl.Attr { return new(Switchactions) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const SwitchtypeAttr zcl.AttrID = 0

type Switchtype zcl.Zenum8

func (a Switchtype) ID() zcl.AttrID         { return SwitchtypeAttr }
func (a Switchtype) Cluster() zcl.ClusterID { return OnOffSwitchConfigurationID }
func (a *Switchtype) Value() *Switchtype    { return a }
func (a Switchtype) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Switchtype) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Switchtype(*nt)
	return br, err
}

func (a Switchtype) Readable() bool   { return true }
func (a Switchtype) Writable() bool   { return false }
func (a Switchtype) Reportable() bool { return false }
func (a Switchtype) SceneIndex() int  { return -1 }

func (a Switchtype) String() string {
	switch a {
	case 0x00:
		return "Toggle"
	case 0x01:
		return "Momentary"
	case 0x02:
		return "Multifunction"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsToggle checks if Switchtype equals the value for Toggle (0x00)
func (a Switchtype) IsToggle() bool { return a == 0x00 }

// SetToggle sets Switchtype to Toggle (0x00)
func (a *Switchtype) SetToggle() { *a = 0x00 }

// IsMomentary checks if Switchtype equals the value for Momentary (0x01)
func (a Switchtype) IsMomentary() bool { return a == 0x01 }

// SetMomentary sets Switchtype to Momentary (0x01)
func (a *Switchtype) SetMomentary() { *a = 0x01 }

// IsMultifunction checks if Switchtype equals the value for Multifunction (0x02)
func (a Switchtype) IsMultifunction() bool { return a == 0x02 }

// SetMultifunction sets Switchtype to Multifunction (0x02)
func (a *Switchtype) SetMultifunction() { *a = 0x02 }

const SwitchactionsAttr zcl.AttrID = 16

type Switchactions zcl.Zenum8

func (a Switchactions) ID() zcl.AttrID         { return SwitchactionsAttr }
func (a Switchactions) Cluster() zcl.ClusterID { return OnOffSwitchConfigurationID }
func (a *Switchactions) Value() *Switchactions { return a }
func (a Switchactions) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Switchactions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Switchactions(*nt)
	return br, err
}

func (a Switchactions) Readable() bool   { return true }
func (a Switchactions) Writable() bool   { return true }
func (a Switchactions) Reportable() bool { return false }
func (a Switchactions) SceneIndex() int  { return -1 }

func (a Switchactions) String() string {
	switch a {
	case 0x00:
		return "On-Off"
	case 0x01:
		return "Off-On"
	case 0x02:
		return "Toggle"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsOnOff checks if Switchactions equals the value for On-Off (0x00)
func (a Switchactions) IsOnOff() bool { return a == 0x00 }

// SetOnOff sets Switchactions to On-Off (0x00)
func (a *Switchactions) SetOnOff() { *a = 0x00 }

// IsOffOn checks if Switchactions equals the value for Off-On (0x01)
func (a Switchactions) IsOffOn() bool { return a == 0x01 }

// SetOffOn sets Switchactions to Off-On (0x01)
func (a *Switchactions) SetOffOn() { *a = 0x01 }

// IsToggle checks if Switchactions equals the value for Toggle (0x02)
func (a Switchactions) IsToggle() bool { return a == 0x02 }

// SetToggle sets Switchactions to Toggle (0x02)
func (a *Switchactions) SetToggle() { *a = 0x02 }
