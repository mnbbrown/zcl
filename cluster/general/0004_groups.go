package general

import "hemtjan.st/zcl"

// Groups
const GroupsID zcl.ClusterID = 4

var GroupsCluster = zcl.Cluster{
	Name: "Groups",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		AddGroupCommand:              func() zcl.Command { return new(AddGroup) },
		ViewGroupCommand:             func() zcl.Command { return new(ViewGroup) },
		GetGroupMembershipCommand:    func() zcl.Command { return new(GetGroupMembership) },
		RemoveGroupCommand:           func() zcl.Command { return new(RemoveGroup) },
		RemoveAllGroupsCommand:       func() zcl.Command { return new(RemoveAllGroups) },
		AddGroupIfIdentifyingCommand: func() zcl.Command { return new(AddGroupIfIdentifying) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		AddGroupResponseCommand:           func() zcl.Command { return new(AddGroupResponse) },
		ViewGroupResponseCommand:          func() zcl.Command { return new(ViewGroupResponse) },
		GetGroupMembershipResponseCommand: func() zcl.Command { return new(GetGroupMembershipResponse) },
		RemoveGroupResponseCommand:        func() zcl.Command { return new(RemoveGroupResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		GroupNameSupportAttr: func() zcl.Attr { return new(GroupNameSupport) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// AddGroup will add a group to the device
type AddGroup struct {
	GroupId   GroupId
	GroupName GroupName
}

// AddGroupCommand is the Command ID of AddGroup
const AddGroupCommand CommandID = 0x0000

// Values returns all values of AddGroup
func (v *AddGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.GroupName,
	}
}

// Arguments returns all values of AddGroup
func (v *AddGroup) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "GroupName", Argument: &v.GroupName},
	}
}

// Name of the command
func (AddGroup) Name() string { return "Add group" }

// ID of the command
func (AddGroup) ID() CommandID { return AddGroupCommand }

// Required
func (AddGroup) Required() bool { return true }

// Cluster ID of the command
func (AddGroup) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (AddGroup) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (AddGroup) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of AddGroup
func (v AddGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupName.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the AddGroup struct
func (v *AddGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v AddGroup) String() string {
	return zcl.Sprintf(
		"AddGroup{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"GroupName(%v)",
		}, " ")+"}",
		v.GroupId,
		v.GroupName,
	)
}

// ViewGroup requests the name of a group
type ViewGroup struct {
	GroupId GroupId
}

// ViewGroupCommand is the Command ID of ViewGroup
const ViewGroupCommand CommandID = 0x0001

// Values returns all values of ViewGroup
func (v *ViewGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

// Arguments returns all values of ViewGroup
func (v *ViewGroup) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
	}
}

// Name of the command
func (ViewGroup) Name() string { return "View group" }

// ID of the command
func (ViewGroup) ID() CommandID { return ViewGroupCommand }

// Required
func (ViewGroup) Required() bool { return true }

// Cluster ID of the command
func (ViewGroup) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (ViewGroup) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ViewGroup) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of ViewGroup
func (v ViewGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ViewGroup struct
func (v *ViewGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ViewGroup) String() string {
	return zcl.Sprintf(
		"ViewGroup{"+zcl.StrJoin([]string{
			"GroupId(%v)",
		}, " ")+"}",
		v.GroupId,
	)
}

// GetGroupMembership fetches group membership(s). Request with empty list to request all memberships
type GetGroupMembership struct {
	GroupList GroupList
}

// GetGroupMembershipCommand is the Command ID of GetGroupMembership
const GetGroupMembershipCommand CommandID = 0x0002

// Values returns all values of GetGroupMembership
func (v *GetGroupMembership) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupList,
	}
}

// Arguments returns all values of GetGroupMembership
func (v *GetGroupMembership) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupList", Argument: &v.GroupList},
	}
}

// Name of the command
func (GetGroupMembership) Name() string { return "Get group membership" }

// ID of the command
func (GetGroupMembership) ID() CommandID { return GetGroupMembershipCommand }

// Required
func (GetGroupMembership) Required() bool { return true }

// Cluster ID of the command
func (GetGroupMembership) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (GetGroupMembership) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetGroupMembership) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

// MarshalZcl returns the wire format representation of GetGroupMembership
func (v GetGroupMembership) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.GroupList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetGroupMembership struct
func (v *GetGroupMembership) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.GroupList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetGroupMembership) String() string {
	return zcl.Sprintf(
		"GetGroupMembership{"+zcl.StrJoin([]string{
			"GroupList(%v)",
		}, " ")+"}",
		v.GroupList,
	)
}

// RemoveGroup Remove a group from the device.
type RemoveGroup struct {
	GroupList GroupList
}

// RemoveGroupCommand is the Command ID of RemoveGroup
const RemoveGroupCommand CommandID = 0x0003

// Values returns all values of RemoveGroup
func (v *RemoveGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupList,
	}
}

// Arguments returns all values of RemoveGroup
func (v *RemoveGroup) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupList", Argument: &v.GroupList},
	}
}

// Name of the command
func (RemoveGroup) Name() string { return "Remove group" }

// ID of the command
func (RemoveGroup) ID() CommandID { return RemoveGroupCommand }

// Required
func (RemoveGroup) Required() bool { return true }

// Cluster ID of the command
func (RemoveGroup) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveGroup) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RemoveGroup) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

// MarshalZcl returns the wire format representation of RemoveGroup
func (v RemoveGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.GroupList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveGroup struct
func (v *RemoveGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.GroupList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RemoveGroup) String() string {
	return zcl.Sprintf(
		"RemoveGroup{"+zcl.StrJoin([]string{
			"GroupList(%v)",
		}, " ")+"}",
		v.GroupList,
	)
}

// RemoveAllGroups Remove all group from the device.
type RemoveAllGroups struct {
}

// RemoveAllGroupsCommand is the Command ID of RemoveAllGroups
const RemoveAllGroupsCommand CommandID = 0x0004

// Values returns all values of RemoveAllGroups
func (v *RemoveAllGroups) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of RemoveAllGroups
func (v *RemoveAllGroups) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (RemoveAllGroups) Name() string { return "Remove all groups" }

// ID of the command
func (RemoveAllGroups) ID() CommandID { return RemoveAllGroupsCommand }

// Required
func (RemoveAllGroups) Required() bool { return true }

// Cluster ID of the command
func (RemoveAllGroups) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveAllGroups) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RemoveAllGroups) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

// MarshalZcl returns the wire format representation of RemoveAllGroups
func (v RemoveAllGroups) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveAllGroups struct
func (v *RemoveAllGroups) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RemoveAllGroups) String() string {
	return zcl.Sprintf(
		"RemoveAllGroups{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// AddGroupIfIdentifying Add a group to the device if the device is currently identifying itself (using the identify cluster)
type AddGroupIfIdentifying struct {
	GroupId   GroupId
	GroupName GroupName
}

// AddGroupIfIdentifyingCommand is the Command ID of AddGroupIfIdentifying
const AddGroupIfIdentifyingCommand CommandID = 0x0005

// Values returns all values of AddGroupIfIdentifying
func (v *AddGroupIfIdentifying) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.GroupName,
	}
}

// Arguments returns all values of AddGroupIfIdentifying
func (v *AddGroupIfIdentifying) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "GroupName", Argument: &v.GroupName},
	}
}

// Name of the command
func (AddGroupIfIdentifying) Name() string { return "Add group if identifying" }

// ID of the command
func (AddGroupIfIdentifying) ID() CommandID { return AddGroupIfIdentifyingCommand }

// Required
func (AddGroupIfIdentifying) Required() bool { return true }

// Cluster ID of the command
func (AddGroupIfIdentifying) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (AddGroupIfIdentifying) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (AddGroupIfIdentifying) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

// MarshalZcl returns the wire format representation of AddGroupIfIdentifying
func (v AddGroupIfIdentifying) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupName.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the AddGroupIfIdentifying struct
func (v *AddGroupIfIdentifying) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v AddGroupIfIdentifying) String() string {
	return zcl.Sprintf(
		"AddGroupIfIdentifying{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"GroupName(%v)",
		}, " ")+"}",
		v.GroupId,
		v.GroupName,
	)
}

// AddGroupResponse The Response to the add group request.
type AddGroupResponse struct {
	Status  Status
	GroupId GroupId
}

// AddGroupResponseCommand is the Command ID of AddGroupResponse
const AddGroupResponseCommand CommandID = 0x0000

// Values returns all values of AddGroupResponse
func (v *AddGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

// Arguments returns all values of AddGroupResponse
func (v *AddGroupResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
	}
}

// Name of the command
func (AddGroupResponse) Name() string { return "Add group response" }

// ID of the command
func (AddGroupResponse) ID() CommandID { return AddGroupResponseCommand }

// Required
func (AddGroupResponse) Required() bool { return true }

// Cluster ID of the command
func (AddGroupResponse) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (AddGroupResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (AddGroupResponse) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of AddGroupResponse
func (v AddGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the AddGroupResponse struct
func (v *AddGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v AddGroupResponse) String() string {
	return zcl.Sprintf(
		"AddGroupResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
	)
}

// ViewGroupResponse The Response to the view group request.
type ViewGroupResponse struct {
	Status    Status
	GroupId   GroupId
	GroupName GroupName
}

// ViewGroupResponseCommand is the Command ID of ViewGroupResponse
const ViewGroupResponseCommand CommandID = 0x0001

// Values returns all values of ViewGroupResponse
func (v *ViewGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.GroupName,
	}
}

// Arguments returns all values of ViewGroupResponse
func (v *ViewGroupResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "GroupName", Argument: &v.GroupName},
	}
}

// Name of the command
func (ViewGroupResponse) Name() string { return "View group response" }

// ID of the command
func (ViewGroupResponse) ID() CommandID { return ViewGroupResponseCommand }

// Required
func (ViewGroupResponse) Required() bool { return true }

// Cluster ID of the command
func (ViewGroupResponse) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (ViewGroupResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ViewGroupResponse) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of ViewGroupResponse
func (v ViewGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupName.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ViewGroupResponse struct
func (v *ViewGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ViewGroupResponse) String() string {
	return zcl.Sprintf(
		"ViewGroupResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
			"GroupName(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
		v.GroupName,
	)
}

// GetGroupMembershipResponse The Response to the get group membership request.
type GetGroupMembershipResponse struct {
	// GroupCapacity specifies remaining number of groups that can be added.
	// If set to 0xFE, at least one more group can be added (exact number unknown)
	// If set to 0xFF, it's unknown if any more groups can be added
	GroupCapacity GroupCapacity
	GroupList     GroupList
}

// GetGroupMembershipResponseCommand is the Command ID of GetGroupMembershipResponse
const GetGroupMembershipResponseCommand CommandID = 0x0002

// Values returns all values of GetGroupMembershipResponse
func (v *GetGroupMembershipResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupCapacity,
		&v.GroupList,
	}
}

// Arguments returns all values of GetGroupMembershipResponse
func (v *GetGroupMembershipResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupCapacity", Argument: &v.GroupCapacity},
		{Name: "GroupList", Argument: &v.GroupList},
	}
}

// Name of the command
func (GetGroupMembershipResponse) Name() string { return "Get group membership response" }

// ID of the command
func (GetGroupMembershipResponse) ID() CommandID { return GetGroupMembershipResponseCommand }

// Required
func (GetGroupMembershipResponse) Required() bool { return true }

// Cluster ID of the command
func (GetGroupMembershipResponse) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (GetGroupMembershipResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetGroupMembershipResponse) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

// MarshalZcl returns the wire format representation of GetGroupMembershipResponse
func (v GetGroupMembershipResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.GroupCapacity.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetGroupMembershipResponse struct
func (v *GetGroupMembershipResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.GroupCapacity).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetGroupMembershipResponse) String() string {
	return zcl.Sprintf(
		"GetGroupMembershipResponse{"+zcl.StrJoin([]string{
			"GroupCapacity(%v)",
			"GroupList(%v)",
		}, " ")+"}",
		v.GroupCapacity,
		v.GroupList,
	)
}

// RemoveGroupResponse The Response to the remove group request.
type RemoveGroupResponse struct {
	Status  Status
	GroupId GroupId
}

// RemoveGroupResponseCommand is the Command ID of RemoveGroupResponse
const RemoveGroupResponseCommand CommandID = 0x0003

// Values returns all values of RemoveGroupResponse
func (v *RemoveGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

// Arguments returns all values of RemoveGroupResponse
func (v *RemoveGroupResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
	}
}

// Name of the command
func (RemoveGroupResponse) Name() string { return "Remove group response" }

// ID of the command
func (RemoveGroupResponse) ID() CommandID { return RemoveGroupResponseCommand }

// Required
func (RemoveGroupResponse) Required() bool { return true }

// Cluster ID of the command
func (RemoveGroupResponse) Cluster() zcl.ClusterID { return GroupsID }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveGroupResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RemoveGroupResponse) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

// MarshalZcl returns the wire format representation of RemoveGroupResponse
func (v RemoveGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveGroupResponse struct
func (v *RemoveGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RemoveGroupResponse) String() string {
	return zcl.Sprintf(
		"RemoveGroupResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
	)
}
