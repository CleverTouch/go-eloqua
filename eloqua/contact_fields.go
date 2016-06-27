package eloqua

import (
	"fmt"
)

// ContactFieldService provides access to all the endpoints related
// to contact field data within eloqua
//
// Eloqua API docs: https://goo.gl/Y8sF8P
type ContactFieldService struct {
	client *Client
}

// ContactField represents an Eloqua contact field object.
// Fields that are not listed in the ContactField model itself can be retrieved/updated
// using the 'FieldValues' property.
type ContactField struct {
	Type      string `json:"type,omitempty"`
	ID        int    `json:"id,omitempty,string"`
	CreatedAt int    `json:"createdAt,omitempty,string"`
	Depth     string `json:"depth,omitempty"`
	Name      string `json:"name,omitempty"`
	UpdatedAt int    `json:"updatedAt,omitempty,string"`

	DataType     string `json:"dataType,omitempty"`
	DisplayType  string `json:"displayType,omitempty"`
	InternalName string `json:"internalName,omitempty"`
	IsReadOnly   bool   `json:"isReadOnly,string"`
	IsRequired   bool   `json:"isRequired,string"`
	IsStandard   bool   `json:"isStandard,string"`
	IsProtected  bool   `json:"isProtected,string"`

	IsPopulatedInOutlookPlugin bool   `json:"isPopulatedInOutlookPlugin,string"`
	UpdateType                 string `json:"updateType,omitempty"`
}

// Create a new contact field in eloqua
func (e *ContactFieldService) Create(name string, dataType string, displayType string, updateType string, contactField *ContactField) (*ContactField, *Response, error) {
	if contactField == nil {
		contactField = &ContactField{}
	}

	contactField.Name = name
	contactField.DataType = dataType
	contactField.DisplayType = displayType
	contactField.UpdateType = updateType
	// Undocumented by seemed to be required during testing
	contactField.IsProtected = false

	endpoint := "/assets/contact/field"
	resp, err := e.client.postRequestDecode(endpoint, contactField)
	return contactField, resp, err
}

// Get an contact field object via its ID
func (e *ContactFieldService) Get(id int) (*ContactField, *Response, error) {
	endpoint := fmt.Sprintf("/assets/contact/field/%d?depth=complete", id)
	contactField := &ContactField{}
	resp, err := e.client.getRequestDecode(endpoint, contactField)
	return contactField, resp, err
}

// List many eloqua contact fields
func (e *ContactFieldService) List(opts *ListOptions) ([]ContactField, *Response, error) {
	endpoint := "/assets/contact/fields"
	contactFields := new([]ContactField)
	resp, err := e.client.getRequestListDecode(endpoint, contactFields, opts)
	return *contactFields, resp, err
}

// Update an existing contact field in eloqua
func (e *ContactFieldService) Update(id int, name string, dataType string, displayType string, updateType string, contactField *ContactField) (*ContactField, *Response, error) {
	if contactField == nil {
		contactField = &ContactField{}
	}

	contactField.ID = id
	contactField.Name = name
	contactField.DataType = dataType
	contactField.DisplayType = displayType
	contactField.UpdateType = updateType

	endpoint := fmt.Sprintf("/assets/contact/field/%d", contactField.ID)
	resp, err := e.client.putRequestDecode(endpoint, contactField)
	return contactField, resp, err
}

// Delete an existing contact field from eloqua
func (e *ContactFieldService) Delete(id int) (*Response, error) {
	contactField := &ContactField{ID: id}
	endpoint := fmt.Sprintf("/assets/contact/field/%d", contactField.ID)
	resp, err := e.client.deleteRequest(endpoint, contactField)
	return resp, err
}
