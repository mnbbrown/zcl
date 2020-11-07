package security

import "hemtjan.st/zcl"

// IasZone
const IasZoneID zcl.ClusterID = 1280

var IasZoneCluster = zcl.Cluster{
	Name:      "IAS Zone",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		ZoneStateAttr:  func() zcl.Attr { return new(ZoneState) },
		ZoneTypeAttr:   func() zcl.Attr { return new(ZoneType) },
		ZoneStatusAttr: func() zcl.Attr { return new(ZoneStatus) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
