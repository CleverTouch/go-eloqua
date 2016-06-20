package eloqua

import (
	"fmt"
)

// ContentSectionService provides access to all the endpoints related
// to content section data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Content sections/get-assets-contentSection.htm
type ContentSectionService struct {
	client *Client
}

// ContentSection represents an Eloqua content section object.
// In Eloqua these are known as 'Shared Content'
type ContentSection struct {
	Type        string      `json:"type,omitempty"`
	ID          int         `json:"id,omitempty,string"`
	CreatedAt   int         `json:"createdAt,omitempty,string"`
	CreatedBy   int         `json:"createdBy,omitempty,string"`
	Depth       string      `json:"depth,omitempty"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	UpdatedAt   int         `json:"updatedAt,omitempty,string"`
	UpdatedBy   int         `json:"updatedBy,omitempty,string"`
	ContentHTML string      `json:"contentHtml,omitempty"`
	ContentText string      `json:"contentText,omitempty"`
	Scope       string      `json:"scope,omitempty"`
	Forms       []Form      `json:"forms,omitempty"`
	Hyperlinks  []Hyperlink `json:"hyperlinks,omitempty"`
	Images      []Image     `json:"images,omitempty"`
	Size        Size        `json:"size,omitempty"`
}

// Create a new content section in eloqua
func (e *ContentSectionService) Create(name string, contentSection *ContentSection) (*ContentSection, *Response, error) {
	if contentSection == nil {
		contentSection = &ContentSection{}
	}

	contentSection.Name = name
	endpoint := "/assets/contentSection"
	resp, err := e.client.postRequestDecode(endpoint, contentSection)
	return contentSection, resp, err
}

// Get a content section object via its ID
func (e *ContentSectionService) Get(id int) (*ContentSection, *Response, error) {
	endpoint := fmt.Sprintf("/assets/contentSection/%d?depth=complete", id)
	contentSection := &ContentSection{}
	resp, err := e.client.getRequestDecode(endpoint, contentSection)
	return contentSection, resp, err
}

// List many eloqua content sections
func (e *ContentSectionService) List(opts *ListOptions) ([]ContentSection, *Response, error) {
	endpoint := "/assets/contentSections"
	contentSections := new([]ContentSection)
	resp, err := e.client.getRequestListDecode(endpoint, contentSections, opts)
	return *contentSections, resp, err
}

// Update an existing content section in eloqua
func (e *ContentSectionService) Update(id int, name string, contentSection *ContentSection) (*ContentSection, *Response, error) {
	if contentSection == nil {
		contentSection = &ContentSection{}
	}

	contentSection.ID = id
	contentSection.Name = name

	endpoint := fmt.Sprintf("/assets/contentSection/%d", contentSection.ID)
	resp, err := e.client.putRequestDecode(endpoint, contentSection)
	return contentSection, resp, err
}

// Delete an existing content section from eloqua
func (e *ContentSectionService) Delete(id int) (*Response, error) {
	contentSection := &ContentSection{ID: id}
	endpoint := fmt.Sprintf("/assets/contentSection/%d", contentSection.ID)
	resp, err := e.client.deleteRequest(endpoint, contentSection)
	return resp, err
}
