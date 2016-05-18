package eloqua

import (
	"fmt"
)

// ContactSegmentService provides access to all the endpoints related
// to contact segment data within eloqua
//
// Eloqua API docs: https://docs.oracle.com/cloud/latest/marketingcs_gs/OMCAB/#Developers/RESTAPI/1.0 Endpoints/Contact segments/contactSegments-API.htm
type ContactSegmentService struct {
	client *Client
}

// ContactSegment represents an Eloqua contact segment object.
// Segments that are not listed in the ContactSegment model itself can be retrieved/updated
// using the 'SegmentValues' property.
type ContactSegment struct {
	Type          string `json:"type,omitempty"`
	CurrentStatus string `json:"currentStatus,omitempty"`
	ID            int    `json:"id,omitempty,string"`
	CreatedAt     int    `json:"createdAt,omitempty,string"`
	CreatedBy     int    `json:"createdBy,omitempty,string"`
	RequestDepth  string `json:"depth,omitempty"`

	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	FolderId    int      `json:"folderId,omitempty,string"`
	UpdatedAt   int      `json:"updatedAt,omitempty,string"`
	UpdatedBy   int      `json:"updatedBy,omitempty,string"`
	Permissions []string `json:"permissions,omitempty"`
	Count       int      `json:"count,omitempty,string"`

	// Todo - Elements
}

// Create a new contact segment in eloqua
// The email must not already exists otherwise Eloqua will return an error.
func (e *ContactSegmentService) Create(name string, contactSegment *ContactSegment) (*ContactSegment, *Response, error) {
	if contactSegment == nil {
		contactSegment = &ContactSegment{}
	}
	contactSegment.Name = name

	endpoint := "/assets/contact/segment"
	resp, err := e.client.postRequestDecode(endpoint, contactSegment)
	return contactSegment, resp, err
}

// Get an contact segment object via its ID
func (e *ContactSegmentService) Get(id int) (*ContactSegment, *Response, error) {
	endpoint := fmt.Sprintf("/assets/contact/segment/%d?depth=complete", id)
	contactSegment := &ContactSegment{}
	resp, err := e.client.getRequestDecode(endpoint, contactSegment)
	return contactSegment, resp, err
}

// List many eloqua contact segments
func (e *ContactSegmentService) List(opts *ListOptions) ([]ContactSegment, *Response, error) {
	endpoint := "/assets/contact/segments"
	contactSegments := new([]ContactSegment)
	resp, err := e.client.getRequestListDecode(endpoint, contactSegments, opts)
	return *contactSegments, resp, err
}

// Update an existing contact segment in eloqua
func (e *ContactSegmentService) Update(id int, name string, contactSegment *ContactSegment) (*ContactSegment, *Response, error) {
	if contactSegment == nil {
		contactSegment = &ContactSegment{}
	}

	contactSegment.ID = id
	contactSegment.Name = name

	endpoint := fmt.Sprintf("/assets/contact/segment/%d", contactSegment.ID)
	resp, err := e.client.putRequestDecode(endpoint, contactSegment)
	return contactSegment, resp, err
}

// Delete an existing contact segment from eloqua
func (e *ContactSegmentService) Delete(id int) (*Response, error) {
	contactSegment := &ContactSegment{ID: id}
	endpoint := fmt.Sprintf("/assets/contact/segment/%d", contactSegment.ID)
	resp, err := e.client.deleteRequest(endpoint, contactSegment)
	return resp, err
}
