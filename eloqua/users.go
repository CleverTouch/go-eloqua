package eloqua

import (
	"fmt"
)

// UserService provides access to all endpoints related to
// managing Eloqua system users.
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Users/users-API.htm
type UserService struct {
	client *Client
}

// User represents an Eloqua system user.
type User struct {
	Type                 string   `json:"type,omitempty"`
	AccessedAt           int      `json:"accessedAt,omitempty,string"`
	CurrentStatus        string   `json:"currentStatus,omitempty"`
	ID                   int      `json:"id,omitempty,string"`
	CreatedAt            int      `json:"createdAt,omitempty,string"`
	CreatedBy            int      `json:"createdBy,omitempty,string"`
	Description          string   `json:"description,omitempty"`
	Depth                string   `json:"depth,omitempty"`
	FolderID             int      `json:"folderId,omitempty,string"`
	Name                 string   `json:"name,omitempty"`
	Permissions          []string `json:"permissions,omitempty"`
	UpdatedAt            int      `json:"updatedAt,omitempty,string"`
	UpdatedBy            int      `json:"updatedBy,omitempty,string"`
	ScheduledFor         int      `json:"scheduledFor,omitempty,string"`
	SourceTemplateID     string   `json:"sourceTemplateId,omitempty"`
	BetaAccess           []string `json:"betaAccess,omitempty"`
	Capabilities         []string `json:"capabilities,omitempty"`
	Company              string   `json:"company,omitempty"`
	DefaultAccountViewID int      `json:"defaultAccountViewId,omitempty,string"`
	DefaultContactViewID int      `json:"defaultContactViewId,omitempty,string"`
	EmailAddress         string   `json:"emailAddress,omitempty"`

	LoggedInAt string `json:"loggedInAt,omitempty"`
	LoginName  string `json:"loginName,omitempty"`

	InterfacePermissions []InterfacePermission `json:"interfacePermissions,omitempty"`
	Preferences          UserPreferences       `json:"preferences,omitempty"`
	TypePermissions      []TypePermission      `json:"typePermissions,omitempty"`
	ProductPermissions   []ProductPermission   `json:"productPermissions,omitempty"`
}

// UserPreferences holds user-specific Eloqua preferences.
// This may be limited compared to what can possibly be fetched/set.
type UserPreferences struct {
	Type       string `json:"type,omitempty"`
	TimezoneID int    `json:"timezoneId,omitempty,string"`
}

// InterfacePermission is a permission assigned to a user to control
// the parts of the Eloqua interface they see.
type InterfacePermission struct {
	Type                       string                `json:"type,omitempty"`
	InterfaceCode              string                `json:"interfaceCode,omitempty"`
	NestedInterfacePermissions []InterfacePermission `json:"nestedInterfacePermissions,omitempty"`
}

// TypePermission represents the user's permissions for a particular Eloqua Type.
type TypePermission struct {
	Type        string          `json:"type,omitempty"`
	ObjectType  string          `json:"objectType,omitempty"`
	Permissions TypePermissions `json:"permissions,omitempty"`
}

// TypePermissions are the actual permission set on a TypePermission.
type TypePermissions struct {
	Type   string `json:"type,omitempty"`
	Create bool   `json:"create,omitempty,string"`
}

// ProductPermission displays the access that the user has to particular Eloqua prouducts such as Profiler.
type ProductPermission struct {
	Type        string `json:"type,omitempty"`
	ProductCode string `json:"productCode,omitempty"`
}

// Get an user object via its ID
func (e *UserService) Get(id int) (*User, *Response, error) {
	endpoint := fmt.Sprintf("/system/user/%d?depth=complete", id)
	user := &User{}
	resp, err := e.client.getRequestDecode(endpoint, user)
	return user, resp, err
}

// List many Eloqua users
func (e *UserService) List(opts *ListOptions) ([]User, *Response, error) {
	endpoint := "/system/users"
	users := new([]User)
	resp, err := e.client.getRequestListDecode(endpoint, users, opts)
	return *users, resp, err
}

// Update an existing user in eloqua
// This endpoint does not seem to be fully stable and/or working fully
// Could not get reliably functioning during testing
func (e *UserService) Update(id int, name string, user *User) (*User, *Response, error) {
	if user == nil {
		user = &User{}
	}
	user.ID = id
	user.Name = name
	endpoint := fmt.Sprintf("/system/user/%d", user.ID)
	resp, err := e.client.putRequestDecode(endpoint, user)
	return user, resp, err
}
