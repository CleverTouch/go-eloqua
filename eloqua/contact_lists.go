package eloqua

import (
	"fmt"
)

// ContactListService provides access to all the endpoints related
// to contact list data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Contact lists/contactLists-API.htm
type ContactListService struct {
	client *Client
}

// ContactList represents an Eloqua email object.
// Fields that are not listed in the ContactList model itself can be retrieved/updated
// using the 'FieldValues' property.
type ContactList struct {
	Type         string   `json:"type,omitempty"`
	ID           int      `json:"id,omitempty,string"`
	CreatedAt    int      `json:"createdAt,omitempty,string"`
	RequestDepth string   `json:"depth,omitempty"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	UpdatedAt    int      `json:"updatedAt,omitempty,string"`
	FolderID     int      `json:"folderId,omitempty,string"`
	Permissions  []string `json:"permissions,omitempty"`
	Count        int      `json:"count,omitempty,string"`
	DataLookupId string   `json:"dataLookupId,omitempty"`
	Scope        string   `json:"scope,omitempty"`

	// Used to add contact ID's to to add or delete from a list
	// Writeonly, Not listed in official Eloqua developer documents
	MembershipAdditions []int `json:"membershipAdditions,omitempty,string"`
	MembershipDeletions []int `json:"membershipDeletions,omitempty,string"`
}

// Create a new contact list in eloqua
// The email must not already exists otherwise Eloqua will return an error.
func (e *ContactListService) Create(name string, contactList *ContactList) (*ContactList, *Response, error) {
	if contactList == nil {
		contactList = &ContactList{}
	}

	contactList.Name = name
	endpoint := "/assets/contact/list"
	resp, err := e.client.postRequestDecode(endpoint, contactList)
	return contactList, resp, err
}

// Get a contact list object via its ID
func (e *ContactListService) Get(id int) (*ContactList, *Response, error) {
	endpoint := fmt.Sprintf("/assets/contact/list/%d?depth=complete", id)
	contactList := &ContactList{}
	resp, err := e.client.getRequestDecode(endpoint, contactList)
	return contactList, resp, err
}

// List many eloqua contact lists
func (e *ContactListService) List(opts *ListOptions) ([]ContactList, *Response, error) {
	endpoint := "/assets/contact/lists"
	contactLists := new([]ContactList)
	resp, err := e.client.getRequestListDecode(endpoint, contactLists, opts)
	return *contactLists, resp, err
}

// Update an existing contact list in eloqua
func (e *ContactListService) Update(id int, name string, contactList *ContactList) (*ContactList, *Response, error) {
	if contactList == nil {
		contactList = &ContactList{}
	}

	contactList.ID = id
	contactList.Name = name

	endpoint := fmt.Sprintf("/assets/contact/list/%d", contactList.ID)
	resp, err := e.client.putRequestDecode(endpoint, contactList)
	return contactList, resp, err
}

// Delete an existing contact list from eloqua
func (e *ContactListService) Delete(id int) (*Response, error) {
	contactList := &ContactList{ID: id}
	endpoint := fmt.Sprintf("/assets/contact/list/%d", contactList.ID)
	resp, err := e.client.deleteRequest(endpoint, contactList)
	return resp, err
}
