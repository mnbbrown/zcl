// The IAS WD cluster provides an interface to the functionality of any Warning Device equipment of the IAS system. Using this cluster, a ZigBee enabled CIE device can access a ZigBee enabled IAS WD device and issue alarm warning indications (siren, strobe lighting, etc.) when a system alarm condition is detected.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// IasWd
const IasWdID zcl.ClusterID = 1282

var IasWdCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		StartWarningCommand: func() zcl.Command { return new(StartWarning) },
		SquawkCommand:       func() zcl.Command { return new(Squawk) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MaxDurationAttr: func() zcl.Attr { return new(MaxDuration) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type StartWarning struct {
	Options         zcl.Zbmp8
	WarningDuration zcl.Zu16
	StrobeDutyCycle zcl.Zu8
	StrobeLevel     zcl.Zenum8
}

const StartWarningCommand zcl.CommandID = 0

func (v *StartWarning) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
		&v.WarningDuration,
		&v.StrobeDutyCycle,
		&v.StrobeLevel,
	}
}

func (v StartWarning) ID() zcl.CommandID {
	return StartWarningCommand
}

func (v StartWarning) Cluster() zcl.ClusterID {
	return IasWdID
}

func (v StartWarning) MnfCode() []byte {
	return []byte{}
}

func (v StartWarning) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.WarningDuration.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StrobeDutyCycle.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StrobeLevel.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *StartWarning) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.WarningDuration).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StrobeDutyCycle).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StrobeLevel).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type Squawk struct {
	Options zcl.Zbmp8
}

const SquawkCommand zcl.CommandID = 1

func (v *Squawk) Values() []zcl.Val {
	return []zcl.Val{
		&v.Options,
	}
}

func (v Squawk) ID() zcl.CommandID {
	return SquawkCommand
}

func (v Squawk) Cluster() zcl.ClusterID {
	return IasWdID
}

func (v Squawk) MnfCode() []byte {
	return []byte{}
}

func (v Squawk) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Options.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *Squawk) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Options).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

const MaxDurationAttr zcl.AttrID = 0

type MaxDuration zcl.Zu16

func (a MaxDuration) ID() zcl.AttrID         { return MaxDurationAttr }
func (a MaxDuration) Cluster() zcl.ClusterID { return IasWdID }
func (a *MaxDuration) Value() *MaxDuration   { return a }
func (a MaxDuration) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxDuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxDuration(*nt)
	return br, err
}

func (a MaxDuration) Readable() bool   { return true }
func (a MaxDuration) Writable() bool   { return true }
func (a MaxDuration) Reportable() bool { return false }
func (a MaxDuration) SceneIndex() int  { return -1 }

func (a MaxDuration) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}
