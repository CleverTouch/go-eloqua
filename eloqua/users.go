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
	// TODO - interfacePermissions
	LoggedInAt string `json:"loggedInAt,omitempty"`
	LoginName  string `json:"loginName,omitempty"`
	// TODO - preferences
	// TODO - productPermissions
	// TODO - typePermissions
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
