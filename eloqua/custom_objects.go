package eloqua

import (
	"fmt"
)

// CustomObjectService provides access to all the endpoints related
// to custom object data within eloqua
//
// Eloqua API docs: https://goo.gl/cQyZYx
type CustomObjectService struct {
	client *Client
}

// CustomObject represents an Eloqua custom object.
// These fields are taken from an API response since the Eloqua documentation
// ,at the time of building, did not appear to be correct.
type CustomObject struct {
	Type        string `json:"type,omitempty"`
	ID          int    `json:"id,omitempty,string"`
	CreatedAt   int    `json:"createdAt,omitempty,string"`
	CreatedBy   int    `json:"createdBy,omitempty,string"`
	Depth       string `json:"depth,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	UpdatedAt   int    `json:"updatedAt,omitempty,string"`
	UpdatedBy   int    `json:"updatedBy,omitempty,string"`

	DisplayNameFieldID string              `json:"displayNameFieldId,omitempty"`
	ContentText        string              `json:"contentText,omitempty"`
	RecordCount        int                 `json:"recordCount,omitempty"`
	Fields             []CustomObjectField `json:"fields,omitempty"`
}

// CustomObjectField represents a database field within an Eloqua custom data object.
type CustomObjectField struct {
	Type         string `json:"type,omitempty"`
	ID           int    `json:"id,omitempty,string"`
	Depth        string `json:"depth,omitempty"`
	Name         string `json:"name,omitempty"`
	DataType     string `json:"dataType,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	DisplayType  string `json:"displayType,omitempty"`
	InternalName string `json:"internalName,omitempty"`
}

// Create a new custom object in eloqua
func (e *CustomObjectService) Create(name string, customObject *CustomObject) (*CustomObject, *Response, error) {
	if customObject == nil {
		customObject = &CustomObject{}
	}

	customObject.Name = name
	endpoint := "/assets/customObject"
	resp, err := e.client.postRequestDecode(endpoint, customObject)
	return customObject, resp, err
}

// Get a custom object via its ID
func (e *CustomObjectService) Get(id int) (*CustomObject, *Response, error) {
	endpoint := fmt.Sprintf("/assets/customObject/%d?depth=complete", id)
	customObject := &CustomObject{}
	resp, err := e.client.getRequestDecode(endpoint, customObject)
	return customObject, resp, err
}

// List many eloqua custom objects
func (e *CustomObjectService) List(opts *ListOptions) ([]CustomObject, *Response, error) {
	endpoint := "/assets/customObjects"
	customObjects := new([]CustomObject)
	resp, err := e.client.getRequestListDecode(endpoint, customObjects, opts)
	return *customObjects, resp, err
}

// Update an existing custom object in eloqua
func (e *CustomObjectService) Update(id int, name string, customObject *CustomObject) (*CustomObject, *Response, error) {
	if customObject == nil {
		customObject = &CustomObject{}
	}

	customObject.ID = id
	customObject.Name = name

	endpoint := fmt.Sprintf("/assets/customObject/%d", customObject.ID)
	resp, err := e.client.putRequestDecode(endpoint, customObject)
	return customObject, resp, err
}

// Delete an existing custom object from eloqua
func (e *CustomObjectService) Delete(id int) (*Response, error) {
	customObject := &CustomObject{ID: id}
	endpoint := fmt.Sprintf("/assets/customObject/%d", customObject.ID)
	resp, err := e.client.deleteRequest(endpoint, customObject)
	return resp, err
}
