package smart_energy

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const CurrentMaxDemandDeliveredAttr zcl.AttrID = 2

func (CurrentMaxDemandDelivered) ID() zcl.AttrID   { return CurrentMaxDemandDeliveredAttr }
func (CurrentMaxDemandDelivered) Readable() bool   { return true }
func (CurrentMaxDemandDelivered) Writable() bool   { return false }
func (CurrentMaxDemandDelivered) Reportable() bool { return false }
func (CurrentMaxDemandDelivered) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandDelivered) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandDelivered) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandDelivered) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandDelivered) Name() string        { return `Current Max Demand Delivered` }
func (CurrentMaxDemandDelivered) Description() string { return `` }

type CurrentMaxDemandDelivered zcl.Zu48

func (v *CurrentMaxDemandDelivered) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentMaxDemandDelivered) Value() zcl.Val     { return v }

func (v CurrentMaxDemandDelivered) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentMaxDemandDelivered) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandDelivered(*nt)
	return br, err
}

func (v CurrentMaxDemandDelivered) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentMaxDemandDelivered) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandDelivered(*a)
	return nil
}

func (v *CurrentMaxDemandDelivered) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentMaxDemandDelivered(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandDelivered) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentMaxDemandDeliveredTimeAttr zcl.AttrID = 8

func (CurrentMaxDemandDeliveredTime) ID() zcl.AttrID   { return CurrentMaxDemandDeliveredTimeAttr }
func (CurrentMaxDemandDeliveredTime) Readable() bool   { return true }
func (CurrentMaxDemandDeliveredTime) Writable() bool   { return false }
func (CurrentMaxDemandDeliveredTime) Reportable() bool { return false }
func (CurrentMaxDemandDeliveredTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandDeliveredTime) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandDeliveredTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandDeliveredTime) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandDeliveredTime) Name() string        { return `Current Max Demand Delivered Time` }
func (CurrentMaxDemandDeliveredTime) Description() string { return `` }

type CurrentMaxDemandDeliveredTime zcl.Zutc

func (v *CurrentMaxDemandDeliveredTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *CurrentMaxDemandDeliveredTime) Value() zcl.Val     { return v }

func (v CurrentMaxDemandDeliveredTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *CurrentMaxDemandDeliveredTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandDeliveredTime(*nt)
	return br, err
}

func (v CurrentMaxDemandDeliveredTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *CurrentMaxDemandDeliveredTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandDeliveredTime(*a)
	return nil
}

func (v *CurrentMaxDemandDeliveredTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = CurrentMaxDemandDeliveredTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandDeliveredTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const CurrentMaxDemandReceivedAttr zcl.AttrID = 3

func (CurrentMaxDemandReceived) ID() zcl.AttrID   { return CurrentMaxDemandReceivedAttr }
func (CurrentMaxDemandReceived) Readable() bool   { return true }
func (CurrentMaxDemandReceived) Writable() bool   { return false }
func (CurrentMaxDemandReceived) Reportable() bool { return false }
func (CurrentMaxDemandReceived) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandReceived) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandReceived) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandReceived) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandReceived) Name() string        { return `Current Max Demand Received` }
func (CurrentMaxDemandReceived) Description() string { return `` }

type CurrentMaxDemandReceived zcl.Zu48

func (v *CurrentMaxDemandReceived) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentMaxDemandReceived) Value() zcl.Val     { return v }

func (v CurrentMaxDemandReceived) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentMaxDemandReceived) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandReceived(*nt)
	return br, err
}

func (v CurrentMaxDemandReceived) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentMaxDemandReceived) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandReceived(*a)
	return nil
}

func (v *CurrentMaxDemandReceived) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentMaxDemandReceived(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandReceived) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentMaxDemandReceivedTimeAttr zcl.AttrID = 9

func (CurrentMaxDemandReceivedTime) ID() zcl.AttrID   { return CurrentMaxDemandReceivedTimeAttr }
func (CurrentMaxDemandReceivedTime) Readable() bool   { return true }
func (CurrentMaxDemandReceivedTime) Writable() bool   { return false }
func (CurrentMaxDemandReceivedTime) Reportable() bool { return false }
func (CurrentMaxDemandReceivedTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentMaxDemandReceivedTime) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentMaxDemandReceivedTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentMaxDemandReceivedTime) AttrValue() zcl.Val  { return v.Value() }

func (CurrentMaxDemandReceivedTime) Name() string        { return `Current Max Demand Received Time` }
func (CurrentMaxDemandReceivedTime) Description() string { return `` }

type CurrentMaxDemandReceivedTime zcl.Zutc

func (v *CurrentMaxDemandReceivedTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *CurrentMaxDemandReceivedTime) Value() zcl.Val     { return v }

func (v CurrentMaxDemandReceivedTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *CurrentMaxDemandReceivedTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentMaxDemandReceivedTime(*nt)
	return br, err
}

func (v CurrentMaxDemandReceivedTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *CurrentMaxDemandReceivedTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentMaxDemandReceivedTime(*a)
	return nil
}

func (v *CurrentMaxDemandReceivedTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = CurrentMaxDemandReceivedTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentMaxDemandReceivedTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const CurrentSummationDeliveredAttr zcl.AttrID = 0

func (CurrentSummationDelivered) ID() zcl.AttrID   { return CurrentSummationDeliveredAttr }
func (CurrentSummationDelivered) Readable() bool   { return true }
func (CurrentSummationDelivered) Writable() bool   { return false }
func (CurrentSummationDelivered) Reportable() bool { return false }
func (CurrentSummationDelivered) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentSummationDelivered) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentSummationDelivered) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentSummationDelivered) AttrValue() zcl.Val  { return v.Value() }

func (CurrentSummationDelivered) Name() string        { return `Current Summation Delivered` }
func (CurrentSummationDelivered) Description() string { return `` }

type CurrentSummationDelivered zcl.Zu48

func (v *CurrentSummationDelivered) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentSummationDelivered) Value() zcl.Val     { return v }

func (v CurrentSummationDelivered) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentSummationDelivered) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentSummationDelivered(*nt)
	return br, err
}

func (v CurrentSummationDelivered) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentSummationDelivered) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentSummationDelivered(*a)
	return nil
}

func (v *CurrentSummationDelivered) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentSummationDelivered(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentSummationDelivered) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const CurrentSummationReceivedAttr zcl.AttrID = 1

func (CurrentSummationReceived) ID() zcl.AttrID   { return CurrentSummationReceivedAttr }
func (CurrentSummationReceived) Readable() bool   { return true }
func (CurrentSummationReceived) Writable() bool   { return false }
func (CurrentSummationReceived) Reportable() bool { return false }
func (CurrentSummationReceived) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentSummationReceived) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentSummationReceived) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentSummationReceived) AttrValue() zcl.Val  { return v.Value() }

func (CurrentSummationReceived) Name() string        { return `Current Summation Received` }
func (CurrentSummationReceived) Description() string { return `` }

type CurrentSummationReceived zcl.Zu48

func (v *CurrentSummationReceived) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *CurrentSummationReceived) Value() zcl.Val     { return v }

func (v CurrentSummationReceived) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *CurrentSummationReceived) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentSummationReceived(*nt)
	return br, err
}

func (v CurrentSummationReceived) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *CurrentSummationReceived) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentSummationReceived(*a)
	return nil
}

func (v *CurrentSummationReceived) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = CurrentSummationReceived(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentSummationReceived) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const DftSummationAttr zcl.AttrID = 4

func (DftSummation) ID() zcl.AttrID   { return DftSummationAttr }
func (DftSummation) Readable() bool   { return true }
func (DftSummation) Writable() bool   { return false }
func (DftSummation) Reportable() bool { return false }
func (DftSummation) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DftSummation) AttrID() zcl.AttrID   { return v.ID() }
func (v DftSummation) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DftSummation) AttrValue() zcl.Val  { return v.Value() }

func (DftSummation) Name() string        { return `DFT Summation` }
func (DftSummation) Description() string { return `` }

type DftSummation zcl.Zu48

func (v *DftSummation) TypeID() zcl.TypeID { return new(zcl.Zu48).TypeID() }
func (v *DftSummation) Value() zcl.Val     { return v }

func (v DftSummation) MarshalZcl() ([]byte, error) { return zcl.Zu48(v).MarshalZcl() }

func (v *DftSummation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*v = DftSummation(*nt)
	return br, err
}

func (v DftSummation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu48(v))
}

func (v *DftSummation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu48)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DftSummation(*a)
	return nil
}

func (v *DftSummation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu48); ok {
		*v = DftSummation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DftSummation) String() string {
	return zcl.Sprintf("%v", zcl.Zu48(v))
}

const DailyFreezeTimeAttr zcl.AttrID = 5

func (DailyFreezeTime) ID() zcl.AttrID   { return DailyFreezeTimeAttr }
func (DailyFreezeTime) Readable() bool   { return true }
func (DailyFreezeTime) Writable() bool   { return false }
func (DailyFreezeTime) Reportable() bool { return false }
func (DailyFreezeTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DailyFreezeTime) AttrID() zcl.AttrID   { return v.ID() }
func (v DailyFreezeTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DailyFreezeTime) AttrValue() zcl.Val  { return v.Value() }

func (DailyFreezeTime) Name() string        { return `Daily Freeze Time` }
func (DailyFreezeTime) Description() string { return `` }

type DailyFreezeTime zcl.Zu16

func (v *DailyFreezeTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DailyFreezeTime) Value() zcl.Val     { return v }

func (v DailyFreezeTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DailyFreezeTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DailyFreezeTime(*nt)
	return br, err
}

func (v DailyFreezeTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DailyFreezeTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DailyFreezeTime(*a)
	return nil
}

func (v *DailyFreezeTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DailyFreezeTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DailyFreezeTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DefaultUpdatePeriodAttr zcl.AttrID = 10

func (DefaultUpdatePeriod) ID() zcl.AttrID   { return DefaultUpdatePeriodAttr }
func (DefaultUpdatePeriod) Readable() bool   { return true }
func (DefaultUpdatePeriod) Writable() bool   { return false }
func (DefaultUpdatePeriod) Reportable() bool { return false }
func (DefaultUpdatePeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DefaultUpdatePeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v DefaultUpdatePeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DefaultUpdatePeriod) AttrValue() zcl.Val  { return v.Value() }

func (DefaultUpdatePeriod) Name() string        { return `Default Update Period` }
func (DefaultUpdatePeriod) Description() string { return `` }

type DefaultUpdatePeriod zcl.Zu8

func (v *DefaultUpdatePeriod) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DefaultUpdatePeriod) Value() zcl.Val     { return v }

func (v DefaultUpdatePeriod) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DefaultUpdatePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DefaultUpdatePeriod(*nt)
	return br, err
}

func (v DefaultUpdatePeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DefaultUpdatePeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DefaultUpdatePeriod(*a)
	return nil
}

func (v *DefaultUpdatePeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DefaultUpdatePeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DefaultUpdatePeriod) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const FastPollUpdatePeriodAttr zcl.AttrID = 11

func (FastPollUpdatePeriod) ID() zcl.AttrID   { return FastPollUpdatePeriodAttr }
func (FastPollUpdatePeriod) Readable() bool   { return true }
func (FastPollUpdatePeriod) Writable() bool   { return false }
func (FastPollUpdatePeriod) Reportable() bool { return false }
func (FastPollUpdatePeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FastPollUpdatePeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v FastPollUpdatePeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FastPollUpdatePeriod) AttrValue() zcl.Val  { return v.Value() }

func (FastPollUpdatePeriod) Name() string        { return `Fast Poll Update Period` }
func (FastPollUpdatePeriod) Description() string { return `` }

type FastPollUpdatePeriod zcl.Zu8

func (v *FastPollUpdatePeriod) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *FastPollUpdatePeriod) Value() zcl.Val     { return v }

func (v FastPollUpdatePeriod) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *FastPollUpdatePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = FastPollUpdatePeriod(*nt)
	return br, err
}

func (v FastPollUpdatePeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *FastPollUpdatePeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FastPollUpdatePeriod(*a)
	return nil
}

func (v *FastPollUpdatePeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = FastPollUpdatePeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FastPollUpdatePeriod) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const InstantaneousDemandAttr zcl.AttrID = 1024

func (InstantaneousDemand) ID() zcl.AttrID   { return InstantaneousDemandAttr }
func (InstantaneousDemand) Readable() bool   { return false }
func (InstantaneousDemand) Writable() bool   { return false }
func (InstantaneousDemand) Reportable() bool { return false }
func (InstantaneousDemand) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v InstantaneousDemand) AttrID() zcl.AttrID   { return v.ID() }
func (v InstantaneousDemand) AttrType() zcl.TypeID { return v.TypeID() }
func (v *InstantaneousDemand) AttrValue() zcl.Val  { return v.Value() }

func (InstantaneousDemand) Name() string        { return `Instantaneous Demand` }
func (InstantaneousDemand) Description() string { return `` }

type InstantaneousDemand zcl.Zs24

func (v *InstantaneousDemand) TypeID() zcl.TypeID { return new(zcl.Zs24).TypeID() }
func (v *InstantaneousDemand) Value() zcl.Val     { return v }

func (v InstantaneousDemand) MarshalZcl() ([]byte, error) { return zcl.Zs24(v).MarshalZcl() }

func (v *InstantaneousDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs24)
	br, err := nt.UnmarshalZcl(b)
	*v = InstantaneousDemand(*nt)
	return br, err
}

func (v InstantaneousDemand) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs24(v))
}

func (v *InstantaneousDemand) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = InstantaneousDemand(*a)
	return nil
}

func (v *InstantaneousDemand) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs24); ok {
		*v = InstantaneousDemand(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v InstantaneousDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zs24(v))
}

const PowerFactorAttr zcl.AttrID = 6

func (PowerFactor) ID() zcl.AttrID   { return PowerFactorAttr }
func (PowerFactor) Readable() bool   { return true }
func (PowerFactor) Writable() bool   { return false }
func (PowerFactor) Reportable() bool { return false }
func (PowerFactor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerFactor) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerFactor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerFactor) AttrValue() zcl.Val  { return v.Value() }

func (PowerFactor) Name() string        { return `Power Factor` }
func (PowerFactor) Description() string { return `` }

type PowerFactor zcl.Zs8

func (v *PowerFactor) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *PowerFactor) Value() zcl.Val     { return v }

func (v PowerFactor) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *PowerFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerFactor(*nt)
	return br, err
}

func (v PowerFactor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *PowerFactor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerFactor(*a)
	return nil
}

func (v *PowerFactor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = PowerFactor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerFactor) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const ReadingSnapShotTimeAttr zcl.AttrID = 7

func (ReadingSnapShotTime) ID() zcl.AttrID   { return ReadingSnapShotTimeAttr }
func (ReadingSnapShotTime) Readable() bool   { return true }
func (ReadingSnapShotTime) Writable() bool   { return false }
func (ReadingSnapShotTime) Reportable() bool { return false }
func (ReadingSnapShotTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReadingSnapShotTime) AttrID() zcl.AttrID   { return v.ID() }
func (v ReadingSnapShotTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReadingSnapShotTime) AttrValue() zcl.Val  { return v.Value() }

func (ReadingSnapShotTime) Name() string        { return `Reading Snap Shot Time` }
func (ReadingSnapShotTime) Description() string { return `` }

type ReadingSnapShotTime zcl.Zutc

func (v *ReadingSnapShotTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *ReadingSnapShotTime) Value() zcl.Val     { return v }

func (v ReadingSnapShotTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *ReadingSnapShotTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = ReadingSnapShotTime(*nt)
	return br, err
}

func (v ReadingSnapShotTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *ReadingSnapShotTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReadingSnapShotTime(*a)
	return nil
}

func (v *ReadingSnapShotTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = ReadingSnapShotTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReadingSnapShotTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}
