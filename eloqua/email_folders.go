package eloqua

import (
	"fmt"
)

// EmailFolderService provides access to all the endpoints related
// to email folder data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Email folders/emailFolders-API.htm
type EmailFolderService struct {
	client *Client
}

// EmailFolder represents an Eloqua email folder object.
// Segments that are not listed in the EmailFolder model itself can be retrieved/updated
// using the 'SegmentValues' property.
type EmailFolder struct {
	Type         string `json:"type,omitempty"`
	ID           int    `json:"id,omitempty,string"`
	CreatedAt    int    `json:"createdAt,omitempty,string"`
	RequestDepth string `json:"depth,omitempty"`

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FolderId    int    `json:"folderId,omitempty,string"`
	UpdatedAt   int    `json:"updatedAt,omitempty,string"`
	UpdatedBy   int    `json:"updatedBy,omitempty,string"`
	IsSystem    bool   `json:"isSystem,omitempty,string"`
	Archive     bool   `json:"archive,omitempty,string"`
}

// Create a new email folder in eloqua
// The email must not already exists otherwise Eloqua will return an error.
func (e *EmailFolderService) Create(name string, emailFolder *EmailFolder) (*EmailFolder, *Response, error) {
	if emailFolder == nil {
		emailFolder = &EmailFolder{}
	}
	emailFolder.Name = name

	endpoint := "/assets/email/folder"
	resp, err := e.client.postRequestDecode(endpoint, emailFolder)
	return emailFolder, resp, err
}

// Get an email folder object via its ID
func (e *EmailFolderService) Get(id int) (*EmailFolder, *Response, error) {
	endpoint := fmt.Sprintf("/assets/email/folder/%d?depth=complete", id)
	emailFolder := &EmailFolder{}
	resp, err := e.client.getRequestDecode(endpoint, emailFolder)
	return emailFolder, resp, err
}

// List many eloqua email folders
func (e *EmailFolderService) List(opts *ListOptions) ([]EmailFolder, *Response, error) {
	endpoint := "/assets/email/folders"
	emailFolders := new([]EmailFolder)
	resp, err := e.client.getRequestListDecode(endpoint, emailFolders, opts)
	return *emailFolders, resp, err
}

// Update an existing email folder in eloqua
func (e *EmailFolderService) Update(id int, name string, emailFolder *EmailFolder) (*EmailFolder, *Response, error) {
	if emailFolder == nil {
		emailFolder = &EmailFolder{}
	}

	emailFolder.ID = id
	emailFolder.Name = name

	endpoint := fmt.Sprintf("/assets/email/folder/%d", emailFolder.ID)
	resp, err := e.client.putRequestDecode(endpoint, emailFolder)
	return emailFolder, resp, err
}

// Delete an existing email folder from eloqua
func (e *EmailFolderService) Delete(id int) (*Response, error) {
	emailFolder := &EmailFolder{ID: id}
	endpoint := fmt.Sprintf("/assets/email/folder/%d", emailFolder.ID)
	resp, err := e.client.deleteRequest(endpoint, emailFolder)
	return resp, err
}
