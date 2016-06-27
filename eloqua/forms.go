package eloqua

import (
	"fmt"
)

// FormService provides access to all the endpoints related
// to form assets within eloqua
//
// Eloqua API docs: https://goo.gl/WW0ehX
type FormService struct {
	client *Client
}

// Form represents an Eloqua form object.
type Form struct {
	Type          string      `json:"Type,omitempty"`
	CurrentStatus string      `json:"currentStatus,omitempty"`
	ID            int         `json:"id,omitempty,string"`
	CreatedAt     int         `json:"createdAt,omitempty,string"`
	CreatedBy     int         `json:"createdBy,omitempty,string"`
	Depth         string      `json:"depth,omitempty"`
	FolderID      int         `json:"folderId,omitempty,string"`
	Name          string      `json:"name,omitempty"`
	Permissions   []string    `json:"permissions,omitempty"`
	UpdatedAt     int         `json:"updatedAt,omitempty,string"`
	UpdatedBy     int         `json:"updatedBy,omitempty,string"`
	FormFields    []FormField `json:"elements,omitempty"`

	EmailAddressFormFieldID int        `json:"emailAddressFormFieldId,omitempty,string"`
	HTML                    string     `json:"html,omitempty"`
	HTMLName                string     `json:"htmlName,omitempty"`
	ProcessingType          string     `json:"processingType,omitempty"`
	ProcessingSteps         []FormStep `json:"processingSteps,omitempty"`
	Size                    Size       `json:"size,omitempty"`
	Style                   string     `json:"style,omitempty"`
}

// FormField represents the Eloqua model for a single field
// and its settings on a form.
type FormField struct {
	Type         string `json:"type,omitempty"`
	ID           int    `json:"id,omitempty,string"`
	Name         string `json:"name,omitempty"`
	Instructions string `json:"instructions,omitempty"`
	Style        string `json:"style,omitempty"`
	DataType     string `json:"dataType,omitempty"`
	DisplayType  string `json:"displayType,omitempty"`
	FieldMergeID int    `json:"fieldMergeId,omitempty,string"`
	HTMLName     string `json:"htmlName,omitempty"`

	Validations []FieldValidation `json:"validations,omitempty"`

	CreatedFromContactFieldID int `json:"createdFromContactFieldID,omitempty,string"`
}

// FieldValidation represents the Eloqua model for a validation
// rule, Typically found on a form field
type FieldValidation struct {
	Type        string     `json:"type,omitempty"`
	ID          int        `json:"id,omitempty,string"`
	Depth       string     `json:"depth,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Condition   TypeObject `json:"condition,omitempty"`
	IsEnabled   bool       `json:"isEnabled,omitempty,string"`
	Message     string     `json:"message,omitempty"`
}

// FormStep is a generic representation of an Eloqua form Step object.
// Within Eloqua there is really a range of 'FormStep*' object.
// FormStep just holds the common fields.
//
// A list of the various 'FormStep*' types can be found here:
// https://goo.gl/EhRfbg
type FormStep struct {
	Type     string `json:"type,omitempty"`
	ID       int    `json:"id,omitempty,string"`
	Name     string `json:"name,omitempty"`
	Exectute string `json:"execute,omitempty"`
}

// Create a new form in eloqua
func (e *FormService) Create(name string, form *Form) (*Form, *Response, error) {
	if form == nil {
		form = &Form{}
	}
	form.Name = name
	endpoint := "/assets/form"
	resp, err := e.client.postRequestDecode(endpoint, form)
	return form, resp, err
}

// Get an form object via its ID
func (e *FormService) Get(id int) (*Form, *Response, error) {
	endpoint := fmt.Sprintf("/assets/form/%d?depth=complete", id)
	form := &Form{}
	resp, err := e.client.getRequestDecode(endpoint, form)
	return form, resp, err
}

// List many Eloqua form objetcs
func (e *FormService) List(opts *ListOptions) ([]Form, *Response, error) {
	endpoint := "/assets/forms"
	forms := new([]Form)
	resp, err := e.client.getRequestListDecode(endpoint, forms, opts)
	return *forms, resp, err
}

// Update an existing form in eloqua
func (e *FormService) Update(id int, name string, form *Form) (*Form, *Response, error) {
	if form == nil {
		form = &Form{}
	}
	form.ID = id
	form.Name = name
	endpoint := fmt.Sprintf("/assets/form/%d", form.ID)
	resp, err := e.client.putRequestDecode(endpoint, form)
	return form, resp, err
}

// Delete an existing form from eloqua
func (e *FormService) Delete(id int) (*Response, error) {
	form := &Form{ID: id}
	endpoint := fmt.Sprintf("/assets/form/%d", form.ID)
	resp, err := e.client.deleteRequest(endpoint, form)
	return resp, err
}
