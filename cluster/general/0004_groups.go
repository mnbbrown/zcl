// Attributes and commands for group configuration and manipulation.
package general

import (
	"neotor.se/zcl"
)

// Groups
// Attributes and commands for group configuration and manipulation.

func NewGroupsServer(profile zcl.ProfileID) *GroupsServer { return &GroupsServer{p: profile} }
func NewGroupsClient(profile zcl.ProfileID) *GroupsClient { return &GroupsClient{p: profile} }

const GroupsCluster zcl.ClusterID = 4

type GroupsServer struct {
	p zcl.ProfileID

	GroupNameSupport *GroupNameSupport
}

func (s *GroupsServer) AddGroup() *AddGroup                     { return new(AddGroup) }
func (s *GroupsServer) ViewGroup() *ViewGroup                   { return new(ViewGroup) }
func (s *GroupsServer) GetGroupMembership() *GetGroupMembership { return new(GetGroupMembership) }
func (s *GroupsServer) RemoveGroup() *RemoveGroup               { return new(RemoveGroup) }
func (s *GroupsServer) RemoveAllGroups() *RemoveAllGroups       { return new(RemoveAllGroups) }
func (s *GroupsServer) AddGroupIfIdentifying() *AddGroupIfIdentifying {
	return new(AddGroupIfIdentifying)
}

type GroupsClient struct {
	p zcl.ProfileID
}

func (s *GroupsClient) AddGroupResponse() *AddGroupResponse   { return new(AddGroupResponse) }
func (s *GroupsClient) ViewGroupResponse() *ViewGroupResponse { return new(ViewGroupResponse) }
func (s *GroupsClient) GetGroupMembershipResponse() *GetGroupMembershipResponse {
	return new(GetGroupMembershipResponse)
}
func (s *GroupsClient) RemoveGroupResponse() *RemoveGroupResponse { return new(RemoveGroupResponse) }

/*
var GroupsServer = map[zcl.CommandID]func() zcl.Command{
    AddGroupID: func() zcl.Command { return new(AddGroup) },
    ViewGroupID: func() zcl.Command { return new(ViewGroup) },
    GetGroupMembershipID: func() zcl.Command { return new(GetGroupMembership) },
    RemoveGroupID: func() zcl.Command { return new(RemoveGroup) },
    RemoveAllGroupsID: func() zcl.Command { return new(RemoveAllGroups) },
    AddGroupIfIdentifyingID: func() zcl.Command { return new(AddGroupIfIdentifying) },
}

var GroupsClient = map[zcl.CommandID]func() zcl.Command{
    AddGroupResponseID: func() zcl.Command { return new(AddGroupResponse) },
    ViewGroupResponseID: func() zcl.Command { return new(ViewGroupResponse) },
    GetGroupMembershipResponseID: func() zcl.Command { return new(GetGroupMembershipResponse) },
    RemoveGroupResponseID: func() zcl.Command { return new(RemoveGroupResponse) },
}
*/

// Add a group to the device.
type AddGroup struct {
	GroupId   zcl.Zu16
	GroupName zcl.Zcstring
}

const AddGroupCommand zcl.CommandID = 0

func (v *AddGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.GroupName,
	}
}

func (v AddGroup) ID() zcl.CommandID {
	return AddGroupCommand
}

func (v AddGroup) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v AddGroup) MnfCode() []byte {
	return []byte{}
}

func (v AddGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Get the name of a group.
type ViewGroup struct {
	GroupId zcl.Zu16
}

const ViewGroupCommand zcl.CommandID = 1

func (v *ViewGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

func (v ViewGroup) ID() zcl.CommandID {
	return ViewGroupCommand
}

func (v ViewGroup) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v ViewGroup) MnfCode() []byte {
	return []byte{}
}

func (v ViewGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ViewGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Get the group membership of the device. Send an empty group list to request all group memberships
type GetGroupMembership struct {
	GroupList zcl.Zset
}

const GetGroupMembershipCommand zcl.CommandID = 2

func (v *GetGroupMembership) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupList,
	}
}

func (v GetGroupMembership) ID() zcl.CommandID {
	return GetGroupMembershipCommand
}

func (v GetGroupMembership) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v GetGroupMembership) MnfCode() []byte {
	return []byte{}
}

func (v GetGroupMembership) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupList.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetGroupMembership) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Remove a group from the device.
type RemoveGroup struct {
	GroupId zcl.Zu16
}

const RemoveGroupCommand zcl.CommandID = 3

func (v *RemoveGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

func (v RemoveGroup) ID() zcl.CommandID {
	return RemoveGroupCommand
}

func (v RemoveGroup) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v RemoveGroup) MnfCode() []byte {
	return []byte{}
}

func (v RemoveGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RemoveGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Remove all group from the device.
type RemoveAllGroups struct {
}

const RemoveAllGroupsCommand zcl.CommandID = 4

func (v *RemoveAllGroups) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v RemoveAllGroups) ID() zcl.CommandID {
	return RemoveAllGroupsCommand
}

func (v RemoveAllGroups) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v RemoveAllGroups) MnfCode() []byte {
	return []byte{}
}

func (v RemoveAllGroups) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *RemoveAllGroups) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// Add a group to the device if the device is currently identifying itself (using the identify cluster)
type AddGroupIfIdentifying struct {
	GroupId   zcl.Zu16
	GroupName zcl.Zcstring
}

const AddGroupIfIdentifyingCommand zcl.CommandID = 5

func (v *AddGroupIfIdentifying) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.GroupName,
	}
}

func (v AddGroupIfIdentifying) ID() zcl.CommandID {
	return AddGroupIfIdentifyingCommand
}

func (v AddGroupIfIdentifying) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v AddGroupIfIdentifying) MnfCode() []byte {
	return []byte{}
}

func (v AddGroupIfIdentifying) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddGroupIfIdentifying) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// The Response to the add group request.
type AddGroupResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
}

const AddGroupResponseCommand zcl.CommandID = 0

func (v *AddGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

func (v AddGroupResponse) ID() zcl.CommandID {
	return AddGroupResponseCommand
}

func (v AddGroupResponse) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v AddGroupResponse) MnfCode() []byte {
	return []byte{}
}

func (v AddGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// The Response to the view group request.
type ViewGroupResponse struct {
	Status    zcl.Status
	GroupId   zcl.Zu16
	GroupName zcl.Zcstring
}

const ViewGroupResponseCommand zcl.CommandID = 1

func (v *ViewGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.GroupName,
	}
}

func (v ViewGroupResponse) ID() zcl.CommandID {
	return ViewGroupResponseCommand
}

func (v ViewGroupResponse) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v ViewGroupResponse) MnfCode() []byte {
	return []byte{}
}

func (v ViewGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ViewGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

// The Response to the get group membership request.
type GetGroupMembershipResponse struct {
	// The remaining number of groups that can be added.
	// If set to 0xFE, at least one more group can be added (exact number unknown)
	// If set to 0xFF, it's unknown if any more groups can be added
	//
	Capacity  zcl.Zu8
	GroupList zcl.Zset
}

const GetGroupMembershipResponseCommand zcl.CommandID = 2

func (v *GetGroupMembershipResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Capacity,
		&v.GroupList,
	}
}

func (v GetGroupMembershipResponse) ID() zcl.CommandID {
	return GetGroupMembershipResponseCommand
}

func (v GetGroupMembershipResponse) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v GetGroupMembershipResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetGroupMembershipResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Capacity.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupList.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetGroupMembershipResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Capacity).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// The Response to the remove group request.
type RemoveGroupResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
}

const RemoveGroupResponseCommand zcl.CommandID = 3

func (v *RemoveGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

func (v RemoveGroupResponse) ID() zcl.CommandID {
	return RemoveGroupResponseCommand
}

func (v RemoveGroupResponse) Cluster() zcl.ClusterID {
	return GroupsCluster
}

func (v RemoveGroupResponse) MnfCode() []byte {
	return []byte{}
}

func (v RemoveGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RemoveGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type GroupNameSupport zcl.Zbmp8

func (a GroupNameSupport) ID() zcl.AttrID            { return 0 }
func (a GroupNameSupport) Cluster() zcl.ClusterID    { return GroupsCluster }
func (a *GroupNameSupport) Value() *GroupNameSupport { return a }
func (a GroupNameSupport) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *GroupNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupNameSupport(*nt)
	return br, err
}

func (a GroupNameSupport) Readable() bool   { return true }
func (a GroupNameSupport) Writable() bool   { return false }
func (a GroupNameSupport) Reportable() bool { return false }
func (a GroupNameSupport) SceneIndex() int  { return -1 }

func (a GroupNameSupport) String() string {

	var bstr []string
	if a.IsNamesSupported() {
		bstr = append(bstr, "Names Supported")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a GroupNameSupport) IsNamesSupported() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *GroupNameSupport) SetNamesSupported(b bool) {
	*a = GroupNameSupport(zcl.BitmapSet([]byte(*a), 7, b))
}
