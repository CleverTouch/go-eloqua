package eloqua

import (
	"fmt"
)

// EmailFooterService provides access to all the endpoints related
// to email footer data within eloqua
//
// The official Eloqua documentation did not seem to contain footers but everything
// appears to be the same as EmailHeaders references in the below link.
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Email headers/emailHeaders-API.htm
type EmailFooterService struct {
	client *Client
}

// EmailFooter represents an Eloqua email footer object.
type EmailFooter struct {
	Type      string `json:"type,omitempty"`
	ID        int    `json:"id,omitempty,string"`
	CreatedAt int    `json:"createdAt,omitempty,string"`
	CreatedBy int    `json:"createdBy,omitempty,string"`
	Depth     string `json:"depth,omitempty"`

	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	FolderID    int      `json:"folderId,omitempty,string"`
	UpdatedAt   int      `json:"updatedAt,omitempty,string"`
	UpdatedBy   int      `json:"updatedBy,omitempty,string"`

	Body                string `json:"body,omitempty"`
	PlainText           string `json:"plainText,omitempty"`
	IsPlainTextEditable bool   `json:"isPlainTextEditable,omitempty,string"`

	FieldMerges []FieldMerge `json:"fieldMerges,omitempty"`
	Images      []Image      `json:"images,omitempty"`
	Hyperlinks  []Hyperlink  `json:"hyperlinks,omitempty"`
}

// Create a new email footer in eloqua
func (e *EmailFooterService) Create(name string, emailFooter *EmailFooter) (*EmailFooter, *Response, error) {
	if emailFooter == nil {
		emailFooter = &EmailFooter{}
	}
	emailFooter.Name = name

	endpoint := "/assets/email/footer"
	resp, err := e.client.postRequestDecode(endpoint, emailFooter)
	return emailFooter, resp, err
}

// Get an email footer object via its ID
func (e *EmailFooterService) Get(id int) (*EmailFooter, *Response, error) {
	endpoint := fmt.Sprintf("/assets/email/footer/%d?depth=complete", id)
	emailFooter := &EmailFooter{}
	resp, err := e.client.getRequestDecode(endpoint, emailFooter)
	return emailFooter, resp, err
}

// List many eloqua email footers
func (e *EmailFooterService) List(opts *ListOptions) ([]EmailFooter, *Response, error) {
	endpoint := "/assets/email/footers"
	emailFooters := new([]EmailFooter)
	resp, err := e.client.getRequestListDecode(endpoint, emailFooters, opts)
	return *emailFooters, resp, err
}

// Update an existing email footer in eloqua
func (e *EmailFooterService) Update(id int, name string, emailFooter *EmailFooter) (*EmailFooter, *Response, error) {
	if emailFooter == nil {
		emailFooter = &EmailFooter{}
	}

	emailFooter.ID = id
	emailFooter.Name = name

	endpoint := fmt.Sprintf("/assets/email/footer/%d", emailFooter.ID)
	resp, err := e.client.putRequestDecode(endpoint, emailFooter)
	return emailFooter, resp, err
}

// Delete an existing email footer from eloqua
func (e *EmailFooterService) Delete(id int) (*Response, error) {
	emailFooter := &EmailFooter{ID: id}
	endpoint := fmt.Sprintf("/assets/email/footer/%d", emailFooter.ID)
	resp, err := e.client.deleteRequest(endpoint, emailFooter)
	return resp, err
}
