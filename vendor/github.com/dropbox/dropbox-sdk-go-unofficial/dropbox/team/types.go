// Copyright (c) Dropbox, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package team : has no documentation (yet)
package team

import (
	"encoding/json"
	"time"

	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/properties"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_common"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_policies"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/users"
)

// DeviceSession : has no documentation (yet)
type DeviceSession struct {
	// SessionId : The session id
	SessionId string `json:"session_id"`
	// IpAddress : The IP address of the last activity from this session
	IpAddress string `json:"ip_address,omitempty"`
	// Country : The country from which the last activity from this session was
	// made
	Country string `json:"country,omitempty"`
	// Created : The time this session was created
	Created time.Time `json:"created,omitempty"`
	// Updated : The time of the last activity from this session
	Updated time.Time `json:"updated,omitempty"`
}

// NewDeviceSession returns a new DeviceSession instance
func NewDeviceSession(SessionId string) *DeviceSession {
	s := new(DeviceSession)
	s.SessionId = SessionId
	return s
}

// ActiveWebSession : Information on active web sessions
type ActiveWebSession struct {
	DeviceSession
	// UserAgent : Information on the hosting device
	UserAgent string `json:"user_agent"`
	// Os : Information on the hosting operating system
	Os string `json:"os"`
	// Browser : Information on the browser used for this web session
	Browser string `json:"browser"`
}

// NewActiveWebSession returns a new ActiveWebSession instance
func NewActiveWebSession(SessionId string, UserAgent string, Os string, Browser string) *ActiveWebSession {
	s := new(ActiveWebSession)
	s.SessionId = SessionId
	s.UserAgent = UserAgent
	s.Os = Os
	s.Browser = Browser
	return s
}

// AddPropertyTemplateArg : Arguments for adding property templates.
type AddPropertyTemplateArg struct {
	properties.PropertyGroupTemplate
}

// NewAddPropertyTemplateArg returns a new AddPropertyTemplateArg instance
func NewAddPropertyTemplateArg(Name string, Description string, Fields []*properties.PropertyFieldTemplate) *AddPropertyTemplateArg {
	s := new(AddPropertyTemplateArg)
	s.Name = Name
	s.Description = Description
	s.Fields = Fields
	return s
}

// AddPropertyTemplateResult : has no documentation (yet)
type AddPropertyTemplateResult struct {
	// TemplateId : An identifier for property template added by
	// `propertiesTemplateAdd`.
	TemplateId string `json:"template_id"`
}

// NewAddPropertyTemplateResult returns a new AddPropertyTemplateResult instance
func NewAddPropertyTemplateResult(TemplateId string) *AddPropertyTemplateResult {
	s := new(AddPropertyTemplateResult)
	s.TemplateId = TemplateId
	return s
}

// AdminTier : Describes which team-related admin permissions a user has.
type AdminTier struct {
	dropbox.Tagged
}

// Valid tag values for AdminTier
const (
	AdminTierTeamAdmin           = "team_admin"
	AdminTierUserManagementAdmin = "user_management_admin"
	AdminTierSupportAdmin        = "support_admin"
	AdminTierMemberOnly          = "member_only"
)

// GroupCreateArg : has no documentation (yet)
type GroupCreateArg struct {
	// GroupName : Group name.
	GroupName string `json:"group_name"`
	// GroupExternalId : The creator of a team can associate an arbitrary
	// external ID to the group.
	GroupExternalId string `json:"group_external_id,omitempty"`
}

// NewGroupCreateArg returns a new GroupCreateArg instance
func NewGroupCreateArg(GroupName string) *GroupCreateArg {
	s := new(GroupCreateArg)
	s.GroupName = GroupName
	return s
}

// AlphaGroupCreateArg : has no documentation (yet)
type AlphaGroupCreateArg struct {
	GroupCreateArg
	// GroupManagementType : Whether the team can be managed by selected users,
	// or only by team admins
	GroupManagementType *team_common.GroupManagementType `json:"group_management_type"`
}

// NewAlphaGroupCreateArg returns a new AlphaGroupCreateArg instance
func NewAlphaGroupCreateArg(GroupName string) *AlphaGroupCreateArg {
	s := new(AlphaGroupCreateArg)
	s.GroupName = GroupName
	s.GroupManagementType = &team_common.GroupManagementType{Tagged: dropbox.Tagged{"company_managed"}}
	return s
}

// AlphaGroupFullInfo : Full description of a group.
type AlphaGroupFullInfo struct {
	team_common.AlphaGroupSummary
	// Members : List of group members.
	Members []*GroupMemberInfo `json:"members,omitempty"`
	// Created : The group creation time as a UTC timestamp in milliseconds
	// since the Unix epoch.
	Created uint64 `json:"created"`
}

// NewAlphaGroupFullInfo returns a new AlphaGroupFullInfo instance
func NewAlphaGroupFullInfo(GroupName string, GroupId string, GroupManagementType *team_common.GroupManagementType, Created uint64) *AlphaGroupFullInfo {
	s := new(AlphaGroupFullInfo)
	s.GroupName = GroupName
	s.GroupId = GroupId
	s.GroupManagementType = GroupManagementType
	s.Created = Created
	return s
}

// IncludeMembersArg : has no documentation (yet)
type IncludeMembersArg struct {
	// ReturnMembers : Whether to return the list of members in the group.  Note
	// that the default value will cause all the group members  to be returned
	// in the response. This may take a long time for large groups.
	ReturnMembers bool `json:"return_members"`
}

// NewIncludeMembersArg returns a new IncludeMembersArg instance
func NewIncludeMembersArg() *IncludeMembersArg {
	s := new(IncludeMembersArg)
	s.ReturnMembers = true
	return s
}

// GroupUpdateArgs : has no documentation (yet)
type GroupUpdateArgs struct {
	IncludeMembersArg
	// Group : Specify a group.
	Group *GroupSelector `json:"group"`
	// NewGroupName : Optional argument. Set group name to this if provided.
	NewGroupName string `json:"new_group_name,omitempty"`
	// NewGroupExternalId : Optional argument. New group external ID. If the
	// argument is None, the group's external_id won't be updated. If the
	// argument is empty string, the group's external id will be cleared.
	NewGroupExternalId string `json:"new_group_external_id,omitempty"`
}

// NewGroupUpdateArgs returns a new GroupUpdateArgs instance
func NewGroupUpdateArgs(Group *GroupSelector) *GroupUpdateArgs {
	s := new(GroupUpdateArgs)
	s.Group = Group
	s.ReturnMembers = true
	return s
}

// AlphaGroupUpdateArgs : has no documentation (yet)
type AlphaGroupUpdateArgs struct {
	GroupUpdateArgs
	// NewGroupManagementType : Set new group management type, if provided.
	NewGroupManagementType *team_common.GroupManagementType `json:"new_group_management_type,omitempty"`
}

// NewAlphaGroupUpdateArgs returns a new AlphaGroupUpdateArgs instance
func NewAlphaGroupUpdateArgs(Group *GroupSelector) *AlphaGroupUpdateArgs {
	s := new(AlphaGroupUpdateArgs)
	s.Group = Group
	s.ReturnMembers = true
	return s
}

// AlphaGroupsGetInfoItem : has no documentation (yet)
type AlphaGroupsGetInfoItem struct {
	dropbox.Tagged
	// IdNotFound : An ID that was provided as a parameter to
	// `alphaGroupsGetInfo`, and did not match a corresponding group. The ID can
	// be a group ID, or an external ID, depending on how the method was called.
	IdNotFound string `json:"id_not_found,omitempty"`
	// GroupInfo : Info about a group.
	GroupInfo *AlphaGroupFullInfo `json:"group_info,omitempty"`
}

// Valid tag values for AlphaGroupsGetInfoItem
const (
	AlphaGroupsGetInfoItemIdNotFound = "id_not_found"
	AlphaGroupsGetInfoItemGroupInfo  = "group_info"
)

// UnmarshalJSON deserializes into a AlphaGroupsGetInfoItem instance
func (u *AlphaGroupsGetInfoItem) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// GroupInfo : Info about a group.
		GroupInfo json.RawMessage `json:"group_info,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "id_not_found":
		if err := json.Unmarshal(body, &u.IdNotFound); err != nil {
			return err
		}

	case "group_info":
		if err := json.Unmarshal(body, &u.GroupInfo); err != nil {
			return err
		}

	}
	return nil
}

// AlphaGroupsListResult : has no documentation (yet)
type AlphaGroupsListResult struct {
	// Groups : has no documentation (yet)
	Groups []*team_common.AlphaGroupSummary `json:"groups"`
	// Cursor : Pass the cursor into `alphaGroupsListContinue` to obtain the
	// additional groups.
	Cursor string `json:"cursor"`
	// HasMore : Is true if there are additional groups that have not been
	// returned yet. An additional call to `alphaGroupsListContinue` can
	// retrieve them.
	HasMore bool `json:"has_more"`
}

// NewAlphaGroupsListResult returns a new AlphaGroupsListResult instance
func NewAlphaGroupsListResult(Groups []*team_common.AlphaGroupSummary, Cursor string, HasMore bool) *AlphaGroupsListResult {
	s := new(AlphaGroupsListResult)
	s.Groups = Groups
	s.Cursor = Cursor
	s.HasMore = HasMore
	return s
}

// ApiApp : Information on linked third party applications
type ApiApp struct {
	// AppId : The application unique id
	AppId string `json:"app_id"`
	// AppName : The application name
	AppName string `json:"app_name"`
	// Publisher : The application publisher name
	Publisher string `json:"publisher,omitempty"`
	// PublisherUrl : The publisher's URL
	PublisherUrl string `json:"publisher_url,omitempty"`
	// Linked : The time this application was linked
	Linked time.Time `json:"linked,omitempty"`
	// IsAppFolder : Whether the linked application uses a dedicated folder
	IsAppFolder bool `json:"is_app_folder"`
}

// NewApiApp returns a new ApiApp instance
func NewApiApp(AppId string, AppName string, IsAppFolder bool) *ApiApp {
	s := new(ApiApp)
	s.AppId = AppId
	s.AppName = AppName
	s.IsAppFolder = IsAppFolder
	return s
}

// BaseDfbReport : Base report structure.
type BaseDfbReport struct {
	// StartDate : First date present in the results as 'YYYY-MM-DD' or None.
	StartDate string `json:"start_date"`
}

// NewBaseDfbReport returns a new BaseDfbReport instance
func NewBaseDfbReport(StartDate string) *BaseDfbReport {
	s := new(BaseDfbReport)
	s.StartDate = StartDate
	return s
}

// DateRange : Input arguments that can be provided for most reports.
type DateRange struct {
	// StartDate : Optional starting date (inclusive)
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate : Optional ending date (exclusive)
	EndDate time.Time `json:"end_date,omitempty"`
}

// NewDateRange returns a new DateRange instance
func NewDateRange() *DateRange {
	s := new(DateRange)
	return s
}

// DateRangeError : Errors that can originate from problems in input arguments
// to reports.
type DateRangeError struct {
	dropbox.Tagged
}

// Valid tag values for DateRangeError
const (
	DateRangeErrorOther = "other"
)

// DesktopClientSession : Information about linked Dropbox desktop client
// sessions
type DesktopClientSession struct {
	DeviceSession
	// HostName : Name of the hosting desktop
	HostName string `json:"host_name"`
	// ClientType : The Dropbox desktop client type
	ClientType *DesktopPlatform `json:"client_type"`
	// ClientVersion : The Dropbox client version
	ClientVersion string `json:"client_version"`
	// Platform : Information on the hosting platform
	Platform string `json:"platform"`
	// IsDeleteOnUnlinkSupported : Whether it's possible to delete all of the
	// account files upon unlinking
	IsDeleteOnUnlinkSupported bool `json:"is_delete_on_unlink_supported"`
}

// NewDesktopClientSession returns a new DesktopClientSession instance
func NewDesktopClientSession(SessionId string, HostName string, ClientType *DesktopPlatform, ClientVersion string, Platform string, IsDeleteOnUnlinkSupported bool) *DesktopClientSession {
	s := new(DesktopClientSession)
	s.SessionId = SessionId
	s.HostName = HostName
	s.ClientType = ClientType
	s.ClientVersion = ClientVersion
	s.Platform = Platform
	s.IsDeleteOnUnlinkSupported = IsDeleteOnUnlinkSupported
	return s
}

// DesktopPlatform : has no documentation (yet)
type DesktopPlatform struct {
	dropbox.Tagged
}

// Valid tag values for DesktopPlatform
const (
	DesktopPlatformWindows = "windows"
	DesktopPlatformMac     = "mac"
	DesktopPlatformLinux   = "linux"
	DesktopPlatformOther   = "other"
)

// DeviceSessionArg : has no documentation (yet)
type DeviceSessionArg struct {
	// SessionId : The session id
	SessionId string `json:"session_id"`
	// TeamMemberId : The unique id of the member owning the device
	TeamMemberId string `json:"team_member_id"`
}

// NewDeviceSessionArg returns a new DeviceSessionArg instance
func NewDeviceSessionArg(SessionId string, TeamMemberId string) *DeviceSessionArg {
	s := new(DeviceSessionArg)
	s.SessionId = SessionId
	s.TeamMemberId = TeamMemberId
	return s
}

// DevicesActive : Each of the items is an array of values, one value per day.
// The value is the number of devices active within a time window, ending with
// that day. If there is no data for a day, then the value will be None.
type DevicesActive struct {
	// Windows : Array of number of linked windows (desktop) clients with
	// activity.
	Windows []uint64 `json:"windows"`
	// Macos : Array of number of linked mac (desktop) clients with activity.
	Macos []uint64 `json:"macos"`
	// Linux : Array of number of linked linus (desktop) clients with activity.
	Linux []uint64 `json:"linux"`
	// Ios : Array of number of linked ios devices with activity.
	Ios []uint64 `json:"ios"`
	// Android : Array of number of linked android devices with activity.
	Android []uint64 `json:"android"`
	// Other : Array of number of other linked devices (blackberry, windows
	// phone, etc)  with activity.
	Other []uint64 `json:"other"`
	// Total : Array of total number of linked clients with activity.
	Total []uint64 `json:"total"`
}

// NewDevicesActive returns a new DevicesActive instance
func NewDevicesActive(Windows []uint64, Macos []uint64, Linux []uint64, Ios []uint64, Android []uint64, Other []uint64, Total []uint64) *DevicesActive {
	s := new(DevicesActive)
	s.Windows = Windows
	s.Macos = Macos
	s.Linux = Linux
	s.Ios = Ios
	s.Android = Android
	s.Other = Other
	s.Total = Total
	return s
}

// GetActivityReport : Activity Report Result. Each of the items in the storage
// report is an array of values, one value per day. If there is no data for a
// day, then the value will be None.
type GetActivityReport struct {
	BaseDfbReport
	// Adds : Array of total number of adds by team members.
	Adds []uint64 `json:"adds"`
	// Edits : Array of number of edits by team members. If the same user edits
	// the same file multiple times this is counted as a single edit.
	Edits []uint64 `json:"edits"`
	// Deletes : Array of total number of deletes by team members.
	Deletes []uint64 `json:"deletes"`
	// ActiveUsers28Day : Array of the number of users who have been active in
	// the last 28 days.
	ActiveUsers28Day []uint64 `json:"active_users_28_day"`
	// ActiveUsers7Day : Array of the number of users who have been active in
	// the last week.
	ActiveUsers7Day []uint64 `json:"active_users_7_day"`
	// ActiveUsers1Day : Array of the number of users who have been active in
	// the last day.
	ActiveUsers1Day []uint64 `json:"active_users_1_day"`
	// ActiveSharedFolders28Day : Array of the number of shared folders with
	// some activity in the last 28 days.
	ActiveSharedFolders28Day []uint64 `json:"active_shared_folders_28_day"`
	// ActiveSharedFolders7Day : Array of the number of shared folders with some
	// activity in the last week.
	ActiveSharedFolders7Day []uint64 `json:"active_shared_folders_7_day"`
	// ActiveSharedFolders1Day : Array of the number of shared folders with some
	// activity in the last day.
	ActiveSharedFolders1Day []uint64 `json:"active_shared_folders_1_day"`
	// SharedLinksCreated : Array of the number of shared links created.
	SharedLinksCreated []uint64 `json:"shared_links_created"`
	// SharedLinksViewedByTeam : Array of the number of views by team users to
	// shared links created by the team.
	SharedLinksViewedByTeam []uint64 `json:"shared_links_viewed_by_team"`
	// SharedLinksViewedByOutsideUser : Array of the number of views by users
	// outside of the team to shared links created by the team.
	SharedLinksViewedByOutsideUser []uint64 `json:"shared_links_viewed_by_outside_user"`
	// SharedLinksViewedByNotLoggedIn : Array of the number of views by
	// non-logged-in users to shared links created by the team.
	SharedLinksViewedByNotLoggedIn []uint64 `json:"shared_links_viewed_by_not_logged_in"`
	// SharedLinksViewedTotal : Array of the total number of views to shared
	// links created by the team.
	SharedLinksViewedTotal []uint64 `json:"shared_links_viewed_total"`
}

// NewGetActivityReport returns a new GetActivityReport instance
func NewGetActivityReport(StartDate string, Adds []uint64, Edits []uint64, Deletes []uint64, ActiveUsers28Day []uint64, ActiveUsers7Day []uint64, ActiveUsers1Day []uint64, ActiveSharedFolders28Day []uint64, ActiveSharedFolders7Day []uint64, ActiveSharedFolders1Day []uint64, SharedLinksCreated []uint64, SharedLinksViewedByTeam []uint64, SharedLinksViewedByOutsideUser []uint64, SharedLinksViewedByNotLoggedIn []uint64, SharedLinksViewedTotal []uint64) *GetActivityReport {
	s := new(GetActivityReport)
	s.StartDate = StartDate
	s.Adds = Adds
	s.Edits = Edits
	s.Deletes = Deletes
	s.ActiveUsers28Day = ActiveUsers28Day
	s.ActiveUsers7Day = ActiveUsers7Day
	s.ActiveUsers1Day = ActiveUsers1Day
	s.ActiveSharedFolders28Day = ActiveSharedFolders28Day
	s.ActiveSharedFolders7Day = ActiveSharedFolders7Day
	s.ActiveSharedFolders1Day = ActiveSharedFolders1Day
	s.SharedLinksCreated = SharedLinksCreated
	s.SharedLinksViewedByTeam = SharedLinksViewedByTeam
	s.SharedLinksViewedByOutsideUser = SharedLinksViewedByOutsideUser
	s.SharedLinksViewedByNotLoggedIn = SharedLinksViewedByNotLoggedIn
	s.SharedLinksViewedTotal = SharedLinksViewedTotal
	return s
}

// GetDevicesReport : Devices Report Result. Contains subsections for different
// time ranges of activity. Each of the items in each subsection of the storage
// report is an array of values, one value per day. If there is no data for a
// day, then the value will be None.
type GetDevicesReport struct {
	BaseDfbReport
	// Active1Day : Report of the number of devices active in the last day.
	Active1Day *DevicesActive `json:"active_1_day"`
	// Active7Day : Report of the number of devices active in the last 7 days.
	Active7Day *DevicesActive `json:"active_7_day"`
	// Active28Day : Report of the number of devices active in the last 28 days.
	Active28Day *DevicesActive `json:"active_28_day"`
}

// NewGetDevicesReport returns a new GetDevicesReport instance
func NewGetDevicesReport(StartDate string, Active1Day *DevicesActive, Active7Day *DevicesActive, Active28Day *DevicesActive) *GetDevicesReport {
	s := new(GetDevicesReport)
	s.StartDate = StartDate
	s.Active1Day = Active1Day
	s.Active7Day = Active7Day
	s.Active28Day = Active28Day
	return s
}

// GetMembershipReport : Membership Report Result. Each of the items in the
// storage report is an array of values, one value per day. If there is no data
// for a day, then the value will be None.
type GetMembershipReport struct {
	BaseDfbReport
	// TeamSize : Team size, for each day.
	TeamSize []uint64 `json:"team_size"`
	// PendingInvites : The number of pending invites to the team, for each day.
	PendingInvites []uint64 `json:"pending_invites"`
	// MembersJoined : The number of members that joined the team, for each day.
	MembersJoined []uint64 `json:"members_joined"`
	// SuspendedMembers : The number of suspended team members, for each day.
	SuspendedMembers []uint64 `json:"suspended_members"`
	// Licenses : The total number of licenses the team has, for each day.
	Licenses []uint64 `json:"licenses"`
}

// NewGetMembershipReport returns a new GetMembershipReport instance
func NewGetMembershipReport(StartDate string, TeamSize []uint64, PendingInvites []uint64, MembersJoined []uint64, SuspendedMembers []uint64, Licenses []uint64) *GetMembershipReport {
	s := new(GetMembershipReport)
	s.StartDate = StartDate
	s.TeamSize = TeamSize
	s.PendingInvites = PendingInvites
	s.MembersJoined = MembersJoined
	s.SuspendedMembers = SuspendedMembers
	s.Licenses = Licenses
	return s
}

// GetStorageReport : Storage Report Result. Each of the items in the storage
// report is an array of values, one value per day. If there is no data for a
// day, then the value will be None.
type GetStorageReport struct {
	BaseDfbReport
	// TotalUsage : Sum of the shared, unshared, and datastore usages, for each
	// day.
	TotalUsage []uint64 `json:"total_usage"`
	// SharedUsage : Array of the combined size (bytes) of team members' shared
	// folders, for each day.
	SharedUsage []uint64 `json:"shared_usage"`
	// UnsharedUsage : Array of the combined size (bytes) of team members' root
	// namespaces, for each day.
	UnsharedUsage []uint64 `json:"unshared_usage"`
	// SharedFolders : Array of the number of shared folders owned by team
	// members, for each day.
	SharedFolders []uint64 `json:"shared_folders"`
	// MemberStorageMap : Array of storage summaries of team members' account
	// sizes. Each storage summary is an array of key, value pairs, where each
	// pair describes a storage bucket. The key indicates the upper bound of the
	// bucket and the value is the number of users in that bucket. There is one
	// such summary per day. If there is no data for a day, the storage summary
	// will be empty.
	MemberStorageMap [][]*StorageBucket `json:"member_storage_map"`
}

// NewGetStorageReport returns a new GetStorageReport instance
func NewGetStorageReport(StartDate string, TotalUsage []uint64, SharedUsage []uint64, UnsharedUsage []uint64, SharedFolders []uint64, MemberStorageMap [][]*StorageBucket) *GetStorageReport {
	s := new(GetStorageReport)
	s.StartDate = StartDate
	s.TotalUsage = TotalUsage
	s.SharedUsage = SharedUsage
	s.UnsharedUsage = UnsharedUsage
	s.SharedFolders = SharedFolders
	s.MemberStorageMap = MemberStorageMap
	return s
}

// GroupAccessType : Role of a user in group.
type GroupAccessType struct {
	dropbox.Tagged
}

// Valid tag values for GroupAccessType
const (
	GroupAccessTypeMember = "member"
	GroupAccessTypeOwner  = "owner"
)

// GroupCreateError : has no documentation (yet)
type GroupCreateError struct {
	dropbox.Tagged
}

// Valid tag values for GroupCreateError
const (
	GroupCreateErrorGroupNameAlreadyUsed   = "group_name_already_used"
	GroupCreateErrorGroupNameInvalid       = "group_name_invalid"
	GroupCreateErrorExternalIdAlreadyInUse = "external_id_already_in_use"
	GroupCreateErrorOther                  = "other"
)

// GroupSelectorError : Error that can be raised when `GroupSelector` is used.
type GroupSelectorError struct {
	dropbox.Tagged
}

// Valid tag values for GroupSelectorError
const (
	GroupSelectorErrorGroupNotFound = "group_not_found"
	GroupSelectorErrorOther         = "other"
)

// GroupDeleteError : has no documentation (yet)
type GroupDeleteError struct {
	dropbox.Tagged
}

// Valid tag values for GroupDeleteError
const (
	GroupDeleteErrorGroupAlreadyDeleted = "group_already_deleted"
)

// GroupFullInfo : Full description of a group.
type GroupFullInfo struct {
	team_common.GroupSummary
	// Members : List of group members.
	Members []*GroupMemberInfo `json:"members,omitempty"`
	// Created : The group creation time as a UTC timestamp in milliseconds
	// since the Unix epoch.
	Created uint64 `json:"created"`
}

// NewGroupFullInfo returns a new GroupFullInfo instance
func NewGroupFullInfo(GroupName string, GroupId string, Created uint64) *GroupFullInfo {
	s := new(GroupFullInfo)
	s.GroupName = GroupName
	s.GroupId = GroupId
	s.Created = Created
	return s
}

// GroupMemberInfo : Profile of group member, and role in group.
type GroupMemberInfo struct {
	// Profile : Profile of group member.
	Profile *MemberProfile `json:"profile"`
	// AccessType : The role that the user has in the group.
	AccessType *GroupAccessType `json:"access_type"`
}

// NewGroupMemberInfo returns a new GroupMemberInfo instance
func NewGroupMemberInfo(Profile *MemberProfile, AccessType *GroupAccessType) *GroupMemberInfo {
	s := new(GroupMemberInfo)
	s.Profile = Profile
	s.AccessType = AccessType
	return s
}

// GroupMemberSelector : Argument for selecting a group and a single user.
type GroupMemberSelector struct {
	// Group : Specify a group.
	Group *GroupSelector `json:"group"`
	// User : Identity of a user that is a member of `group`.
	User *UserSelectorArg `json:"user"`
}

// NewGroupMemberSelector returns a new GroupMemberSelector instance
func NewGroupMemberSelector(Group *GroupSelector, User *UserSelectorArg) *GroupMemberSelector {
	s := new(GroupMemberSelector)
	s.Group = Group
	s.User = User
	return s
}

// GroupMemberSelectorError : Error that can be raised when
// `GroupMemberSelector` is used, and the user is required to be a member of the
// specified group.
type GroupMemberSelectorError struct {
	dropbox.Tagged
}

// Valid tag values for GroupMemberSelectorError
const (
	GroupMemberSelectorErrorMemberNotInGroup = "member_not_in_group"
)

// GroupMemberSetAccessTypeError : has no documentation (yet)
type GroupMemberSetAccessTypeError struct {
	dropbox.Tagged
}

// Valid tag values for GroupMemberSetAccessTypeError
const (
	GroupMemberSetAccessTypeErrorUserCannotBeManagerOfCompanyManagedGroup = "user_cannot_be_manager_of_company_managed_group"
)

// GroupMembersAddArg : has no documentation (yet)
type GroupMembersAddArg struct {
	IncludeMembersArg
	// Group : Group to which users will be added.
	Group *GroupSelector `json:"group"`
	// Members : List of users to be added to the group.
	Members []*MemberAccess `json:"members"`
}

// NewGroupMembersAddArg returns a new GroupMembersAddArg instance
func NewGroupMembersAddArg(Group *GroupSelector, Members []*MemberAccess) *GroupMembersAddArg {
	s := new(GroupMembersAddArg)
	s.Group = Group
	s.Members = Members
	s.ReturnMembers = true
	return s
}

// GroupMembersAddError : has no documentation (yet)
type GroupMembersAddError struct {
	dropbox.Tagged
	// MembersNotInTeam : These members are not part of your team. Currently,
	// you cannot add members to a group if they are not part of your team,
	// though this may change in a subsequent version. To add new members to
	// your Dropbox Business team, use the `membersAdd` endpoint.
	MembersNotInTeam []string `json:"members_not_in_team,omitempty"`
	// UsersNotFound : These users were not found in Dropbox.
	UsersNotFound []string `json:"users_not_found,omitempty"`
	// UserCannotBeManagerOfCompanyManagedGroup : A company-managed group cannot
	// be managed by a user.
	UserCannotBeManagerOfCompanyManagedGroup []string `json:"user_cannot_be_manager_of_company_managed_group,omitempty"`
}

// Valid tag values for GroupMembersAddError
const (
	GroupMembersAddErrorDuplicateUser                            = "duplicate_user"
	GroupMembersAddErrorGroupNotInTeam                           = "group_not_in_team"
	GroupMembersAddErrorMembersNotInTeam                         = "members_not_in_team"
	GroupMembersAddErrorUsersNotFound                            = "users_not_found"
	GroupMembersAddErrorUserMustBeActiveToBeOwner                = "user_must_be_active_to_be_owner"
	GroupMembersAddErrorUserCannotBeManagerOfCompanyManagedGroup = "user_cannot_be_manager_of_company_managed_group"
)

// UnmarshalJSON deserializes into a GroupMembersAddError instance
func (u *GroupMembersAddError) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// MembersNotInTeam : These members are not part of your team.
		// Currently, you cannot add members to a group if they are not part of
		// your team, though this may change in a subsequent version. To add new
		// members to your Dropbox Business team, use the `membersAdd` endpoint.
		MembersNotInTeam json.RawMessage `json:"members_not_in_team,omitempty"`
		// UsersNotFound : These users were not found in Dropbox.
		UsersNotFound json.RawMessage `json:"users_not_found,omitempty"`
		// UserCannotBeManagerOfCompanyManagedGroup : A company-managed group
		// cannot be managed by a user.
		UserCannotBeManagerOfCompanyManagedGroup json.RawMessage `json:"user_cannot_be_manager_of_company_managed_group,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "members_not_in_team":
		if err := json.Unmarshal(body, &u.MembersNotInTeam); err != nil {
			return err
		}

	case "users_not_found":
		if err := json.Unmarshal(body, &u.UsersNotFound); err != nil {
			return err
		}

	case "user_cannot_be_manager_of_company_managed_group":
		if err := json.Unmarshal(body, &u.UserCannotBeManagerOfCompanyManagedGroup); err != nil {
			return err
		}

	}
	return nil
}

// GroupMembersChangeResult : Result returned by `groupsMembersAdd` and
// `groupsMembersRemove`.
type GroupMembersChangeResult struct {
	// GroupInfo : The group info after member change operation has been
	// performed.
	GroupInfo *GroupFullInfo `json:"group_info"`
	// AsyncJobId : An ID that can be used to obtain the status of
	// granting/revoking group-owned resources.
	AsyncJobId string `json:"async_job_id"`
}

// NewGroupMembersChangeResult returns a new GroupMembersChangeResult instance
func NewGroupMembersChangeResult(GroupInfo *GroupFullInfo, AsyncJobId string) *GroupMembersChangeResult {
	s := new(GroupMembersChangeResult)
	s.GroupInfo = GroupInfo
	s.AsyncJobId = AsyncJobId
	return s
}

// GroupMembersRemoveArg : has no documentation (yet)
type GroupMembersRemoveArg struct {
	IncludeMembersArg
	// Group : Group from which users will be removed.
	Group *GroupSelector `json:"group"`
	// Users : List of users to be removed from the group.
	Users []*UserSelectorArg `json:"users"`
}

// NewGroupMembersRemoveArg returns a new GroupMembersRemoveArg instance
func NewGroupMembersRemoveArg(Group *GroupSelector, Users []*UserSelectorArg) *GroupMembersRemoveArg {
	s := new(GroupMembersRemoveArg)
	s.Group = Group
	s.Users = Users
	s.ReturnMembers = true
	return s
}

// GroupMembersSelectorError : Error that can be raised when
// `GroupMembersSelector` is used, and the users are required to be members of
// the specified group.
type GroupMembersSelectorError struct {
	dropbox.Tagged
}

// Valid tag values for GroupMembersSelectorError
const (
	GroupMembersSelectorErrorMemberNotInGroup = "member_not_in_group"
)

// GroupMembersRemoveError : has no documentation (yet)
type GroupMembersRemoveError struct {
	dropbox.Tagged
}

// Valid tag values for GroupMembersRemoveError
const (
	GroupMembersRemoveErrorGroupNotInTeam = "group_not_in_team"
)

// GroupMembersSelector : Argument for selecting a group and a list of users.
type GroupMembersSelector struct {
	// Group : Specify a group.
	Group *GroupSelector `json:"group"`
	// Users : A list of users that are members of `group`.
	Users *UsersSelectorArg `json:"users"`
}

// NewGroupMembersSelector returns a new GroupMembersSelector instance
func NewGroupMembersSelector(Group *GroupSelector, Users *UsersSelectorArg) *GroupMembersSelector {
	s := new(GroupMembersSelector)
	s.Group = Group
	s.Users = Users
	return s
}

// GroupMembersSetAccessTypeArg : has no documentation (yet)
type GroupMembersSetAccessTypeArg struct {
	GroupMemberSelector
	// AccessType : New group access type the user will have.
	AccessType *GroupAccessType `json:"access_type"`
	// ReturnMembers : Whether to return the list of members in the group.  Note
	// that the default value will cause all the group members  to be returned
	// in the response. This may take a long time for large groups.
	ReturnMembers bool `json:"return_members"`
}

// NewGroupMembersSetAccessTypeArg returns a new GroupMembersSetAccessTypeArg instance
func NewGroupMembersSetAccessTypeArg(Group *GroupSelector, User *UserSelectorArg, AccessType *GroupAccessType) *GroupMembersSetAccessTypeArg {
	s := new(GroupMembersSetAccessTypeArg)
	s.Group = Group
	s.User = User
	s.AccessType = AccessType
	s.ReturnMembers = true
	return s
}

// GroupSelector : Argument for selecting a single group, either by group_id or
// by external group ID.
type GroupSelector struct {
	dropbox.Tagged
	// GroupId : Group ID.
	GroupId string `json:"group_id,omitempty"`
	// GroupExternalId : External ID of the group.
	GroupExternalId string `json:"group_external_id,omitempty"`
}

// Valid tag values for GroupSelector
const (
	GroupSelectorGroupId         = "group_id"
	GroupSelectorGroupExternalId = "group_external_id"
)

// UnmarshalJSON deserializes into a GroupSelector instance
func (u *GroupSelector) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "group_id":
		if err := json.Unmarshal(body, &u.GroupId); err != nil {
			return err
		}

	case "group_external_id":
		if err := json.Unmarshal(body, &u.GroupExternalId); err != nil {
			return err
		}

	}
	return nil
}

// GroupUpdateError : has no documentation (yet)
type GroupUpdateError struct {
	dropbox.Tagged
}

// Valid tag values for GroupUpdateError
const (
	GroupUpdateErrorExternalIdAlreadyInUse = "external_id_already_in_use"
)

// GroupsGetInfoError : has no documentation (yet)
type GroupsGetInfoError struct {
	dropbox.Tagged
}

// Valid tag values for GroupsGetInfoError
const (
	GroupsGetInfoErrorGroupNotOnTeam = "group_not_on_team"
	GroupsGetInfoErrorOther          = "other"
)

// GroupsGetInfoItem : has no documentation (yet)
type GroupsGetInfoItem struct {
	dropbox.Tagged
	// IdNotFound : An ID that was provided as a parameter to `groupsGetInfo`,
	// and did not match a corresponding group. The ID can be a group ID, or an
	// external ID, depending on how the method was called.
	IdNotFound string `json:"id_not_found,omitempty"`
	// GroupInfo : Info about a group.
	GroupInfo *GroupFullInfo `json:"group_info,omitempty"`
}

// Valid tag values for GroupsGetInfoItem
const (
	GroupsGetInfoItemIdNotFound = "id_not_found"
	GroupsGetInfoItemGroupInfo  = "group_info"
)

// UnmarshalJSON deserializes into a GroupsGetInfoItem instance
func (u *GroupsGetInfoItem) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// GroupInfo : Info about a group.
		GroupInfo json.RawMessage `json:"group_info,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "id_not_found":
		if err := json.Unmarshal(body, &u.IdNotFound); err != nil {
			return err
		}

	case "group_info":
		if err := json.Unmarshal(body, &u.GroupInfo); err != nil {
			return err
		}

	}
	return nil
}

// GroupsListArg : has no documentation (yet)
type GroupsListArg struct {
	// Limit : Number of results to return per call.
	Limit uint32 `json:"limit"`
}

// NewGroupsListArg returns a new GroupsListArg instance
func NewGroupsListArg() *GroupsListArg {
	s := new(GroupsListArg)
	s.Limit = 1000
	return s
}

// GroupsListContinueArg : has no documentation (yet)
type GroupsListContinueArg struct {
	// Cursor : Indicates from what point to get the next set of groups.
	Cursor string `json:"cursor"`
}

// NewGroupsListContinueArg returns a new GroupsListContinueArg instance
func NewGroupsListContinueArg(Cursor string) *GroupsListContinueArg {
	s := new(GroupsListContinueArg)
	s.Cursor = Cursor
	return s
}

// GroupsListContinueError : has no documentation (yet)
type GroupsListContinueError struct {
	dropbox.Tagged
}

// Valid tag values for GroupsListContinueError
const (
	GroupsListContinueErrorInvalidCursor = "invalid_cursor"
	GroupsListContinueErrorOther         = "other"
)

// GroupsListResult : has no documentation (yet)
type GroupsListResult struct {
	// Groups : has no documentation (yet)
	Groups []*team_common.GroupSummary `json:"groups"`
	// Cursor : Pass the cursor into `groupsListContinue` to obtain the
	// additional groups.
	Cursor string `json:"cursor"`
	// HasMore : Is true if there are additional groups that have not been
	// returned yet. An additional call to `groupsListContinue` can retrieve
	// them.
	HasMore bool `json:"has_more"`
}

// NewGroupsListResult returns a new GroupsListResult instance
func NewGroupsListResult(Groups []*team_common.GroupSummary, Cursor string, HasMore bool) *GroupsListResult {
	s := new(GroupsListResult)
	s.Groups = Groups
	s.Cursor = Cursor
	s.HasMore = HasMore
	return s
}

// GroupsMembersListArg : has no documentation (yet)
type GroupsMembersListArg struct {
	// Group : The group whose members are to be listed.
	Group *GroupSelector `json:"group"`
	// Limit : Number of results to return per call.
	Limit uint32 `json:"limit"`
}

// NewGroupsMembersListArg returns a new GroupsMembersListArg instance
func NewGroupsMembersListArg(Group *GroupSelector) *GroupsMembersListArg {
	s := new(GroupsMembersListArg)
	s.Group = Group
	s.Limit = 1000
	return s
}

// GroupsMembersListContinueArg : has no documentation (yet)
type GroupsMembersListContinueArg struct {
	// Cursor : Indicates from what point to get the next set of groups.
	Cursor string `json:"cursor"`
}

// NewGroupsMembersListContinueArg returns a new GroupsMembersListContinueArg instance
func NewGroupsMembersListContinueArg(Cursor string) *GroupsMembersListContinueArg {
	s := new(GroupsMembersListContinueArg)
	s.Cursor = Cursor
	return s
}

// GroupsMembersListContinueError : has no documentation (yet)
type GroupsMembersListContinueError struct {
	dropbox.Tagged
}

// Valid tag values for GroupsMembersListContinueError
const (
	GroupsMembersListContinueErrorInvalidCursor = "invalid_cursor"
	GroupsMembersListContinueErrorOther         = "other"
)

// GroupsMembersListResult : has no documentation (yet)
type GroupsMembersListResult struct {
	// Members : has no documentation (yet)
	Members []*GroupMemberInfo `json:"members"`
	// Cursor : Pass the cursor into `groupsMembersListContinue` to obtain
	// additional group members.
	Cursor string `json:"cursor"`
	// HasMore : Is true if there are additional group members that have not
	// been returned yet. An additional call to `groupsMembersListContinue` can
	// retrieve them.
	HasMore bool `json:"has_more"`
}

// NewGroupsMembersListResult returns a new GroupsMembersListResult instance
func NewGroupsMembersListResult(Members []*GroupMemberInfo, Cursor string, HasMore bool) *GroupsMembersListResult {
	s := new(GroupsMembersListResult)
	s.Members = Members
	s.Cursor = Cursor
	s.HasMore = HasMore
	return s
}

// GroupsPollError : has no documentation (yet)
type GroupsPollError struct {
	dropbox.Tagged
}

// Valid tag values for GroupsPollError
const (
	GroupsPollErrorAccessDenied = "access_denied"
)

// GroupsSelector : Argument for selecting a list of groups, either by
// group_ids, or external group IDs.
type GroupsSelector struct {
	dropbox.Tagged
	// GroupIds : List of group IDs.
	GroupIds []string `json:"group_ids,omitempty"`
	// GroupExternalIds : List of external IDs of groups.
	GroupExternalIds []string `json:"group_external_ids,omitempty"`
}

// Valid tag values for GroupsSelector
const (
	GroupsSelectorGroupIds         = "group_ids"
	GroupsSelectorGroupExternalIds = "group_external_ids"
)

// UnmarshalJSON deserializes into a GroupsSelector instance
func (u *GroupsSelector) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// GroupIds : List of group IDs.
		GroupIds json.RawMessage `json:"group_ids,omitempty"`
		// GroupExternalIds : List of external IDs of groups.
		GroupExternalIds json.RawMessage `json:"group_external_ids,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "group_ids":
		if err := json.Unmarshal(body, &u.GroupIds); err != nil {
			return err
		}

	case "group_external_ids":
		if err := json.Unmarshal(body, &u.GroupExternalIds); err != nil {
			return err
		}

	}
	return nil
}

// ListMemberAppsArg : has no documentation (yet)
type ListMemberAppsArg struct {
	// TeamMemberId : The team member id
	TeamMemberId string `json:"team_member_id"`
}

// NewListMemberAppsArg returns a new ListMemberAppsArg instance
func NewListMemberAppsArg(TeamMemberId string) *ListMemberAppsArg {
	s := new(ListMemberAppsArg)
	s.TeamMemberId = TeamMemberId
	return s
}

// ListMemberAppsError : Error returned by `linkedAppsListMemberLinkedApps`.
type ListMemberAppsError struct {
	dropbox.Tagged
}

// Valid tag values for ListMemberAppsError
const (
	ListMemberAppsErrorMemberNotFound = "member_not_found"
	ListMemberAppsErrorOther          = "other"
)

// ListMemberAppsResult : has no documentation (yet)
type ListMemberAppsResult struct {
	// LinkedApiApps : List of third party applications linked by this team
	// member
	LinkedApiApps []*ApiApp `json:"linked_api_apps"`
}

// NewListMemberAppsResult returns a new ListMemberAppsResult instance
func NewListMemberAppsResult(LinkedApiApps []*ApiApp) *ListMemberAppsResult {
	s := new(ListMemberAppsResult)
	s.LinkedApiApps = LinkedApiApps
	return s
}

// ListMemberDevicesArg : has no documentation (yet)
type ListMemberDevicesArg struct {
	// TeamMemberId : The team's member id
	TeamMemberId string `json:"team_member_id"`
	// IncludeWebSessions : Whether to list web sessions of the team's member
	IncludeWebSessions bool `json:"include_web_sessions"`
	// IncludeDesktopClients : Whether to list linked desktop devices of the
	// team's member
	IncludeDesktopClients bool `json:"include_desktop_clients"`
	// IncludeMobileClients : Whether to list linked mobile devices of the
	// team's member
	IncludeMobileClients bool `json:"include_mobile_clients"`
}

// NewListMemberDevicesArg returns a new ListMemberDevicesArg instance
func NewListMemberDevicesArg(TeamMemberId string) *ListMemberDevicesArg {
	s := new(ListMemberDevicesArg)
	s.TeamMemberId = TeamMemberId
	s.IncludeWebSessions = true
	s.IncludeDesktopClients = true
	s.IncludeMobileClients = true
	return s
}

// ListMemberDevicesError : has no documentation (yet)
type ListMemberDevicesError struct {
	dropbox.Tagged
}

// Valid tag values for ListMemberDevicesError
const (
	ListMemberDevicesErrorMemberNotFound = "member_not_found"
	ListMemberDevicesErrorOther          = "other"
)

// ListMemberDevicesResult : has no documentation (yet)
type ListMemberDevicesResult struct {
	// ActiveWebSessions : List of web sessions made by this team member
	ActiveWebSessions []*ActiveWebSession `json:"active_web_sessions,omitempty"`
	// DesktopClientSessions : List of desktop clients used by this team member
	DesktopClientSessions []*DesktopClientSession `json:"desktop_client_sessions,omitempty"`
	// MobileClientSessions : List of mobile client used by this team member
	MobileClientSessions []*MobileClientSession `json:"mobile_client_sessions,omitempty"`
}

// NewListMemberDevicesResult returns a new ListMemberDevicesResult instance
func NewListMemberDevicesResult() *ListMemberDevicesResult {
	s := new(ListMemberDevicesResult)
	return s
}

// ListMembersAppsArg : Arguments for `linkedAppsListMembersLinkedApps`.
type ListMembersAppsArg struct {
	// Cursor : At the first call to the `linkedAppsListMembersLinkedApps` the
	// cursor shouldn't be passed. Then, if the result of the call includes a
	// cursor, the following requests should include the received cursors in
	// order to receive the next sub list of the team applications
	Cursor string `json:"cursor,omitempty"`
}

// NewListMembersAppsArg returns a new ListMembersAppsArg instance
func NewListMembersAppsArg() *ListMembersAppsArg {
	s := new(ListMembersAppsArg)
	return s
}

// ListMembersAppsError : Error returned by `linkedAppsListMembersLinkedApps`
type ListMembersAppsError struct {
	dropbox.Tagged
}

// Valid tag values for ListMembersAppsError
const (
	ListMembersAppsErrorReset = "reset"
	ListMembersAppsErrorOther = "other"
)

// ListMembersAppsResult : Information returned by
// `linkedAppsListMembersLinkedApps`.
type ListMembersAppsResult struct {
	// Apps : The linked applications of each member of the team
	Apps []*MemberLinkedApps `json:"apps"`
	// HasMore : If true, then there are more apps available. Pass the cursor to
	// `linkedAppsListMembersLinkedApps` to retrieve the rest.
	HasMore bool `json:"has_more"`
	// Cursor : Pass the cursor into `linkedAppsListMembersLinkedApps` to
	// receive the next sub list of team's applications.
	Cursor string `json:"cursor,omitempty"`
}

// NewListMembersAppsResult returns a new ListMembersAppsResult instance
func NewListMembersAppsResult(Apps []*MemberLinkedApps, HasMore bool) *ListMembersAppsResult {
	s := new(ListMembersAppsResult)
	s.Apps = Apps
	s.HasMore = HasMore
	return s
}

// ListMembersDevicesArg : has no documentation (yet)
type ListMembersDevicesArg struct {
	// Cursor : At the first call to the `devicesListMembersDevices` the cursor
	// shouldn't be passed. Then, if the result of the call includes a cursor,
	// the following requests should include the received cursors in order to
	// receive the next sub list of team devices
	Cursor string `json:"cursor,omitempty"`
	// IncludeWebSessions : Whether to list web sessions of the team members
	IncludeWebSessions bool `json:"include_web_sessions"`
	// IncludeDesktopClients : Whether to list desktop clients of the team
	// members
	IncludeDesktopClients bool `json:"include_desktop_clients"`
	// IncludeMobileClients : Whether to list mobile clients of the team members
	IncludeMobileClients bool `json:"include_mobile_clients"`
}

// NewListMembersDevicesArg returns a new ListMembersDevicesArg instance
func NewListMembersDevicesArg() *ListMembersDevicesArg {
	s := new(ListMembersDevicesArg)
	s.IncludeWebSessions = true
	s.IncludeDesktopClients = true
	s.IncludeMobileClients = true
	return s
}

// ListMembersDevicesError : has no documentation (yet)
type ListMembersDevicesError struct {
	dropbox.Tagged
}

// Valid tag values for ListMembersDevicesError
const (
	ListMembersDevicesErrorReset = "reset"
	ListMembersDevicesErrorOther = "other"
)

// ListMembersDevicesResult : has no documentation (yet)
type ListMembersDevicesResult struct {
	// Devices : The devices of each member of the team
	Devices []*MemberDevices `json:"devices"`
	// HasMore : If true, then there are more devices available. Pass the cursor
	// to `devicesListMembersDevices` to retrieve the rest.
	HasMore bool `json:"has_more"`
	// Cursor : Pass the cursor into `devicesListMembersDevices` to receive the
	// next sub list of team's devices.
	Cursor string `json:"cursor,omitempty"`
}

// NewListMembersDevicesResult returns a new ListMembersDevicesResult instance
func NewListMembersDevicesResult(Devices []*MemberDevices, HasMore bool) *ListMembersDevicesResult {
	s := new(ListMembersDevicesResult)
	s.Devices = Devices
	s.HasMore = HasMore
	return s
}

// ListTeamAppsArg : Arguments for `linkedAppsListTeamLinkedApps`.
type ListTeamAppsArg struct {
	// Cursor : At the first call to the `linkedAppsListTeamLinkedApps` the
	// cursor shouldn't be passed. Then, if the result of the call includes a
	// cursor, the following requests should include the received cursors in
	// order to receive the next sub list of the team applications
	Cursor string `json:"cursor,omitempty"`
}

// NewListTeamAppsArg returns a new ListTeamAppsArg instance
func NewListTeamAppsArg() *ListTeamAppsArg {
	s := new(ListTeamAppsArg)
	return s
}

// ListTeamAppsError : Error returned by `linkedAppsListTeamLinkedApps`
type ListTeamAppsError struct {
	dropbox.Tagged
}

// Valid tag values for ListTeamAppsError
const (
	ListTeamAppsErrorReset = "reset"
	ListTeamAppsErrorOther = "other"
)

// ListTeamAppsResult : Information returned by `linkedAppsListTeamLinkedApps`.
type ListTeamAppsResult struct {
	// Apps : The linked applications of each member of the team
	Apps []*MemberLinkedApps `json:"apps"`
	// HasMore : If true, then there are more apps available. Pass the cursor to
	// `linkedAppsListTeamLinkedApps` to retrieve the rest.
	HasMore bool `json:"has_more"`
	// Cursor : Pass the cursor into `linkedAppsListTeamLinkedApps` to receive
	// the next sub list of team's applications.
	Cursor string `json:"cursor,omitempty"`
}

// NewListTeamAppsResult returns a new ListTeamAppsResult instance
func NewListTeamAppsResult(Apps []*MemberLinkedApps, HasMore bool) *ListTeamAppsResult {
	s := new(ListTeamAppsResult)
	s.Apps = Apps
	s.HasMore = HasMore
	return s
}

// ListTeamDevicesArg : has no documentation (yet)
type ListTeamDevicesArg struct {
	// Cursor : At the first call to the `devicesListTeamDevices` the cursor
	// shouldn't be passed. Then, if the result of the call includes a cursor,
	// the following requests should include the received cursors in order to
	// receive the next sub list of team devices
	Cursor string `json:"cursor,omitempty"`
	// IncludeWebSessions : Whether to list web sessions of the team members
	IncludeWebSessions bool `json:"include_web_sessions"`
	// IncludeDesktopClients : Whether to list desktop clients of the team
	// members
	IncludeDesktopClients bool `json:"include_desktop_clients"`
	// IncludeMobileClients : Whether to list mobile clients of the team members
	IncludeMobileClients bool `json:"include_mobile_clients"`
}

// NewListTeamDevicesArg returns a new ListTeamDevicesArg instance
func NewListTeamDevicesArg() *ListTeamDevicesArg {
	s := new(ListTeamDevicesArg)
	s.IncludeWebSessions = true
	s.IncludeDesktopClients = true
	s.IncludeMobileClients = true
	return s
}

// ListTeamDevicesError : has no documentation (yet)
type ListTeamDevicesError struct {
	dropbox.Tagged
}

// Valid tag values for ListTeamDevicesError
const (
	ListTeamDevicesErrorReset = "reset"
	ListTeamDevicesErrorOther = "other"
)

// ListTeamDevicesResult : has no documentation (yet)
type ListTeamDevicesResult struct {
	// Devices : The devices of each member of the team
	Devices []*MemberDevices `json:"devices"`
	// HasMore : If true, then there are more devices available. Pass the cursor
	// to `devicesListTeamDevices` to retrieve the rest.
	HasMore bool `json:"has_more"`
	// Cursor : Pass the cursor into `devicesListTeamDevices` to receive the
	// next sub list of team's devices.
	Cursor string `json:"cursor,omitempty"`
}

// NewListTeamDevicesResult returns a new ListTeamDevicesResult instance
func NewListTeamDevicesResult(Devices []*MemberDevices, HasMore bool) *ListTeamDevicesResult {
	s := new(ListTeamDevicesResult)
	s.Devices = Devices
	s.HasMore = HasMore
	return s
}

// MemberAccess : Specify access type a member should have when joined to a
// group.
type MemberAccess struct {
	// User : Identity of a user.
	User *UserSelectorArg `json:"user"`
	// AccessType : Access type.
	AccessType *GroupAccessType `json:"access_type"`
}

// NewMemberAccess returns a new MemberAccess instance
func NewMemberAccess(User *UserSelectorArg, AccessType *GroupAccessType) *MemberAccess {
	s := new(MemberAccess)
	s.User = User
	s.AccessType = AccessType
	return s
}

// MemberAddArg : has no documentation (yet)
type MemberAddArg struct {
	// MemberEmail : has no documentation (yet)
	MemberEmail string `json:"member_email"`
	// MemberGivenName : Member's first name.
	MemberGivenName string `json:"member_given_name"`
	// MemberSurname : Member's last name.
	MemberSurname string `json:"member_surname"`
	// MemberExternalId : External ID for member.
	MemberExternalId string `json:"member_external_id,omitempty"`
	// SendWelcomeEmail : Whether to send a welcome email to the member. If
	// send_welcome_email is false, no email invitation will be sent to the
	// user. This may be useful for apps using single sign-on (SSO) flows for
	// onboarding that want to handle announcements themselves.
	SendWelcomeEmail bool `json:"send_welcome_email"`
	// Role : has no documentation (yet)
	Role *AdminTier `json:"role"`
}

// NewMemberAddArg returns a new MemberAddArg instance
func NewMemberAddArg(MemberEmail string, MemberGivenName string, MemberSurname string) *MemberAddArg {
	s := new(MemberAddArg)
	s.MemberEmail = MemberEmail
	s.MemberGivenName = MemberGivenName
	s.MemberSurname = MemberSurname
	s.SendWelcomeEmail = true
	s.Role = &AdminTier{Tagged: dropbox.Tagged{"member_only"}}
	return s
}

// MemberAddResult : Describes the result of attempting to add a single user to
// the team. 'success' is the only value indicating that a user was indeed added
// to the team - the other values explain the type of failure that occurred, and
// include the email of the user for which the operation has failed.
type MemberAddResult struct {
	dropbox.Tagged
	// Success : Describes a user that was successfully added to the team.
	Success *TeamMemberInfo `json:"success,omitempty"`
	// TeamLicenseLimit : Team is already full. The organization has no
	// available licenses.
	TeamLicenseLimit string `json:"team_license_limit,omitempty"`
	// FreeTeamMemberLimitReached : Team is already full. The free team member
	// limit has been reached.
	FreeTeamMemberLimitReached string `json:"free_team_member_limit_reached,omitempty"`
	// UserAlreadyOnTeam : User is already on this team. The provided email
	// address is associated with a user who is already a member of or invited
	// to the team.
	UserAlreadyOnTeam string `json:"user_already_on_team,omitempty"`
	// UserOnAnotherTeam : User is already on another team. The provided email
	// address is associated with a user that is already a member or invited to
	// another team.
	UserOnAnotherTeam string `json:"user_on_another_team,omitempty"`
	// UserAlreadyPaired : User is already paired.
	UserAlreadyPaired string `json:"user_already_paired,omitempty"`
	// UserMigrationFailed : User migration has failed.
	UserMigrationFailed string `json:"user_migration_failed,omitempty"`
	// DuplicateExternalMemberId : A user with the given external member ID
	// already exists on the team.
	DuplicateExternalMemberId string `json:"duplicate_external_member_id,omitempty"`
	// UserCreationFailed : User creation has failed.
	UserCreationFailed string `json:"user_creation_failed,omitempty"`
}

// Valid tag values for MemberAddResult
const (
	MemberAddResultSuccess                    = "success"
	MemberAddResultTeamLicenseLimit           = "team_license_limit"
	MemberAddResultFreeTeamMemberLimitReached = "free_team_member_limit_reached"
	MemberAddResultUserAlreadyOnTeam          = "user_already_on_team"
	MemberAddResultUserOnAnotherTeam          = "user_on_another_team"
	MemberAddResultUserAlreadyPaired          = "user_already_paired"
	MemberAddResultUserMigrationFailed        = "user_migration_failed"
	MemberAddResultDuplicateExternalMemberId  = "duplicate_external_member_id"
	MemberAddResultUserCreationFailed         = "user_creation_failed"
)

// UnmarshalJSON deserializes into a MemberAddResult instance
func (u *MemberAddResult) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// Success : Describes a user that was successfully added to the team.
		Success json.RawMessage `json:"success,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "success":
		if err := json.Unmarshal(body, &u.Success); err != nil {
			return err
		}

	case "team_license_limit":
		if err := json.Unmarshal(body, &u.TeamLicenseLimit); err != nil {
			return err
		}

	case "free_team_member_limit_reached":
		if err := json.Unmarshal(body, &u.FreeTeamMemberLimitReached); err != nil {
			return err
		}

	case "user_already_on_team":
		if err := json.Unmarshal(body, &u.UserAlreadyOnTeam); err != nil {
			return err
		}

	case "user_on_another_team":
		if err := json.Unmarshal(body, &u.UserOnAnotherTeam); err != nil {
			return err
		}

	case "user_already_paired":
		if err := json.Unmarshal(body, &u.UserAlreadyPaired); err != nil {
			return err
		}

	case "user_migration_failed":
		if err := json.Unmarshal(body, &u.UserMigrationFailed); err != nil {
			return err
		}

	case "duplicate_external_member_id":
		if err := json.Unmarshal(body, &u.DuplicateExternalMemberId); err != nil {
			return err
		}

	case "user_creation_failed":
		if err := json.Unmarshal(body, &u.UserCreationFailed); err != nil {
			return err
		}

	}
	return nil
}

// MemberDevices : Information on devices of a team's member.
type MemberDevices struct {
	// TeamMemberId : The member unique Id
	TeamMemberId string `json:"team_member_id"`
	// WebSessions : List of web sessions made by this team member
	WebSessions []*ActiveWebSession `json:"web_sessions,omitempty"`
	// DesktopClients : List of desktop clients by this team member
	DesktopClients []*DesktopClientSession `json:"desktop_clients,omitempty"`
	// MobileClients : List of mobile clients by this team member
	MobileClients []*MobileClientSession `json:"mobile_clients,omitempty"`
}

// NewMemberDevices returns a new MemberDevices instance
func NewMemberDevices(TeamMemberId string) *MemberDevices {
	s := new(MemberDevices)
	s.TeamMemberId = TeamMemberId
	return s
}

// MemberLinkedApps : Information on linked applications of a team member.
type MemberLinkedApps struct {
	// TeamMemberId : The member unique Id
	TeamMemberId string `json:"team_member_id"`
	// LinkedApiApps : List of third party applications linked by this team
	// member
	LinkedApiApps []*ApiApp `json:"linked_api_apps"`
}

// NewMemberLinkedApps returns a new MemberLinkedApps instance
func NewMemberLinkedApps(TeamMemberId string, LinkedApiApps []*ApiApp) *MemberLinkedApps {
	s := new(MemberLinkedApps)
	s.TeamMemberId = TeamMemberId
	s.LinkedApiApps = LinkedApiApps
	return s
}

// MemberProfile : Basic member profile.
type MemberProfile struct {
	// TeamMemberId : ID of user as a member of a team.
	TeamMemberId string `json:"team_member_id"`
	// ExternalId : External ID that a team can attach to the user. An
	// application using the API may find it easier to use their own IDs instead
	// of Dropbox IDs like account_id or team_member_id.
	ExternalId string `json:"external_id,omitempty"`
	// AccountId : A user's account identifier.
	AccountId string `json:"account_id,omitempty"`
	// Email : Email address of user.
	Email string `json:"email"`
	// EmailVerified : Is true if the user's email is verified to be owned by
	// the user.
	EmailVerified bool `json:"email_verified"`
	// Status : The user's status as a member of a specific team.
	Status *TeamMemberStatus `json:"status"`
	// Name : Representations for a person's name.
	Name *users.Name `json:"name"`
	// MembershipType : The user's membership type: full (normal team member) vs
	// limited (does not use a license; no access to the team's shared quota).
	MembershipType *TeamMembershipType `json:"membership_type"`
}

// NewMemberProfile returns a new MemberProfile instance
func NewMemberProfile(TeamMemberId string, Email string, EmailVerified bool, Status *TeamMemberStatus, Name *users.Name, MembershipType *TeamMembershipType) *MemberProfile {
	s := new(MemberProfile)
	s.TeamMemberId = TeamMemberId
	s.Email = Email
	s.EmailVerified = EmailVerified
	s.Status = Status
	s.Name = Name
	s.MembershipType = MembershipType
	return s
}

// UserSelectorError : Error that can be returned whenever a struct derived from
// `UserSelectorArg` is used.
type UserSelectorError struct {
	dropbox.Tagged
}

// Valid tag values for UserSelectorError
const (
	UserSelectorErrorUserNotFound = "user_not_found"
)

// MemberSelectorError : has no documentation (yet)
type MemberSelectorError struct {
	dropbox.Tagged
}

// Valid tag values for MemberSelectorError
const (
	MemberSelectorErrorUserNotInTeam = "user_not_in_team"
)

// MembersAddArg : has no documentation (yet)
type MembersAddArg struct {
	// NewMembers : Details of new members to be added to the team.
	NewMembers []*MemberAddArg `json:"new_members"`
	// ForceAsync : Whether to force the add to happen asynchronously.
	ForceAsync bool `json:"force_async"`
}

// NewMembersAddArg returns a new MembersAddArg instance
func NewMembersAddArg(NewMembers []*MemberAddArg) *MembersAddArg {
	s := new(MembersAddArg)
	s.NewMembers = NewMembers
	s.ForceAsync = false
	return s
}

// MembersAddJobStatus : has no documentation (yet)
type MembersAddJobStatus struct {
	dropbox.Tagged
	// Complete : The asynchronous job has finished. For each member that was
	// specified in the parameter `MembersAddArg` that was provided to
	// `membersAdd`, a corresponding item is returned in this list.
	Complete []*MemberAddResult `json:"complete,omitempty"`
	// Failed : The asynchronous job returned an error. The string contains an
	// error message.
	Failed string `json:"failed,omitempty"`
}

// Valid tag values for MembersAddJobStatus
const (
	MembersAddJobStatusComplete = "complete"
	MembersAddJobStatusFailed   = "failed"
)

// UnmarshalJSON deserializes into a MembersAddJobStatus instance
func (u *MembersAddJobStatus) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// Complete : The asynchronous job has finished. For each member that
		// was specified in the parameter `MembersAddArg` that was provided to
		// `membersAdd`, a corresponding item is returned in this list.
		Complete json.RawMessage `json:"complete,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "complete":
		if err := json.Unmarshal(body, &u.Complete); err != nil {
			return err
		}

	case "failed":
		if err := json.Unmarshal(body, &u.Failed); err != nil {
			return err
		}

	}
	return nil
}

// MembersAddLaunch : has no documentation (yet)
type MembersAddLaunch struct {
	dropbox.Tagged
	// Complete : has no documentation (yet)
	Complete []*MemberAddResult `json:"complete,omitempty"`
}

// Valid tag values for MembersAddLaunch
const (
	MembersAddLaunchComplete = "complete"
)

// UnmarshalJSON deserializes into a MembersAddLaunch instance
func (u *MembersAddLaunch) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// Complete : has no documentation (yet)
		Complete json.RawMessage `json:"complete,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "complete":
		if err := json.Unmarshal(body, &u.Complete); err != nil {
			return err
		}

	}
	return nil
}

// MembersDeactivateArg : Exactly one of team_member_id, email, or external_id
// must be provided to identify the user account.
type MembersDeactivateArg struct {
	// User : Identity of user to remove/suspend.
	User *UserSelectorArg `json:"user"`
	// WipeData : If provided, controls if the user's data will be deleted on
	// their linked devices.
	WipeData bool `json:"wipe_data"`
}

// NewMembersDeactivateArg returns a new MembersDeactivateArg instance
func NewMembersDeactivateArg(User *UserSelectorArg) *MembersDeactivateArg {
	s := new(MembersDeactivateArg)
	s.User = User
	s.WipeData = true
	return s
}

// MembersDeactivateError : has no documentation (yet)
type MembersDeactivateError struct {
	dropbox.Tagged
}

// Valid tag values for MembersDeactivateError
const (
	MembersDeactivateErrorUserNotInTeam = "user_not_in_team"
	MembersDeactivateErrorOther         = "other"
)

// MembersGetInfoArgs : has no documentation (yet)
type MembersGetInfoArgs struct {
	// Members : List of team members.
	Members []*UserSelectorArg `json:"members"`
}

// NewMembersGetInfoArgs returns a new MembersGetInfoArgs instance
func NewMembersGetInfoArgs(Members []*UserSelectorArg) *MembersGetInfoArgs {
	s := new(MembersGetInfoArgs)
	s.Members = Members
	return s
}

// MembersGetInfoError : has no documentation (yet)
type MembersGetInfoError struct {
	dropbox.Tagged
}

// Valid tag values for MembersGetInfoError
const (
	MembersGetInfoErrorOther = "other"
)

// MembersGetInfoItem : Describes a result obtained for a single user whose id
// was specified in the parameter of `membersGetInfo`.
type MembersGetInfoItem struct {
	dropbox.Tagged
	// IdNotFound : An ID that was provided as a parameter to `membersGetInfo`,
	// and did not match a corresponding user. This might be a team_member_id,
	// an email, or an external ID, depending on how the method was called.
	IdNotFound string `json:"id_not_found,omitempty"`
	// MemberInfo : Info about a team member.
	MemberInfo *TeamMemberInfo `json:"member_info,omitempty"`
}

// Valid tag values for MembersGetInfoItem
const (
	MembersGetInfoItemIdNotFound = "id_not_found"
	MembersGetInfoItemMemberInfo = "member_info"
)

// UnmarshalJSON deserializes into a MembersGetInfoItem instance
func (u *MembersGetInfoItem) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// MemberInfo : Info about a team member.
		MemberInfo json.RawMessage `json:"member_info,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "id_not_found":
		if err := json.Unmarshal(body, &u.IdNotFound); err != nil {
			return err
		}

	case "member_info":
		if err := json.Unmarshal(body, &u.MemberInfo); err != nil {
			return err
		}

	}
	return nil
}

// MembersListArg : has no documentation (yet)
type MembersListArg struct {
	// Limit : Number of results to return per call.
	Limit uint32 `json:"limit"`
}

// NewMembersListArg returns a new MembersListArg instance
func NewMembersListArg() *MembersListArg {
	s := new(MembersListArg)
	s.Limit = 1000
	return s
}

// MembersListContinueArg : has no documentation (yet)
type MembersListContinueArg struct {
	// Cursor : Indicates from what point to get the next set of members.
	Cursor string `json:"cursor"`
}

// NewMembersListContinueArg returns a new MembersListContinueArg instance
func NewMembersListContinueArg(Cursor string) *MembersListContinueArg {
	s := new(MembersListContinueArg)
	s.Cursor = Cursor
	return s
}

// MembersListContinueError : has no documentation (yet)
type MembersListContinueError struct {
	dropbox.Tagged
}

// Valid tag values for MembersListContinueError
const (
	MembersListContinueErrorInvalidCursor = "invalid_cursor"
	MembersListContinueErrorOther         = "other"
)

// MembersListError : has no documentation (yet)
type MembersListError struct {
	dropbox.Tagged
}

// Valid tag values for MembersListError
const (
	MembersListErrorOther = "other"
)

// MembersListResult : has no documentation (yet)
type MembersListResult struct {
	// Members : List of team members.
	Members []*TeamMemberInfo `json:"members"`
	// Cursor : Pass the cursor into `membersListContinue` to obtain the
	// additional members.
	Cursor string `json:"cursor"`
	// HasMore : Is true if there are additional team members that have not been
	// returned yet. An additional call to `membersListContinue` can retrieve
	// them.
	HasMore bool `json:"has_more"`
}

// NewMembersListResult returns a new MembersListResult instance
func NewMembersListResult(Members []*TeamMemberInfo, Cursor string, HasMore bool) *MembersListResult {
	s := new(MembersListResult)
	s.Members = Members
	s.Cursor = Cursor
	s.HasMore = HasMore
	return s
}

// MembersRemoveArg : has no documentation (yet)
type MembersRemoveArg struct {
	MembersDeactivateArg
	// TransferDestId : If provided, files from the deleted member account will
	// be transferred to this user.
	TransferDestId *UserSelectorArg `json:"transfer_dest_id,omitempty"`
	// TransferAdminId : If provided, errors during the transfer process will be
	// sent via email to this user. If the transfer_dest_id argument was
	// provided, then this argument must be provided as well.
	TransferAdminId *UserSelectorArg `json:"transfer_admin_id,omitempty"`
	// KeepAccount : Downgrade the member to a Basic account. The user will
	// retain the email address associated with their Dropbox  account and data
	// in their account that is not restricted to team members.
	KeepAccount bool `json:"keep_account"`
}

// NewMembersRemoveArg returns a new MembersRemoveArg instance
func NewMembersRemoveArg(User *UserSelectorArg) *MembersRemoveArg {
	s := new(MembersRemoveArg)
	s.User = User
	s.WipeData = true
	s.KeepAccount = false
	return s
}

// MembersRemoveError : has no documentation (yet)
type MembersRemoveError struct {
	dropbox.Tagged
}

// Valid tag values for MembersRemoveError
const (
	MembersRemoveErrorRemoveLastAdmin                     = "remove_last_admin"
	MembersRemoveErrorRemovedAndTransferDestShouldDiffer  = "removed_and_transfer_dest_should_differ"
	MembersRemoveErrorRemovedAndTransferAdminShouldDiffer = "removed_and_transfer_admin_should_differ"
	MembersRemoveErrorTransferDestUserNotFound            = "transfer_dest_user_not_found"
	MembersRemoveErrorTransferDestUserNotInTeam           = "transfer_dest_user_not_in_team"
	MembersRemoveErrorTransferAdminUserNotFound           = "transfer_admin_user_not_found"
	MembersRemoveErrorTransferAdminUserNotInTeam          = "transfer_admin_user_not_in_team"
	MembersRemoveErrorUnspecifiedTransferAdminId          = "unspecified_transfer_admin_id"
	MembersRemoveErrorTransferAdminIsNotAdmin             = "transfer_admin_is_not_admin"
	MembersRemoveErrorCannotKeepAccountAndTransfer        = "cannot_keep_account_and_transfer"
	MembersRemoveErrorCannotKeepAccountAndDeleteData      = "cannot_keep_account_and_delete_data"
)

// MembersSendWelcomeError : has no documentation (yet)
type MembersSendWelcomeError struct {
	dropbox.Tagged
}

// Valid tag values for MembersSendWelcomeError
const (
	MembersSendWelcomeErrorOther = "other"
)

// MembersSetPermissionsArg : Exactly one of team_member_id, email, or
// external_id must be provided to identify the user account.
type MembersSetPermissionsArg struct {
	// User : Identity of user whose role will be set.
	User *UserSelectorArg `json:"user"`
	// NewRole : The new role of the member.
	NewRole *AdminTier `json:"new_role"`
}

// NewMembersSetPermissionsArg returns a new MembersSetPermissionsArg instance
func NewMembersSetPermissionsArg(User *UserSelectorArg, NewRole *AdminTier) *MembersSetPermissionsArg {
	s := new(MembersSetPermissionsArg)
	s.User = User
	s.NewRole = NewRole
	return s
}

// MembersSetPermissionsError : has no documentation (yet)
type MembersSetPermissionsError struct {
	dropbox.Tagged
}

// Valid tag values for MembersSetPermissionsError
const (
	MembersSetPermissionsErrorLastAdmin            = "last_admin"
	MembersSetPermissionsErrorUserNotInTeam        = "user_not_in_team"
	MembersSetPermissionsErrorCannotSetPermissions = "cannot_set_permissions"
	MembersSetPermissionsErrorTeamLicenseLimit     = "team_license_limit"
	MembersSetPermissionsErrorOther                = "other"
)

// MembersSetPermissionsResult : has no documentation (yet)
type MembersSetPermissionsResult struct {
	// TeamMemberId : The member ID of the user to which the change was applied.
	TeamMemberId string `json:"team_member_id"`
	// Role : The role after the change.
	Role *AdminTier `json:"role"`
}

// NewMembersSetPermissionsResult returns a new MembersSetPermissionsResult instance
func NewMembersSetPermissionsResult(TeamMemberId string, Role *AdminTier) *MembersSetPermissionsResult {
	s := new(MembersSetPermissionsResult)
	s.TeamMemberId = TeamMemberId
	s.Role = Role
	return s
}

// MembersSetProfileArg : Exactly one of team_member_id, email, or external_id
// must be provided to identify the user account. At least one of new_email,
// new_external_id, new_given_name, and/or new_surname must be provided.
type MembersSetProfileArg struct {
	// User : Identity of user whose profile will be set.
	User *UserSelectorArg `json:"user"`
	// NewEmail : New email for member.
	NewEmail string `json:"new_email,omitempty"`
	// NewExternalId : New external ID for member.
	NewExternalId string `json:"new_external_id,omitempty"`
	// NewGivenName : New given name for member.
	NewGivenName string `json:"new_given_name,omitempty"`
	// NewSurname : New surname for member.
	NewSurname string `json:"new_surname,omitempty"`
}

// NewMembersSetProfileArg returns a new MembersSetProfileArg instance
func NewMembersSetProfileArg(User *UserSelectorArg) *MembersSetProfileArg {
	s := new(MembersSetProfileArg)
	s.User = User
	return s
}

// MembersSetProfileError : has no documentation (yet)
type MembersSetProfileError struct {
	dropbox.Tagged
}

// Valid tag values for MembersSetProfileError
const (
	MembersSetProfileErrorExternalIdAndNewExternalIdUnsafe = "external_id_and_new_external_id_unsafe"
	MembersSetProfileErrorNoNewDataSpecified               = "no_new_data_specified"
	MembersSetProfileErrorEmailReservedForOtherUser        = "email_reserved_for_other_user"
	MembersSetProfileErrorExternalIdUsedByOtherUser        = "external_id_used_by_other_user"
	MembersSetProfileErrorSetProfileDisallowed             = "set_profile_disallowed"
	MembersSetProfileErrorParamCannotBeEmpty               = "param_cannot_be_empty"
	MembersSetProfileErrorOther                            = "other"
)

// MembersSuspendError : has no documentation (yet)
type MembersSuspendError struct {
	dropbox.Tagged
}

// Valid tag values for MembersSuspendError
const (
	MembersSuspendErrorSuspendInactiveUser = "suspend_inactive_user"
	MembersSuspendErrorSuspendLastAdmin    = "suspend_last_admin"
	MembersSuspendErrorTeamLicenseLimit    = "team_license_limit"
)

// MembersUnsuspendArg : Exactly one of team_member_id, email, or external_id
// must be provided to identify the user account.
type MembersUnsuspendArg struct {
	// User : Identity of user to unsuspend.
	User *UserSelectorArg `json:"user"`
}

// NewMembersUnsuspendArg returns a new MembersUnsuspendArg instance
func NewMembersUnsuspendArg(User *UserSelectorArg) *MembersUnsuspendArg {
	s := new(MembersUnsuspendArg)
	s.User = User
	return s
}

// MembersUnsuspendError : has no documentation (yet)
type MembersUnsuspendError struct {
	dropbox.Tagged
}

// Valid tag values for MembersUnsuspendError
const (
	MembersUnsuspendErrorUnsuspendNonSuspendedMember = "unsuspend_non_suspended_member"
	MembersUnsuspendErrorTeamLicenseLimit            = "team_license_limit"
)

// MobileClientPlatform : has no documentation (yet)
type MobileClientPlatform struct {
	dropbox.Tagged
}

// Valid tag values for MobileClientPlatform
const (
	MobileClientPlatformIphone       = "iphone"
	MobileClientPlatformIpad         = "ipad"
	MobileClientPlatformAndroid      = "android"
	MobileClientPlatformWindowsPhone = "windows_phone"
	MobileClientPlatformBlackberry   = "blackberry"
	MobileClientPlatformOther        = "other"
)

// MobileClientSession : Information about linked Dropbox mobile client sessions
type MobileClientSession struct {
	DeviceSession
	// DeviceName : The device name
	DeviceName string `json:"device_name"`
	// ClientType : The mobile application type
	ClientType *MobileClientPlatform `json:"client_type"`
	// ClientVersion : The dropbox client version
	ClientVersion string `json:"client_version,omitempty"`
	// OsVersion : The hosting OS version
	OsVersion string `json:"os_version,omitempty"`
	// LastCarrier : last carrier used by the device
	LastCarrier string `json:"last_carrier,omitempty"`
}

// NewMobileClientSession returns a new MobileClientSession instance
func NewMobileClientSession(SessionId string, DeviceName string, ClientType *MobileClientPlatform) *MobileClientSession {
	s := new(MobileClientSession)
	s.SessionId = SessionId
	s.DeviceName = DeviceName
	s.ClientType = ClientType
	return s
}

// RevokeDesktopClientArg : has no documentation (yet)
type RevokeDesktopClientArg struct {
	DeviceSessionArg
	// DeleteOnUnlink : Whether to delete all files of the account (this is
	// possible only if supported by the desktop client and  will be made the
	// next time the client access the account)
	DeleteOnUnlink bool `json:"delete_on_unlink"`
}

// NewRevokeDesktopClientArg returns a new RevokeDesktopClientArg instance
func NewRevokeDesktopClientArg(SessionId string, TeamMemberId string) *RevokeDesktopClientArg {
	s := new(RevokeDesktopClientArg)
	s.SessionId = SessionId
	s.TeamMemberId = TeamMemberId
	s.DeleteOnUnlink = false
	return s
}

// RevokeDeviceSessionArg : has no documentation (yet)
type RevokeDeviceSessionArg struct {
	dropbox.Tagged
	// WebSession : End an active session
	WebSession *DeviceSessionArg `json:"web_session,omitempty"`
	// DesktopClient : Unlink a linked desktop device
	DesktopClient *RevokeDesktopClientArg `json:"desktop_client,omitempty"`
	// MobileClient : Unlink a linked mobile device
	MobileClient *DeviceSessionArg `json:"mobile_client,omitempty"`
}

// Valid tag values for RevokeDeviceSessionArg
const (
	RevokeDeviceSessionArgWebSession    = "web_session"
	RevokeDeviceSessionArgDesktopClient = "desktop_client"
	RevokeDeviceSessionArgMobileClient  = "mobile_client"
)

// UnmarshalJSON deserializes into a RevokeDeviceSessionArg instance
func (u *RevokeDeviceSessionArg) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// WebSession : End an active session
		WebSession json.RawMessage `json:"web_session,omitempty"`
		// DesktopClient : Unlink a linked desktop device
		DesktopClient json.RawMessage `json:"desktop_client,omitempty"`
		// MobileClient : Unlink a linked mobile device
		MobileClient json.RawMessage `json:"mobile_client,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "web_session":
		if err := json.Unmarshal(body, &u.WebSession); err != nil {
			return err
		}

	case "desktop_client":
		if err := json.Unmarshal(body, &u.DesktopClient); err != nil {
			return err
		}

	case "mobile_client":
		if err := json.Unmarshal(body, &u.MobileClient); err != nil {
			return err
		}

	}
	return nil
}

// RevokeDeviceSessionBatchArg : has no documentation (yet)
type RevokeDeviceSessionBatchArg struct {
	// RevokeDevices : has no documentation (yet)
	RevokeDevices []*RevokeDeviceSessionArg `json:"revoke_devices"`
}

// NewRevokeDeviceSessionBatchArg returns a new RevokeDeviceSessionBatchArg instance
func NewRevokeDeviceSessionBatchArg(RevokeDevices []*RevokeDeviceSessionArg) *RevokeDeviceSessionBatchArg {
	s := new(RevokeDeviceSessionBatchArg)
	s.RevokeDevices = RevokeDevices
	return s
}

// RevokeDeviceSessionBatchError : has no documentation (yet)
type RevokeDeviceSessionBatchError struct {
	dropbox.Tagged
}

// Valid tag values for RevokeDeviceSessionBatchError
const (
	RevokeDeviceSessionBatchErrorUnspecified = "unspecified"
)

// RevokeDeviceSessionBatchResult : has no documentation (yet)
type RevokeDeviceSessionBatchResult struct {
	// RevokeDevicesStatus : has no documentation (yet)
	RevokeDevicesStatus []*RevokeDeviceSessionStatus `json:"revoke_devices_status"`
}

// NewRevokeDeviceSessionBatchResult returns a new RevokeDeviceSessionBatchResult instance
func NewRevokeDeviceSessionBatchResult(RevokeDevicesStatus []*RevokeDeviceSessionStatus) *RevokeDeviceSessionBatchResult {
	s := new(RevokeDeviceSessionBatchResult)
	s.RevokeDevicesStatus = RevokeDevicesStatus
	return s
}

// RevokeDeviceSessionError : has no documentation (yet)
type RevokeDeviceSessionError struct {
	dropbox.Tagged
}

// Valid tag values for RevokeDeviceSessionError
const (
	RevokeDeviceSessionErrorDeviceSessionNotFound = "device_session_not_found"
	RevokeDeviceSessionErrorMemberNotFound        = "member_not_found"
	RevokeDeviceSessionErrorOther                 = "other"
)

// RevokeDeviceSessionStatus : has no documentation (yet)
type RevokeDeviceSessionStatus struct {
	// Success : Result of the revoking request
	Success bool `json:"success"`
	// ErrorType : The error cause in case of a failure
	ErrorType *RevokeDeviceSessionError `json:"error_type,omitempty"`
}

// NewRevokeDeviceSessionStatus returns a new RevokeDeviceSessionStatus instance
func NewRevokeDeviceSessionStatus(Success bool) *RevokeDeviceSessionStatus {
	s := new(RevokeDeviceSessionStatus)
	s.Success = Success
	return s
}

// RevokeLinkedApiAppArg : has no documentation (yet)
type RevokeLinkedApiAppArg struct {
	// AppId : The application's unique id
	AppId string `json:"app_id"`
	// TeamMemberId : The unique id of the member owning the device
	TeamMemberId string `json:"team_member_id"`
	// KeepAppFolder : Whether to keep the application dedicated folder (in case
	// the application uses  one)
	KeepAppFolder bool `json:"keep_app_folder"`
}

// NewRevokeLinkedApiAppArg returns a new RevokeLinkedApiAppArg instance
func NewRevokeLinkedApiAppArg(AppId string, TeamMemberId string) *RevokeLinkedApiAppArg {
	s := new(RevokeLinkedApiAppArg)
	s.AppId = AppId
	s.TeamMemberId = TeamMemberId
	s.KeepAppFolder = true
	return s
}

// RevokeLinkedApiAppBatchArg : has no documentation (yet)
type RevokeLinkedApiAppBatchArg struct {
	// RevokeLinkedApp : has no documentation (yet)
	RevokeLinkedApp []*RevokeLinkedApiAppArg `json:"revoke_linked_app"`
}

// NewRevokeLinkedApiAppBatchArg returns a new RevokeLinkedApiAppBatchArg instance
func NewRevokeLinkedApiAppBatchArg(RevokeLinkedApp []*RevokeLinkedApiAppArg) *RevokeLinkedApiAppBatchArg {
	s := new(RevokeLinkedApiAppBatchArg)
	s.RevokeLinkedApp = RevokeLinkedApp
	return s
}

// RevokeLinkedAppBatchError : Error returned by
// `linkedAppsRevokeLinkedAppBatch`.
type RevokeLinkedAppBatchError struct {
	dropbox.Tagged
}

// Valid tag values for RevokeLinkedAppBatchError
const (
	RevokeLinkedAppBatchErrorUnspecified = "unspecified"
)

// RevokeLinkedAppBatchResult : has no documentation (yet)
type RevokeLinkedAppBatchResult struct {
	// RevokeLinkedAppStatus : has no documentation (yet)
	RevokeLinkedAppStatus []*RevokeLinkedAppStatus `json:"revoke_linked_app_status"`
}

// NewRevokeLinkedAppBatchResult returns a new RevokeLinkedAppBatchResult instance
func NewRevokeLinkedAppBatchResult(RevokeLinkedAppStatus []*RevokeLinkedAppStatus) *RevokeLinkedAppBatchResult {
	s := new(RevokeLinkedAppBatchResult)
	s.RevokeLinkedAppStatus = RevokeLinkedAppStatus
	return s
}

// RevokeLinkedAppError : Error returned by `linkedAppsRevokeLinkedApp`.
type RevokeLinkedAppError struct {
	dropbox.Tagged
}

// Valid tag values for RevokeLinkedAppError
const (
	RevokeLinkedAppErrorAppNotFound    = "app_not_found"
	RevokeLinkedAppErrorMemberNotFound = "member_not_found"
	RevokeLinkedAppErrorOther          = "other"
)

// RevokeLinkedAppStatus : has no documentation (yet)
type RevokeLinkedAppStatus struct {
	// Success : Result of the revoking request
	Success bool `json:"success"`
	// ErrorType : The error cause in case of a failure
	ErrorType *RevokeLinkedAppError `json:"error_type,omitempty"`
}

// NewRevokeLinkedAppStatus returns a new RevokeLinkedAppStatus instance
func NewRevokeLinkedAppStatus(Success bool) *RevokeLinkedAppStatus {
	s := new(RevokeLinkedAppStatus)
	s.Success = Success
	return s
}

// StorageBucket : Describes the number of users in a specific storage bucket.
type StorageBucket struct {
	// Bucket : The name of the storage bucket. For example, '1G' is a bucket of
	// users with storage size up to 1 Giga.
	Bucket string `json:"bucket"`
	// Users : The number of people whose storage is in the range of this
	// storage bucket.
	Users uint64 `json:"users"`
}

// NewStorageBucket returns a new StorageBucket instance
func NewStorageBucket(Bucket string, Users uint64) *StorageBucket {
	s := new(StorageBucket)
	s.Bucket = Bucket
	s.Users = Users
	return s
}

// TeamGetInfoResult : has no documentation (yet)
type TeamGetInfoResult struct {
	// Name : The name of the team.
	Name string `json:"name"`
	// TeamId : The ID of the team.
	TeamId string `json:"team_id"`
	// NumLicensedUsers : The number of licenses available to the team.
	NumLicensedUsers uint32 `json:"num_licensed_users"`
	// NumProvisionedUsers : The number of accounts that have been invited or
	// are already active members of the team.
	NumProvisionedUsers uint32 `json:"num_provisioned_users"`
	// Policies : has no documentation (yet)
	Policies *team_policies.TeamMemberPolicies `json:"policies"`
}

// NewTeamGetInfoResult returns a new TeamGetInfoResult instance
func NewTeamGetInfoResult(Name string, TeamId string, NumLicensedUsers uint32, NumProvisionedUsers uint32, Policies *team_policies.TeamMemberPolicies) *TeamGetInfoResult {
	s := new(TeamGetInfoResult)
	s.Name = Name
	s.TeamId = TeamId
	s.NumLicensedUsers = NumLicensedUsers
	s.NumProvisionedUsers = NumProvisionedUsers
	s.Policies = Policies
	return s
}

// TeamMemberInfo : Information about a team member.
type TeamMemberInfo struct {
	// Profile : Profile of a user as a member of a team.
	Profile *TeamMemberProfile `json:"profile"`
	// Role : The user's role in the team.
	Role *AdminTier `json:"role"`
}

// NewTeamMemberInfo returns a new TeamMemberInfo instance
func NewTeamMemberInfo(Profile *TeamMemberProfile, Role *AdminTier) *TeamMemberInfo {
	s := new(TeamMemberInfo)
	s.Profile = Profile
	s.Role = Role
	return s
}

// TeamMemberProfile : Profile of a user as a member of a team.
type TeamMemberProfile struct {
	MemberProfile
	// Groups : List of group IDs of groups that the user belongs to.
	Groups []string `json:"groups"`
}

// NewTeamMemberProfile returns a new TeamMemberProfile instance
func NewTeamMemberProfile(TeamMemberId string, Email string, EmailVerified bool, Status *TeamMemberStatus, Name *users.Name, MembershipType *TeamMembershipType, Groups []string) *TeamMemberProfile {
	s := new(TeamMemberProfile)
	s.TeamMemberId = TeamMemberId
	s.Email = Email
	s.EmailVerified = EmailVerified
	s.Status = Status
	s.Name = Name
	s.MembershipType = MembershipType
	s.Groups = Groups
	return s
}

// TeamMemberStatus : The user's status as a member of a specific team.
type TeamMemberStatus struct {
	dropbox.Tagged
}

// Valid tag values for TeamMemberStatus
const (
	TeamMemberStatusActive    = "active"
	TeamMemberStatusInvited   = "invited"
	TeamMemberStatusSuspended = "suspended"
)

// TeamMembershipType : has no documentation (yet)
type TeamMembershipType struct {
	dropbox.Tagged
}

// Valid tag values for TeamMembershipType
const (
	TeamMembershipTypeFull    = "full"
	TeamMembershipTypeLimited = "limited"
)

// UpdatePropertyTemplateArg : has no documentation (yet)
type UpdatePropertyTemplateArg struct {
	// TemplateId : An identifier for property template added by
	// `propertiesTemplateAdd`.
	TemplateId string `json:"template_id"`
	// Name : A display name for the property template. Property template names
	// can be up to 256 bytes.
	Name string `json:"name,omitempty"`
	// Description : Description for new property template. Property template
	// descriptions can be up to 1024 bytes.
	Description string `json:"description,omitempty"`
	// AddFields : This is a list of custom properties to add to the property
	// template. There can be up to 64 properties in a single property template.
	AddFields []*properties.PropertyFieldTemplate `json:"add_fields,omitempty"`
}

// NewUpdatePropertyTemplateArg returns a new UpdatePropertyTemplateArg instance
func NewUpdatePropertyTemplateArg(TemplateId string) *UpdatePropertyTemplateArg {
	s := new(UpdatePropertyTemplateArg)
	s.TemplateId = TemplateId
	return s
}

// UpdatePropertyTemplateResult : has no documentation (yet)
type UpdatePropertyTemplateResult struct {
	// TemplateId : An identifier for property template added by
	// `propertiesTemplateAdd`.
	TemplateId string `json:"template_id"`
}

// NewUpdatePropertyTemplateResult returns a new UpdatePropertyTemplateResult instance
func NewUpdatePropertyTemplateResult(TemplateId string) *UpdatePropertyTemplateResult {
	s := new(UpdatePropertyTemplateResult)
	s.TemplateId = TemplateId
	return s
}

// UserSelectorArg : Argument for selecting a single user, either by
// team_member_id, external_id or email.
type UserSelectorArg struct {
	dropbox.Tagged
	// TeamMemberId : has no documentation (yet)
	TeamMemberId string `json:"team_member_id,omitempty"`
	// ExternalId : has no documentation (yet)
	ExternalId string `json:"external_id,omitempty"`
	// Email : has no documentation (yet)
	Email string `json:"email,omitempty"`
}

// Valid tag values for UserSelectorArg
const (
	UserSelectorArgTeamMemberId = "team_member_id"
	UserSelectorArgExternalId   = "external_id"
	UserSelectorArgEmail        = "email"
)

// UnmarshalJSON deserializes into a UserSelectorArg instance
func (u *UserSelectorArg) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "team_member_id":
		if err := json.Unmarshal(body, &u.TeamMemberId); err != nil {
			return err
		}

	case "external_id":
		if err := json.Unmarshal(body, &u.ExternalId); err != nil {
			return err
		}

	case "email":
		if err := json.Unmarshal(body, &u.Email); err != nil {
			return err
		}

	}
	return nil
}

// UsersSelectorArg : Argument for selecting a list of users, either by
// team_member_ids, external_ids or emails.
type UsersSelectorArg struct {
	dropbox.Tagged
	// TeamMemberIds : List of member IDs.
	TeamMemberIds []string `json:"team_member_ids,omitempty"`
	// ExternalIds : List of external user IDs.
	ExternalIds []string `json:"external_ids,omitempty"`
	// Emails : List of email addresses.
	Emails []string `json:"emails,omitempty"`
}

// Valid tag values for UsersSelectorArg
const (
	UsersSelectorArgTeamMemberIds = "team_member_ids"
	UsersSelectorArgExternalIds   = "external_ids"
	UsersSelectorArgEmails        = "emails"
)

// UnmarshalJSON deserializes into a UsersSelectorArg instance
func (u *UsersSelectorArg) UnmarshalJSON(body []byte) error {
	type wrap struct {
		dropbox.Tagged
		// TeamMemberIds : List of member IDs.
		TeamMemberIds json.RawMessage `json:"team_member_ids,omitempty"`
		// ExternalIds : List of external user IDs.
		ExternalIds json.RawMessage `json:"external_ids,omitempty"`
		// Emails : List of email addresses.
		Emails json.RawMessage `json:"emails,omitempty"`
	}
	var w wrap
	if err := json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "team_member_ids":
		if err := json.Unmarshal(body, &u.TeamMemberIds); err != nil {
			return err
		}

	case "external_ids":
		if err := json.Unmarshal(body, &u.ExternalIds); err != nil {
			return err
		}

	case "emails":
		if err := json.Unmarshal(body, &u.Emails); err != nil {
			return err
		}

	}
	return nil
}
