package eloqua

import (
	"fmt"
)

// EmailHeaderService provides access to all the endpoints related
// to email header data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Email headers/post-assets-emailHeader.htm
type EmailHeaderService struct {
	client *Client
}

// EmailHeader represents an Eloqua email header object.
type EmailHeader struct {
	Type         string `json:"type,omitempty"`
	ID           int    `json:"id,omitempty,string"`
	CreatedAt    int    `json:"createdAt,omitempty,string"`
	CreatedBy    int    `json:"createdBy,omitempty,string"`
	RequestDepth string `json:"depth,omitempty"`

	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	FolderID    int      `json:"folderId,omitempty,string"`
	UpdatedAt   int      `json:"updatedAt,omitempty,string"`
	UpdatedBy   int      `json:"updatedBy,omitempty,string"`

	Body                string `json:"body,omitempty"`
	PlainText           string `json:"plainText,omitempty"`
	IsPlainTextEditable bool   `json:"isPlainTextEditable,omitempty,string"`

	FieldMerges []FieldMerge `json:"hyperlinks,fieldMerges"`
	Images      []Image      `json:"images,omitempty"`
	Hyperlinks  []Hyperlink  `json:"hyperlinks,omitempty"`
}

// Create a new email header in eloqua
func (e *EmailHeaderService) Create(name string, emailHeader *EmailHeader) (*EmailHeader, *Response, error) {
	if emailHeader == nil {
		emailHeader = &EmailHeader{}
	}
	emailHeader.Name = name

	endpoint := "/assets/email/header"
	resp, err := e.client.postRequestDecode(endpoint, emailHeader)
	return emailHeader, resp, err
}

// Get an email header object via its ID
func (e *EmailHeaderService) Get(id int) (*EmailHeader, *Response, error) {
	endpoint := fmt.Sprintf("/assets/email/header/%d?depth=complete", id)
	emailHeader := &EmailHeader{}
	resp, err := e.client.getRequestDecode(endpoint, emailHeader)
	return emailHeader, resp, err
}

// List many eloqua email headers
func (e *EmailHeaderService) List(opts *ListOptions) ([]EmailHeader, *Response, error) {
	endpoint := "/assets/email/headers"
	emailHeaders := new([]EmailHeader)
	resp, err := e.client.getRequestListDecode(endpoint, emailHeaders, opts)
	return *emailHeaders, resp, err
}

// Update an existing email header in eloqua
func (e *EmailHeaderService) Update(id int, name string, emailHeader *EmailHeader) (*EmailHeader, *Response, error) {
	if emailHeader == nil {
		emailHeader = &EmailHeader{}
	}

	emailHeader.ID = id
	emailHeader.Name = name

	endpoint := fmt.Sprintf("/assets/email/header/%d", emailHeader.ID)
	resp, err := e.client.putRequestDecode(endpoint, emailHeader)
	return emailHeader, resp, err
}

// Delete an existing email header from eloqua
func (e *EmailHeaderService) Delete(id int) (*Response, error) {
	emailHeader := &EmailHeader{ID: id}
	endpoint := fmt.Sprintf("/assets/email/header/%d", emailHeader.ID)
	resp, err := e.client.deleteRequest(endpoint, emailHeader)
	return resp, err
}
